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

// Validation errors
const (
	ErrorEmptyServerURL          = "server URL should not be empty"
	ErrorEmptyAdminEmail         = "admin email should not be empty"
	ErrorEmptyAdminPassword      = "admin password should not be empty"
	ErrorEmptyUsername           = "username should not be empty"
	ErrorEmptyUserPassword       = "user password should not be empty"
	ErrorEmptyUserEmail          = "user email should not be empty"
	ErrorEmptyChannelDisplayName = "channel display name should not be empty"
	ErrorEmptyChannelSlugName    = "channel slug name should not be empty"
	ErrorEmptyChannelType        = "channel type should not be empty"
	ErrorEmptyMMTeamName         = "mattermost team name should not be empty"
	ErrorEmptyMSTeamsTeamID      = "ms teams team ID should not be empty"
	ErrorEmptyMSTeamsChannelID   = "ms teams channel ID should not be empty"
	ErrorInvalidChannelType      = "invalid channel type"
)
