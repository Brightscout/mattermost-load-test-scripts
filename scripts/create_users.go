package scripts

import (
	"github.com/mattermost/mattermost-server/v6/model"
	"go.uber.org/zap"

	"github.com/Brightscout/mattermost-load-test-scripts/serializers"
	"github.com/Brightscout/mattermost-load-test-scripts/utils"
)

func CreateUsers(config *serializers.Config, logger *zap.Logger) error {
	client := model.NewAPIv4Client(config.ConnectionConfiguration.ServerURL)
	var newUsers []*serializers.UserResponse
	for _, user := range config.UsersConfiguration {
		createdUser, _, err := client.CreateUser(&model.User{
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
		})
		if err != nil {
			logger.Info("Unable to create new user", zap.String("Username", user.Username), zap.Error(err))
			continue
		}

		_, userResponse, err := client.Login(user.Username, user.Password)
		if err != nil {
			logger.Info("Unable to login new user", zap.String("Username", user.Username), zap.Error(err))
			continue
		}

		newUsers = append(newUsers, &serializers.UserResponse{
			ID:       createdUser.Id,
			Email:    createdUser.Email,
			Token:    userResponse.Header.Get(model.HeaderToken),
		})
	}

	response, err := utils.LoadCreds()
	if err != nil {
		return err
	}

	response.UserResponse = newUsers
	if err := utils.StoreCreds(response); err != nil {
		return err
	}

	return nil
}
