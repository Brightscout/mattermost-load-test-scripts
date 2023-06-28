package scripts

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/mattermost/mattermost-server/v6/model"
	"go.uber.org/zap"

	"github.com/Brightscout/mattermost-load-test-scripts/constants"
	"github.com/Brightscout/mattermost-load-test-scripts/serializers"
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
			logger.Info("unable to create new user",
				zap.String("username", user.Username),
				zap.Error(err),
			)
			continue
		}

		newUsers = append(newUsers, &serializers.UserResponse{
			ID:       createdUser.Id,
			Username: createdUser.Username,
			Email:    createdUser.Email,
		})
	}

	userMap := make(map[string]interface{})
	userMap[constants.NewUsersKey] = newUsers
	userMapBytes, err := json.Marshal(userMap)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(constants.TempStoreFile, userMapBytes, os.ModePerm); err != nil {
		return err
	}

	return nil
}
