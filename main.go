package main

import (
	"errors"
	"os"

	"github.com/Brightscout/mattermost-load-test-scripts/constants"
	"github.com/Brightscout/mattermost-load-test-scripts/scripts"
	"github.com/Brightscout/mattermost-load-test-scripts/utils"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	defer logger.Sync()
	config, err := utils.LoadConfig()
	if err != nil {
		logger.Error("failed to load the config", zap.Error(err))
		return
	}

	args := os.Args
	if len(args) > 1 {
		switch args[1] {
		case constants.CreateUsers:
			err = scripts.CreateUsers(config, logger)
		case constants.ClearStore:
			err = scripts.ClearStore()
		case constants.CreateChannels:
			err = scripts.CreateChannels(config, logger)
		case constants.CreateDMAndGM:
			err = scripts.CreateDMAndGM(config, logger)
		default:
			err = errors.New("invalid arguments")
		}
	}

	if err != nil {
		logger.Error("failed to run the script", zap.String("arg", args[1]), zap.Error(err))
	}
}
