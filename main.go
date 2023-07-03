package main

import (
	"errors"
	"os"

	"github.com/Brightscout/mattermost-load-test-scripts/constants"
	"github.com/Brightscout/mattermost-load-test-scripts/scripts"
	"github.com/Brightscout/mattermost-load-test-scripts/utils"
	"go.uber.org/zap"
)

// TODO: Update documentation
func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	config, err := utils.LoadConfig()
	if err != nil {
		logger.Error("failed to load the config", zap.Error(err))
		return
	}

	if err := config.IsConnectionConfigurationValid(); err != nil {
		logger.Error("Error in validating the connection configuration", zap.Error(err))
		return
	}

	args := os.Args
	if len(args) > 1 {
		switch args[1] {
		case constants.CreateUsers:
			if err := config.IsUsersConfigurationValid(); err != nil {
				logger.Error("Error in validating the user configuration", zap.Error(err))
				break
			}

			err = scripts.CreateUsers(config, logger)
		case constants.ClearStore:
			err = scripts.ClearStore()
		case constants.CreateChannels:
			if err := config.IsChannelsConfigurationValid(); err != nil {
				logger.Error("Error in validating the channel configuration", zap.Error(err))
				break
			}

			err = scripts.CreateChannels(config, logger)
		case constants.CreateDMAndGMs:
			err = scripts.CreateDMAndGMs(config, logger)
		default:
			err = errors.New("invalid arguments")
		}
	}
	if err != nil {
		logger.Error("failed to run the script", zap.String("arg", args[1]), zap.Error(err))
	}

	_ = logger.Sync()
}
