package models

type ConsumerRequest struct {
	Username string   `json:"username" validate:"required,alphanum,min=1,max=100"`        // Name of the Consumer.
	Desc     string   `json:"desc,omitempty" validate:"omitempty,alphanum,min=1,max=300"` // Description of usage scenarios.
	Plugins  *Plugins `json:"plugins,omitempty" validate:"omitempty"`                     // Plugins associated with the consumer
}

type GetResponse struct {
	List  *[]ConsumerList `json:"list" validate:"required"`
	Total int             `json:"total" validate:"required,number"`
}

type GetByUsernameResponse struct {
	Key           string         `json:"key" validate:"required,ascii,min=1,max=300"`
	Value         *ConsumerValue `json:"value" validate:"required"`
	CreatedIndex  int            `json:"createdIndex" validate:"required,number"`
	ModifiedIndex int            `json:"modifiedIndex" validate:"required,number"`
}

type PutResponse struct {
	Key   string         `json:"key" validate:"required,ascii,min=1,max=300"`
	Value *ConsumerValue `json:"value" validate:"required"`
}

type DeleteResponse struct {
	Key     string `json:"key" validate:"required,ascii,min=1,max=300"`
	Deleted string `json:"deleted" validate:"required,alphanum,min=1,max=300"`
}

/*
Sub-Models
*/
type Plugins struct {
	BasicAuth *BasicAuth `json:"basic-auth,omitempty" validate:"omitempty"`
	JwtAuth   *JWTAuth   `json:"jwt-auth,omitempty" validate:"omitempty"`
	KeyAuth   *KeyAuth   `json:"key-auth,omitempty" validate:"omitempty"`
}
type ConsumerValue struct {
	CreateTime int      `json:"create_time" validate:"required"`
	UpdateTime int      `json:"update_time" validate:"required"`
	Username   string   `json:"username" validate:"required,alphanum,min=1,max=200"`
	Desc       string   `json:"desc,omitempty" validate:"omitempty,alphanum,min=1,max=300"`
	Plugins    *Plugins `json:"plugins,omitempty" validate:"omitempty"`
}
type ConsumerList struct {
	Key           string         `json:"key" validate:"required,ascii,min=1,max=300"`
	Value         *ConsumerValue `json:"value,omitempty" validate:"omitempty"`
	CreatedIndex  int            `json:"createdIndex,omitempty" validate:"required,number"`
	ModifiedIndex int            `json:"modifiedIndex,omitempty" validate:"required,number"`
}
