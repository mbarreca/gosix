package models

type Secret struct {
	ID         string `json:"id,omitempty" validate:"omitempty,ascii,min=1,max=64"`
	Token      string `json:"token" validate:"required,ascii,min=1,max=2000"`
	Namespace  string `json:"namespace,omitempty" validate:"omitempty,ascii,min=1,max=2000"`
	URI        string `json:"uri" validate:"required,uri,min=1,max=1000"`
	Prefix     string `json:"prefix" validate:"required,ascii,min=1,max=2000"`
	CreateTime int    `json:"create_time,omitempty" validate:"omitempty,number"`
	UpdateTime int    `json:"update_time,omitempty" validate:"omitempty,number"`
}

func (c Secret) GetID() string {
	return c.ID
}
func (c Secret) GetName() string {
	return ""
}
func (c Secret) GetTTL() string {
	return ""
}
func (c Secret) GetURI() string {
	return ""
}
func (c Secret) GetURIs() []string {
	return nil
}
func (c Secret) GetUsername() string {
	return ""
}
