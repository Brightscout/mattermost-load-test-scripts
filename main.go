package main

import (
	"errors"
	"os"

	"github.com/Brightscout/mattermost-load-test-scripts/scripts"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	config, err := scripts.LoadConfig()
	if err != nil {
		logger.Error("failed to load the config",
			zap.Error(err),
		)

		return
	}

	args := os.Args
	if len(args) > 1 {
		switch args[1] {
		case "create_users":
			err = scripts.CreateUsers(config, logger)
		case "clear_store":
			err = scripts.ClearStore()
		default:
			err = errors.New("invalid arguments")
		}
	}

	if err != nil {
		logger.Error("failed to run the script",
			zap.String("arg", args[1]),
			zap.Error(err),
		)
	}
}
