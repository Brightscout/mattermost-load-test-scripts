package scripts

import (
	"github.com/mattermost/mattermost-server/v6/model"

	"github.com/Brightscout/mattermost-load-test-scripts/serializers"
	"github.com/Brightscout/mattermost-load-test-scripts/utils"
)

func CreateDMAndGM(config *serializers.Config) error {
	client := model.NewAPIv4Client(config.ConnectionConfiguration.ServerURL)
	if _, _, err := client.Login(config.ConnectionConfiguration.AdminEmail, config.ConnectionConfiguration.AdminPassword); err != nil {
		return err
	}

	response, err := utils.LoadResponse()
	if err != nil {
		return err
	}

	if len(response.UserResponse) > 1 {
		newDM, _, err := client.CreateDirectChannel(response.UserResponse[0].ID, response.UserResponse[1].ID)
		if err != nil {
			return err
		}

		response.DMResponse = &serializers.ChannelResponse{
			ID: newDM.Id,
		}
	}

	if len(response.UserResponse) > 2 {
		newGM, _, err := client.CreateGroupChannel([]string{
			response.UserResponse[0].ID,
			response.UserResponse[1].ID,
			response.UserResponse[2].ID,
		})

		if err != nil {
			return err
		}

		response.GMResponse = &serializers.ChannelResponse{
			ID: newGM.Id,
		}
	}

	if err := utils.StoreResponse(response); err != nil {
		return err
	}

	return nil
}
