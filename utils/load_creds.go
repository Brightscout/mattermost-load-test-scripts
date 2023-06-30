package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/Brightscout/mattermost-load-test-scripts/constants"
	"github.com/Brightscout/mattermost-load-test-scripts/serializers"
)

func LoadCreds() (*serializers.ClientResponse, error) {
	tempStoreFile, err := os.Open(constants.TempStoreFile)
	if err != nil {
		return nil, err
	}

	defer tempStoreFile.Close()
	byteValue, err := ioutil.ReadAll(tempStoreFile)
	if err != nil {
		return nil, err
	}

	var response *serializers.ClientResponse
	if err := json.Unmarshal(byteValue, &response); err != nil {
		return nil, err
	}

	return response, nil
}
