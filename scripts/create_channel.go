package scripts

import (
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-server/v6/model"
	"go.uber.org/zap"

	"github.com/Brightscout/mattermost-load-test-scripts/serializers"
	"github.com/Brightscout/mattermost-load-test-scripts/utils"
)

func CreateChannels(config *serializers.Config, logger *zap.Logger) error {
	connectionConfiguration := config.ConnectionConfiguration
	client := model.NewAPIv4Client(connectionConfiguration.ServerURL)
	if _, _, err := client.Login(connectionConfiguration.AdminEmail, connectionConfiguration.AdminPassword); err != nil {
		return err
	}

	var newChannels []*serializers.ChannelResponse
	response, err := utils.LoadCreds()
	if err != nil {
		return err
	}

	for _, channel := range config.ChannelsConfiguration {
		team, err := CreateOrGetTeam(client, channel.MMTeamName)
		if err != nil {
			logger.Error("unable to get the team details", zap.String("TeamName", channel.MMTeamName), zap.Error(err))
			continue
		}

		createdChannel, err := CreateOrGetChannel(client, team, channel)
		if err != nil {
			logger.Error("unable to create the channel", zap.String("ChannelName", channel.Name), zap.Error(err))
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
			logger.Error("unable to add users to the team", zap.String("TeamName", channel.MMTeamName), zap.Error(err))
			continue
		}

		channelLinkCommand := fmt.Sprintf("/msteams-sync link %s %s", channel.MSTeamsTeamID, channel.MSTeamsChannelID)
		if _, _, err := client.ExecuteCommand(createdChannel.Id, channelLinkCommand); err != nil {
			logger.Error("unable to execute the command to link the channel", zap.Error(err))
			continue
		}

	}

	response.ChannelResponse = newChannels
	if err := utils.StoreCreds(response); err != nil {
		return err
	}

	return nil
}

func CreateOrGetTeam(client *model.Client4, teamName string) (*model.Team, error) {
	team, response, err := client.GetTeamByName(teamName, "")
	if err != nil {
		if response.StatusCode == http.StatusNotFound {
			newTeam, _, cErr := client.CreateTeam(&model.Team{
				Name:        teamName,
				DisplayName: teamName,
				Type:        model.TeamOpen,
			})
			if cErr != nil {
				return nil, cErr
			}

			return newTeam, nil
		}

		return nil, err
	}

	return team, nil
}

func CreateOrGetChannel(client *model.Client4, team *model.Team, channel serializers.ChannelsConfiguration) (*model.Channel, error) {
	createdChannel, response, err := client.CreateChannel(&model.Channel{
		TeamId:      team.Id,
		Name:        channel.Name,
		DisplayName: channel.DisplayName,
		Type:        model.ChannelType(channel.Type),
	})
	if err != nil {
		if response.StatusCode == http.StatusBadRequest {
			channel, _, gErr := client.GetChannelByName(channel.Name, team.Id, "")
			if gErr != nil {
				return nil, gErr
			}

			return channel, nil
		}

		return nil, err
	}

	return createdChannel, nil
}
