package main

import (
	"context"
	"time"

	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/robot/client"
	"go.viam.com/utils/rpc"
)

func connectWebRTC(logger logging.Logger) (*client.RobotClient, error) {
	return client.New(
		context.Background(),
		"something-unique",
		logger,
		client.WithDialOptions(rpc.WithInsecure()),
	)
}

func main() {
	logger := logging.NewDebugLogger("client")
	robot, err := connectWebRTC(logger)

	if err != nil {
		logger.Fatal(err)
	}

	defer robot.Close(context.Background())
	for {
		rs := robot.ResourceNames()
		logger.Infof("Resource count: %d", len(rs))
		time.Sleep(2 * time.Second)
	}
}
