package serializers

type Config struct {
	ConnectionConfiguration ConnectionConfiguration
	UsersConfiguration []UsersConfiguration
	ChannelsConfiguration []ChannelsConfiguration
	PostsConfiguration PostsConfiguration
}

type ConnectionConfiguration struct {
	ServerURL string 
	AdminEmail string
	AdminPassword string 
}

type UsersConfiguration struct {
	Username string 
	Password string 
	Email string
}

type ChannelsConfiguration struct {
	DisplayName string 
	Name string 
	Type string
	TeamId string
}

type PostsConfiguration struct {
	Count int
}
