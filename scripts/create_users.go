package scripts

import (
	"net/http"

	"github.com/mattermost/mattermost-server/v6/model"
	"go.uber.org/zap"

	"github.com/Brightscout/mattermost-load-test-scripts/serializers"
	"github.com/Brightscout/mattermost-load-test-scripts/utils"
)

func CreateUsers(config *serializers.Config, logger *zap.Logger) error {
	client := model.NewAPIv4Client(config.ConnectionConfiguration.ServerURL)
	if _, _, err := client.Login(config.ConnectionConfiguration.AdminEmail, config.ConnectionConfiguration.AdminPassword); err != nil {
		return err
	}

	var newUsers []*serializers.UserResponse
	for _, user := range config.UsersConfiguration {
		createdUser, err := GetOrCreateUser(client, &user)
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
			ID:    createdUser.Id,
			Email: createdUser.Email,
			Token: userResponse.Header.Get(model.HeaderToken),
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

func GetOrCreateUser(client *model.Client4, userDetails *serializers.UsersConfiguration) (*model.User, error) {
	user, response, err := client.GetUserByUsername(userDetails.Username, "")
	if err != nil {
		if response.StatusCode == http.StatusNotFound {
			newUser, _, cErr := client.CreateUser(&model.User{
				Username: userDetails.Username,
				Email:    userDetails.Email,
				Password: userDetails.Password,
			})
			if cErr != nil {
				return nil, cErr
			}

			return newUser, nil
		}

		return nil, err
	}

	return user, nil
}
