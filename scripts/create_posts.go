package scripts

import (
	"math/rand"
	"sync"

	"github.com/mattermost/mattermost-server/v6/model"
	"go.uber.org/zap"

	"github.com/Brightscout/mattermost-load-test-scripts/serializers"
	"github.com/Brightscout/mattermost-load-test-scripts/utils"
)

func CreatePosts(config *serializers.Config, logger *zap.Logger) error {
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
			if _, _, err := client.Login(utils.GetUserNameAndPasswordByID(userID, response.UserResponse, config.UsersConfiguration)); err != nil {
				logger.Info("Unable to login",
					zap.String("UserID", userID),
					zap.Error(err),
				)
				return
			}

			message := utils.GetRandomMessage(5)
			if _, _, err := client.CreatePost(&model.Post{
				ChannelId: channelID,
				Message:   message,
			}); err != nil {
				logger.Info("Unable to create the post",
					zap.String("ChannelID", channelID),
					zap.Error(err),
				)
			}
		}(count)
	}

	wg.Wait()
	return nil
}
