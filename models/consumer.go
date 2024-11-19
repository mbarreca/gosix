package models

type Consumer struct {
	Username   string            `json:"username" validate:"required,alphanum,min=1,max=200"`
	Desc       string            `json:"desc,omitempty" validate:"omitempty,ascii,min=1,max=2000"`
	GroupID    string            `json:"group_id,omitempty" validate:"omitempty,ascii,min=1,max=2000"`
	Labels     map[string]string `json:"labels,omitempty" validate:"omitempty,ascii,min=1,max=2000"`
	Plugins    *Plugins          `json:"plugins,omitempty" validate:"omitempty"`
	CreateTime int               `json:"create_time,omitempty" validate:"omitempty,number"`
	UpdateTime int               `json:"update_time,omitempty" validate:"omitempty,number"`
}

func (c Consumer) GetID() string {
	return ""
}
func (c Consumer) GetName() string {
	return ""
}
func (c Consumer) GetTTL() string {
	return ""
}
func (c Consumer) GetURI() string {
	return ""
}
func (c Consumer) GetURIs() []string {
	return nil
}
func (c Consumer) GetUsername() string {
	return c.Username
}
