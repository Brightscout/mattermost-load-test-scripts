package serializers

type ClientResponse struct {
	UserResponse    []*UserResponse
	ChannelResponse []*ChannelResponse
	DMResponse      *ChannelResponse
	GMResponse      *ChannelResponse
}

type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type ChannelResponse struct {
	ID string `json:"id"`
}
