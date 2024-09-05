package main

import (
	"cmp"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"slices"
	"strconv"
	"strings"
	"syscall"
	"time"

	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/robot"
	"go.viam.com/rdk/robot/client"
	"go.viam.com/utils/rpc"
)

const uri = "revision-demo-main.4uju10uqzz.viamstg.cloud"

var hexRev = false

func formatRevision(rev string) string {
	if hexRev {
		i, err := strconv.Atoi(rev)
		if err != nil {
			log.Fatal("error parsing revision to int", err)
		}

		return fmt.Sprintf("%x", i)
	}

	return rev
}

func printStatus(s robot.MachineStatus) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("%-*s %s @ %s\n\n",
		14,
		"CONFIG",
		formatRevision(s.Config.Revision),
		s.Config.LastUpdated.Format(time.Kitchen)),
	)

	pad := []int{15, 20, 15}

	header := fmt.Sprintf("%-*s%-*s%-*s%s\n",
		pad[0], "RESOURCE", pad[1], "REV", pad[2], "STATE", "UPDATED")
	sb.WriteString(header)
	slices.SortStableFunc(s.Resources, func(r1, r2 resource.Status) int {
		return cmp.Compare(r1.Name.String(), r2.Name.String())
	})
	for _, res := range s.Resources {
		switch res.Name.Name {
		case "builtin", "deferred-manager":
			continue
		}
		line := fmt.Sprintf("%-*s%-*s%-*s%s\n",
			pad[0],
			res.Name.ShortName(),
			pad[1],
			formatRevision(res.Revision),
			pad[2],
			res.State,
			res.LastUpdated.Format(time.Kitchen),
		)
		sb.WriteString(line)
	}
	fmt.Print(sb.String())
}

func clearTerminal() {
	// https://stackoverflow.com/questions/22891645/how-can-i-clear-the-terminal-screen-in-go
	fmt.Print("\033[H\033[2J")
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	logger := logging.NewDebugLogger("client")
	logger.SetLevel(logging.ERROR)
	ctx := context.Background()

	apiKey, found := os.LookupEnv("VIAM_API_KEY")
	if !found {
		logger.Fatal(`missing VIAM_API_KEY`)
	}
	apiSecret, found := os.LookupEnv("VIAM_API_SECRET")
	if !found {
		logger.Fatal(`missing VIAM_API_SECRET`)
	}

	machine, err := client.New(
		ctx,
		uri,
		logger,
		client.WithDialOptions(rpc.WithEntityCredentials(
			apiKey,
			rpc.Credentials{
				Type:    rpc.CredentialsTypeAPIKey,
				Payload: apiSecret,
			})),
	)
	if err != nil {
		logger.Fatal(err)
	}

	for {
		select {
		case sig := <-sigs:
			switch sig {
			case syscall.SIGTERM:
				clearTerminal()
				fmt.Println("reloading...")
			case syscall.SIGINT:
				fmt.Println("bye!")
			}
			return
		default:
			s, err := machine.MachineStatus(ctx)
			if err != nil {
				logger.Fatal(err)
			}

			clearTerminal()
			printStatus(s)
			time.Sleep(300 * time.Millisecond)
		}
	}
}
