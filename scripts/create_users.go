package scripts

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/mattermost/mattermost-server/v5/model"

	"github.com/Brightscout/mattermost-load-test-scripts/constants"
	"github.com/Brightscout/mattermost-load-test-scripts/serializers"
)

func CreateUsers(config *serializers.Config) error {
	client := model.NewAPIv4Client(config.ConnectionConfiguration.ServerURL)
	var newUsers []*serializers.UserResponse
	for _, user := range config.UsersConfiguration {
		createdUser, cErr := client.CreateUser(&model.User{
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
		})

		if cErr.Error != nil {
			return errors.New(cErr.Error.Message)
		}

		newUsers = append(newUsers, &serializers.UserResponse{
			Id:       createdUser.Id,
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

	ioutil.WriteFile(constants.TempStoreFile, userMapBytes, 0644)
	return nil
}
