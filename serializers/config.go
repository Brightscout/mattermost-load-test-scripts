package serializers

import (
	"errors"
	"fmt"
	"strings"

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

	config := c.ConnectionConfiguration
	config.ServerURL = strings.TrimRight(strings.TrimSpace(config.ServerURL), "/")
	config.AdminEmail = strings.TrimSpace(config.AdminEmail)
	config.AdminPassword = strings.TrimSpace(config.AdminPassword)

	return nil
}

func (c *Config) IsUsersConfigurationValid() error {
	for idx, user := range c.UsersConfiguration {
		if user.Username == "" {
			return fmt.Errorf("%s. index: %d", constants.ErrorEmptyUsername, idx)
		}

		if user.Email == "" {
			return fmt.Errorf("%s. index: %d", constants.ErrorEmptyUserEmail, idx)
		}

		if user.Password == "" {
			return fmt.Errorf("%s. index: %d", constants.ErrorEmptyUserPassword, idx)
		}

		user.Email = strings.TrimSpace(user.Email)
		user.Username = strings.TrimSpace(user.Username)
		user.Password = strings.TrimSpace(user.Password)
	}

	return nil
}

func (c *Config) IsChannelsConfigurationValid() error {
	for idx, channel := range c.ChannelsConfiguration {
		if channel.DisplayName == "" {
			return fmt.Errorf("%s. index: %d", constants.ErrorEmptyChannelDisplayName, idx)
		}

		if channel.Name == "" {
			return fmt.Errorf("%s. index: %d", constants.ErrorEmptyChannelSlugName, idx)
		}

		if channel.Type == "" {
			return fmt.Errorf("%s. index: %d", constants.ErrorEmptyChannelType, idx)
		}

		if channel.MMTeamName == "" {
			return fmt.Errorf("%s. index: %d", constants.ErrorEmptyMMTeamName, idx)
		}

		if channel.MSTeamsTeamID == "" {
			return fmt.Errorf("%s. index: %d", constants.ErrorEmptyMSTeamsTeamID, idx)
		}

		if channel.MSTeamsChannelID == "" {
			return fmt.Errorf("%s. index: %d", constants.ErrorEmptyMSTeamsChannelID, idx)
		}

		channel.Name = strings.TrimSpace(channel.Name)
		channel.Type = strings.TrimSpace(channel.Type)
		channel.MSTeamsTeamID = strings.TrimSpace(channel.MSTeamsTeamID)
		channel.MSTeamsChannelID = strings.TrimSpace(channel.MSTeamsChannelID)

		if channel.Type != string(model.ChannelTypePrivate) && channel.Type != string(model.ChannelTypeOpen) {
			return fmt.Errorf("%s. index: %d", constants.ErrorInvalidChannelType, idx)
		}
	}

	return nil
}
