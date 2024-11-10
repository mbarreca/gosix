package models

type KeyAuth struct {
	Key  string `json:"key,omitempty" validate:"required,ascii,min=20,max=200"` // Unique key for a Consumer. This field supports saving the value in Secret Manager using the APISIX Secret resource.
	Meta *Meta  `json:"_meta,omitempty" validate:"omitempty"`
}
