package utils

import "github.com/Brightscout/mattermost-load-test-scripts/serializers"

func GetUserNameAndPasswordByID(userID string, usersResponse []*serializers.UserResponse, usersConfig []serializers.UsersConfiguration) (string, string) {
	username := ""
	password := ""
	for _, user := range usersResponse {
		if user.ID == userID {
			username = user.Username
			break
		}
	}

	for _, user := range usersConfig {
		if user.Username == username {
			password = user.Password
			break
		}
	}

	return username, password
}
