package models

type ConsumerGroup struct {
	ID         string            `json:"id,omitempty" validate:"omitempty,ascii,min=1,max=64"`
	Desc       string            `json:"desc,omitempty" validate:"omitempty,ascii,min=1,max=2000"`
	Labels     map[string]string `json:"labels,omitempty" validate:"omitempty,ascii,min=1,max=2000"`
	Plugins    *Plugins          `json:"plugins,omitempty" validate:"omitempty"`
	CreateTime int               `json:"create_time,omitempty" validate:"omitempty,number"`
	UpdateTime int               `json:"update_time,omitempty" validate:"omitempty,number"`
}

func (c ConsumerGroup) GetID() string {
	return c.ID
}
func (c ConsumerGroup) GetName() string {
	return ""
}
func (c ConsumerGroup) GetTTL() string {
	return ""
}
func (c ConsumerGroup) GetURI() string {
	return ""
}
func (c ConsumerGroup) GetURIs() []string {
	return nil
}
func (c ConsumerGroup) GetUsername() string {
	return ""
}
