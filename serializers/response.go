package serializers

type ClientResponse struct {
	UserResponse    []*UserResponse
	ChannelResponse []*ChannelResponse
	DMResponse      *ChannelResponse
	GMResponse      *ChannelResponse
}

type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type ChannelResponse struct {
	ID string `json:"id"`
}
