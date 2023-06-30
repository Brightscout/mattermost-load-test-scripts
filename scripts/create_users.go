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
			logger.Info("unable to create new user", zap.String("username", user.Username), zap.Error(err))
			continue
		}

		newUsers = append(newUsers, &serializers.UserResponse{
			ID:       createdUser.Id,
			Username: createdUser.Username,
			Email:    createdUser.Email,
		})
	}

	response, err := utils.LoadResponse()
	if err != nil {
		return err
	}

	response.UserResponse = newUsers
	if err := utils.StoreResponse(response); err != nil {
		return err
	}

	return nil
}
