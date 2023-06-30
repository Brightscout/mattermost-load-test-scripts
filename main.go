package main

import (
	"errors"
	"os"

	"github.com/Brightscout/mattermost-load-test-scripts/constants"
	"github.com/Brightscout/mattermost-load-test-scripts/scripts"
	"github.com/Brightscout/mattermost-load-test-scripts/utils"
	"go.uber.org/zap"
)

// TODO: Complete config validation
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

	args := os.Args
	if len(args) > 1 {
		switch args[1] {
		case constants.CreateUsers:
			err = scripts.CreateUsers(config, logger)
		case constants.ClearStore:
			err = scripts.ClearStore()
		case constants.CreateChannels:
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

	if err := logger.Sync(); err != nil {
		panic(err)
	}
}
