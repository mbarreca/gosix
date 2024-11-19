package models

type GlobalRule struct {
	ID         string            `json:"id,omitempty" validate:"omitempty,ascii,min=1,max=64"`
	Labels     map[string]string `json:"labels,omitempty" validate:"omitempty"`
	Plugins    *Plugins          `json:"plugins,omitempty" validate:"omitempty"`
	CreateTime int               `json:"create_time,omitempty" validate:"omitempty,number"`
	UpdateTime int               `json:"update_time,omitempty" validate:"omitempty,number"`
}

func (c GlobalRule) GetID() string {
	return c.ID
}
func (c GlobalRule) GetName() string {
	return ""
}
func (c GlobalRule) GetTTL() string {
	return ""
}
func (c GlobalRule) GetURI() string {
	return ""
}
func (c GlobalRule) GetURIs() []string {
	return nil
}
func (c GlobalRule) GetUsername() string {
	return ""
}
