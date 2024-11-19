package models

type Credential struct {
	Username   string            `json:"-"` // This is here for purposes of allowing the interfaces to work
	ID         string            `json:"id,omitempty" validate:"omitempty,ascii,min=1,max=64"`
	Desc       string            `json:"desc,omitempty" validate:"omitempty,ascii,min=1,max=2000"`
	Labels     map[string]string `json:"labels,omitempty" validate:"omitempty,ascii,min=1,max=2000"`
	Plugins    *Plugins          `json:"plugins,omitempty" validate:"omitempty"`
	CreateTime int               `json:"create_time,omitempty" validate:"omitempty,number"`
	UpdateTime int               `json:"update_time,omitempty" validate:"omitempty,number"`
}

func (c Credential) GetID() string {
	return c.ID
}
func (c Credential) GetName() string {
	return ""
}
func (c Credential) GetTTL() string {
	return ""
}
func (c Credential) GetURI() string {
	return ""
}
func (c Credential) GetURIs() []string {
	return nil
}
func (c Credential) GetUsername() string {
	return c.Username
}
