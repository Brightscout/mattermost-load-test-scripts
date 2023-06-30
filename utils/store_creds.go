package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/Brightscout/mattermost-load-test-scripts/constants"
	"github.com/Brightscout/mattermost-load-test-scripts/serializers"
)

func StoreCreds(response *serializers.ClientResponse) error {
	responseBytes, err := json.Marshal(response)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(constants.TempStoreFile, responseBytes, os.ModePerm); err != nil {
		return err
	}

	return nil
}
