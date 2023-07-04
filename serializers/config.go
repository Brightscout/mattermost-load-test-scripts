package serializers

import (
	"errors"

	"github.com/Brightscout/mattermost-load-test-scripts/constants"
	"github.com/mattermost/mattermost-server/v6/model"
)

type Config struct {
	ConnectionConfiguration ConnectionConfiguration
	UsersConfiguration      []UsersConfiguration
	ChannelsConfiguration   []ChannelsConfiguration
	PostsConfiguration      PostsConfiguration
}

type ConnectionConfiguration struct {
	ServerURL     string
	AdminEmail    string
	AdminPassword string
}

type UsersConfiguration struct {
	Username string
	Password string
	Email    string
}

type ChannelsConfiguration struct {
	DisplayName      string
	Name             string
	Type             string
	MMTeamName       string
	MSTeamsTeamID    string
	MSTeamsChannelID string
}

type PostsConfiguration struct {
	Count         int
	MaxWordsCount int
	MaxWordLength int
	Duration      string
}

func (c *Config) IsConnectionConfigurationValid() error {
	if c.ConnectionConfiguration.ServerURL == "" {
		return errors.New(constants.ErrorEmptyServerURL)
	}

	if c.ConnectionConfiguration.AdminEmail == "" {
		return errors.New(constants.ErrorEmptyAdminEmail)
	}

	if c.ConnectionConfiguration.AdminPassword == "" {
		return errors.New(constants.ErrorEmptyAdminPassword)
	}

	return nil
}

func (c *Config) IsUsersConfigurationValid() error {
	for _, user := range c.UsersConfiguration {
		if user.Username == "" {
			return errors.New(constants.ErrorEmptyUsername)
		}

		if user.Email == "" {
			return errors.New(constants.ErrorEmptyUserEmail)
		}

		if user.Password == "" {
			return errors.New(constants.ErrorEmptyUserPassword)
		}
	}

	return nil
}

func (c *Config) IsChannelsConfigurationValid() error {
	for _, channel := range c.ChannelsConfiguration {
		if channel.DisplayName == "" {
			return errors.New(constants.ErrorEmptyChannelDisplayName)
		}

		if channel.Name == "" {
			return errors.New(constants.ErrorEmptyChannelName)
		}

		if channel.Type == "" {
			return errors.New(constants.ErrorEmptyChannelType)
		}

		if channel.MMTeamName == "" {
			return errors.New(constants.ErrorEmptyMMTeamName)
		}

		if channel.MSTeamsTeamID == "" {
			return errors.New(constants.ErrorEmptyMSTeamsTeamID)
		}

		if channel.MSTeamsChannelID == "" {
			return errors.New(constants.ErrorEmptyMSTeamsChannelID)
		}

		if channel.Type != string(model.ChannelTypePrivate) && channel.Type != string(model.ChannelTypeOpen) {
			return errors.New(constants.ErrorInvalidChannelType)
		}
	}

	return nil
}
