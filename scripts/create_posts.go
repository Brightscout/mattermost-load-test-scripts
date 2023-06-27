package scripts

import (
	"math/rand"
	"sync"

	"github.com/mattermost/mattermost-server/v6/model"

	"github.com/Brightscout/mattermost-load-test-scripts/serializers"
	"github.com/Brightscout/mattermost-load-test-scripts/utils"
)

func CreatePosts(config *serializers.Config) error {
	client := model.NewAPIv4Client(config.ConnectionConfiguration.ServerURL)
	response, err := utils.LoadResponse()
	if err != nil {
		return err
	}

	channelIDs := []string{}
	for _, channel := range response.ChannelResponse {
		channelIDs = append(channelIDs, channel.ID)
	}

	userIDs := []string{}
	for _, user := range response.UserResponse {
		userIDs = append(userIDs, user.ID)
	}

	channelIDs = append(channelIDs, response.DMResponse.ID, response.GMResponse.ID)
	var wg = &sync.WaitGroup{}
	for count := 0; count < config.PostsConfiguration.Count; count++ {
		wg.Add(1)
		go func(count int) {
			defer wg.Done()
			channelID := channelIDs[rand.Intn(len(channelIDs))]
			userID := userIDs[rand.Intn(len(userIDs))]
			client.Login(utils.GetUserNameAndPasswordByID(userID, response.UserResponse, config.UsersConfiguration))
			message := utils.GetRandomMessage(5)
			client.CreatePost(&model.Post{
				ChannelId: channelID,
				Message:   message,
			})
		}(count)
	}

	wg.Wait()

	return nil
}
