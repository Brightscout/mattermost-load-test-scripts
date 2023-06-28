package scripts

import (
	"errors"
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

	if response.ChannelResponse == nil {
		return errors.New("no new channels present to create posts")
	}

	channelIDs := []string{}
	for _, channel := range response.ChannelResponse {
		channelIDs = append(channelIDs, channel.ID)
	}

	if response.UserResponse == nil {
		return errors.New("no new users present to create posts")
	}

	userIDs := []string{}
	for _, user := range response.UserResponse {
		userIDs = append(userIDs, user.ID)
	}

	if response.DMResponse != nil && response.DMResponse.ID != "" {
		channelIDs = append(channelIDs, response.DMResponse.ID)
	}

	if response.GMResponse != nil && response.GMResponse.ID != "" {
		channelIDs = append(channelIDs, response.GMResponse.ID)
	}

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
