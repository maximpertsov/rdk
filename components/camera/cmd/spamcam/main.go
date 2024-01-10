package main

import (
	"context"

	"go.viam.com/rdk/components/camera"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/robot/client"
	"go.viam.com/utils/rpc"
)

func main() {
	logger := logging.NewDebugLogger("client")
	robot, err := client.New(
		context.Background(),
		"rover2-main.tkofq1ck33.viam.cloud",
		logger,
		client.WithDialOptions(rpc.WithEntityCredentials(
			"287cc05f-f383-4991-b221-79422cc851ce",
			rpc.Credentials{
				Type:    rpc.CredentialsTypeAPIKey,
				Payload: "49991pvan6y5qrlf1jdb556afvzqfsas",
			})),
	)
	if err != nil {
		logger.Fatal(err)
	}

	defer robot.Close(context.Background())
	logger.Info("Resources:")
	logger.Info(robot.ResourceNames())

	for {
		camComponent, err := camera.FromRobot(robot, "cam")
		if err != nil {
			logger.Error(err)
			return
		}
		camReturnValue, _, err := camComponent.Images(context.Background())
		if err != nil {
			camReturnValue[0].Image.ColorModel()
			logger.Error(err)
			return
		}
		logger.Infof("cam Properties return value: %+v", camReturnValue)
	}
}
