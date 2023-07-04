# Mattermost load test scripts

Mattermost load-test-scripts provides a set of scripts written in [Go](https://golang.org/) to help profiling [Mattermost](https://github.com/mattermost/mattermost-server) under heavy load, simulating real-world usage of a server installation at scale.

## Setup

Make sure you have the following components installed:  

- Go - v1.18 - [Getting Started](https://golang.org/doc/install)
    > **Note:** If you have installed Go to a custom location, make sure the `$GOROOT` variable is set properly. Refer [Installing to a custom location](https://golang.org/doc/install#install).

- Install the K6 load testing tool from [here](https://k6.io/docs/get-started/installation).

- Clone the repo using the command:
    ```
    git clone git@github.com:Brightscout/mattermost-load-test-scripts.git
    ``` 
    or 
    ```
    git clone https://github.com/Brightscout/mattermost-load-test-scripts.git
    ```

## How to use
- Create a `config.json` file.
    - Run command to copy the sample `config.sample.json` file.
    ```
    cp config/config.sample.json config/config.json
    ```
    - Configure the `config.json` file created according to the load to be tested.
    - Go to config [docs](docs/config.md) to check details on the config settings.

- Run command `make build` to create a new binary file for the load test scripts.

- Run command `make create_users` to create new users with the details present in the config file.

- Connect the system admin account with an MS Teams account to link the new channels with the MS Teams channels.

- Run command `make create_channels` to create the new channels and add the new users to them. Running this command also links the Mattermost channels with the MS Teams channels present in the config file.

- Run command `make create_dm_and_gm` to create the DMs and GMs between the new users.

- Login and connect all the new users with their respective MS Teams accounts to enable the relaying of messages from Mattermost to MS Teams on behalf of these users.

- Run command `make create_posts` to create the random posts in the Mattermost channels, DMs, and GMs.

- Run command `make clear_store` to clear all the stored data present in the temporary database to start load testing with new details.
