package main

import (
	"fmt"
	"os"

	"github.com/Brightscout/mattermost-load-test-scripts/scripts"
)

func main() {
	config, err := scripts.LoadConfig()
	if err != nil {
		panic(err)
	}

	args := os.Args
	if len(args) > 1 {
		switch args[1] {
		case "create_users":
			err = scripts.CreateUsers(config)
		case "clear_store":
			err = scripts.ClearStore()
		default:
			fmt.Println("Invalid arguments")
		}
	}

	if err != nil {
		panic(err)
	}
}
