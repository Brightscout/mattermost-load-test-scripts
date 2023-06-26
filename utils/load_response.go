package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/Brightscout/mattermost-load-test-scripts/constants"
	"github.com/Brightscout/mattermost-load-test-scripts/serializers"
)

func LoadResponse() (*serializers.ClientResponse, error) {
	responseFile, err := os.Open(constants.TempStoreFile)
	if err != nil {
		return nil, err
	}

	defer responseFile.Close()
	byteValue, err := ioutil.ReadAll(responseFile)
	if err != nil {
		return nil, err
	}

	var response *serializers.ClientResponse
	if err := json.Unmarshal(byteValue, &response); err != nil {
		return nil, err
	}

	return response, nil
}
