package models

type All[V Value] struct {
	Total   int          `json:"total" validate:"omitempty,number"`
	Objects *[]Object[V] `json:"list"`
}

type Object[V Value] struct {
	Key           string `json:"key" validate:"required,uri"`
	ModifiedIndex int    `json:"modifiedIndex,omitempty" validate:"omitempty,number"`
	CreatedIndex  int    `json:"createdIndex,omitempty" validate:"omitempty,number"`
	Value         *V     `json:"value,omitempty"`
}

type Value interface {
	Consumer | ConsumerGroup | Credential | GlobalRule | PluginConfig | Proto | Route | Secret | Service | SSL | Stream | Upstream
	GetID() string
	GetName() string
	GetTTL() string
	GetURI() string
	GetURIs() []string
	GetUsername() string
}

type DeleteResponse struct {
	Key     string `json:"key" validate:"required,ascii,min=1,max=500"`
	Deleted string `json:"deleted" validate:"required,ascii,min=1,max=500"`
}

type Timeout struct {
	Connect int `json:"connect,omitempty" validate:"omitempty,number"`
	Send    int `json:"send,omitempty" validate:"omitempty,number"`
	Read    int `json:"read,omitempty" validate:"omitempty,number"`
}

type Kind interface {
	Get | Create | Update | Delete
}

type Get struct {
}
type Create struct {
}
type Update struct {
}
type Delete struct {
}
