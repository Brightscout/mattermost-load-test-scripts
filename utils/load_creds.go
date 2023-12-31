package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Brightscout/mattermost-load-test-scripts/constants"
	"github.com/Brightscout/mattermost-load-test-scripts/serializers"
)

func LoadCreds() (*serializers.ClientResponse, error) {
	tempStoreFile, err := os.Open(constants.TempStoreFile)
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			return &serializers.ClientResponse{}, nil
		}

		return nil, err
	}

	defer tempStoreFile.Close()
	byteValue, err := ioutil.ReadAll(tempStoreFile)
	if err != nil {
		return nil, err
	}

	if len(byteValue) == 0 {
		return &serializers.ClientResponse{}, nil
	}

	var response *serializers.ClientResponse
	if err := json.Unmarshal(byteValue, &response); err != nil {
		return nil, err
	}

	return response, nil
}
