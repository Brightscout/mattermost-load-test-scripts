package scripts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Brightscout/mattermost-load-test-scripts/serializers"
)

// TODO: remove later, if not needed
func LoadConfig() {
	configFile, err := os.Open("config/config.json")
	if err != nil {
		panic(err)
	}

	defer configFile.Close()
	byteValue, err := ioutil.ReadAll(configFile)
	if err != nil {
		panic(err)
	}

	var config serializers.Config
	if err := json.Unmarshal(byteValue, &config); err != nil {
		panic(err)
	}

	fmt.Println(config.ConnectionConfiguration.ServerURL)
}
