package scripts

import (
	"fmt"

	"github.com/mattermost/mattermost-server/v6/model"
	"go.uber.org/zap"

	"github.com/Brightscout/mattermost-load-test-scripts/serializers"
	"github.com/Brightscout/mattermost-load-test-scripts/utils"
)

func CreateChannels(config *serializers.Config, logger *zap.Logger) error {
	client := model.NewAPIv4Client(config.ConnectionConfiguration.ServerURL)
	if _, _, err := client.Login(config.ConnectionConfiguration.AdminEmail, config.ConnectionConfiguration.AdminPassword); err != nil {
		return err
	}

	var newChannels []*serializers.ChannelResponse
	response, err := utils.LoadResponse()
	if err != nil {
		return err
	}

	for _, channel := range config.ChannelsConfiguration {
		team, _, err := client.GetTeamByName(channel.MMTeamName, "")
		if err != nil {
			logger.Error("unable to get the team details",
				zap.String("TeamName", channel.MMTeamName),
				zap.Error(err),
			)
			continue
		}

		createdChannel, _, err := client.CreateChannel(&model.Channel{
			TeamId:      team.Id,
			Name:        channel.Name,
			DisplayName: channel.DisplayName,
			Type:        model.ChannelType(channel.Type),
		})

		if err != nil {
			logger.Error("unable to create the channel",
				zap.String("ChannelName", channel.Name),
				zap.Error(err),
			)
			continue
		}

		newChannels = append(newChannels, &serializers.ChannelResponse{
			ID: createdChannel.Id,
		})

		newUserIDs := []string{}
		for _, user := range response.UserResponse {
			newUserIDs = append(newUserIDs, user.ID)
		}

		if _, _, err := client.AddTeamMembers(team.Id, newUserIDs); err != nil {
			logger.Error("unable to add users to the team",
				zap.String("TeamName", channel.MMTeamName),
				zap.Error(err),
			)
			continue
		}

		channelLinkCommand := fmt.Sprintf("/msteams-sync link %s %s", channel.MSTeamsTeamID, channel.MSTeamsChannelID)
		if _, _, err := client.ExecuteCommand(createdChannel.Id, channelLinkCommand); err != nil {
			logger.Error("unable to execute the command to link the channel",
				zap.Error(err),
			)
			continue
		}

	}

	response.ChannelResponse = newChannels
	if err := utils.StoreResponse(response); err != nil {
		return err
	}

	return nil
}
