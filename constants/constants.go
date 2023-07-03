package constants

// File locations
const (
	ConfigFile    = "config/config.json"
	TempStoreFile = "temp_store.json"
)

// Scripts arguments
const (
	CreateUsers    = "create_users"
	ClearStore     = "clear_store"
	CreateChannels = "create_channels"
	CreateDMAndGMs = "create_dm_and_gm"
)

const (
	MinUsersForDM = 2
	MinUsersForGM = 3
)

// Validations errors
const (
	ErrorEmptyServerURL          = "server URL should not be empty"
	ErrorEmptyAdminEmail         = "admin email should not be empty"
	ErrorEmptyAdminPassword      = "admin password should not be empty"
	ErrorEmptyUsername           = "username should not be empty"
	ErrorEmptyUserPassword       = "user password should not be empty"
	ErrorEmptyUserEmailL         = "user email should not be empty"
	ErrorEmptyChannelDisplayName = "channel display name should not be empty"
	ErrorEmptyChannelName        = "channel name should not be empty"
	ErrorEmptyChannelType        = "channel type should not be empty"
	ErrorEmptyMMTeamName         = "mattermost team name should not be empty"
	ErrorEmptyMSTeamsTeamID      = "ms teams team ID not be empty"
	ErrorEmptyMSTeamsChannelID   = "ms teams channel ID should not be empty"
	ErrorInvalidChannelType      = "invalid channel type"
)
