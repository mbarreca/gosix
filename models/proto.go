package models

type Proto struct {
	ID         string `json:"id,omitempty" validate:"omitempty,ascii,min=1,max=64"`
	Content    string `json:"content,omitempty" validate:"omitempty,min=1,max=10000"`
	CreateTime int    `json:"create_time,omitempty" validate:"omitempty,number"`
	UpdateTime int    `json:"update_time,omitempty" validate:"omitempty,number"`
}

func (c Proto) GetID() string {
	return c.ID
}
func (c Proto) GetName() string {
	return ""
}
func (c Proto) GetTTL() string {
	return ""
}
func (c Proto) GetURI() string {
	return ""
}
func (c Proto) GetURIs() []string {
	return nil
}
func (c Proto) GetUsername() string {
	return ""
}
