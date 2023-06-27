package scripts

import (
	"errors"

	"github.com/mattermost/mattermost-server/v6/model"
	"go.uber.org/zap"

	"github.com/Brightscout/mattermost-load-test-scripts/constants"
	"github.com/Brightscout/mattermost-load-test-scripts/serializers"
	"github.com/Brightscout/mattermost-load-test-scripts/utils"
)

func CreateDMAndGM(config *serializers.Config, logger *zap.Logger) error {
	response, err := utils.LoadResponse()
	if err != nil {
		return err
	}

	if len(response.UserResponse) == 0 {
		return errors.New("no new user created")
	}

	client := model.NewAPIv4Client(config.ConnectionConfiguration.ServerURL)
	if _, _, err := client.Login(utils.GetUserNameAndPasswordByID(response.UserResponse[0].ID, response.UserResponse, config.UsersConfiguration)); err != nil {
		return err
	}

	if len(response.UserResponse) >= constants.MinUsersForDM {
		newDM, _, err := client.CreateDirectChannel(response.UserResponse[0].ID, response.UserResponse[1].ID)
		if err != nil {
			logger.Error("unable to create the DM", zap.Error(err))
		} else {
			response.DMResponse = &serializers.ChannelResponse{
				ID: newDM.Id,
			}
		}
	}

	if len(response.UserResponse) >= constants.MinUsersForGM {
		newGM, _, err := client.CreateGroupChannel([]string{
			response.UserResponse[0].ID,
			response.UserResponse[1].ID,
			response.UserResponse[2].ID,
		})

		if err != nil {
			logger.Error("unable to create the GM", zap.Error(err))
		} else {
			response.GMResponse = &serializers.ChannelResponse{
				ID: newGM.Id,
			}
		}
	}

	if err := utils.StoreResponse(response); err != nil {
		return err
	}

	return nil
}
