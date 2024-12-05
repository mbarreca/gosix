package models

type BasicAuth struct {
	Username string `json:"username,omitempty" validate:"required,ascii,min=1,max=200"` // This library uses emails as usernames - Must be Unique
	Password string `json:"password,omitempty" validate:"required,ascii,min=1,max=200"` // Password of the user. This field supports saving the value in Secret Manager using the APISIX Secret resource.
	Meta     *Meta  `json:"_meta,omitempty" validate:"omitempty"`
}

type JwtAuth struct {
	Key                 string                 `json:"key,omitempty" validate:"required,ascii,min=20,max=200"`         // Unique Key for the Consumer
	Secret              string                 `json:"secret,omitempty" validate:"omitempty,ascii,min=10,max=200"`     // HIGHLY Recommended to use a secret in this instance - The encryption key. If unspecified, auto generated in the background. This field supports saving the value in Secret Manager using the APISIX Secret resource.
	PublicKey           string                 `json:"public_key,omitempty" validate:"omitempty,ascii,min=10,max=200"` // Required if RS256 or ES256 is set for the algorithm attribute. RSA or ECDSA public key. This field supports saving the value in Secret Manager using the APISIX Secret resource.
	Algorithm           JwtEncryptionAlgorithm `json:"algorithm,omitempty" validate:"omitempty"`                       // Default is HS256 - ["HS256", "HS512", "RS256", "ES256"] - Encryption algorithm.
	Exp                 int64                  `json:"exp,omitempty" validate:"omitempty,number"`                      // Default - 86400 - Expiry time of the token in seconds.
	Base64Secret        bool                   `json:"base64_secret,omitempty" validate:"omitempty,boolean"`           // Default - False - Set to true if the secret is base64 encoded.
	LifetimeGracePeriod int                    `json:"lifetime_grace_period,omitempty" validate:"omitempty,number"`    // Default - Don't touch unless you know what you're doing - 0 - Define the leeway in seconds to account for clock skew between the server that generated the jwt and the server validating it. Value should be zero (0) or a positive integer.
	Meta                *Meta                  `json:"_meta,omitempty" validate:"omitempty"`
}

type JwtEncryptionAlgorithm string

const (
	HS256 JwtEncryptionAlgorithm = "HS256"
	HS512 JwtEncryptionAlgorithm = "HS512"
	RS256 JwtEncryptionAlgorithm = "RS256"
	ES256 JwtEncryptionAlgorithm = "ES256"
)

type KeyAuth struct {
	Key  string `json:"key,omitempty" validate:"required,ascii,min=8,max=200"` // Unique key for a Consumer. This field supports saving the value in Secret Manager using the APISIX Secret resource.
	Exp  string `json:"exp,omitempty" validate:"omitempty,ascii,min=15,max=40"`
	Meta *Meta  `json:"_meta,omitempty" validate:"omitempty"`
}

type Meta struct {
	Disable bool `json:"disable,omitempty" validate:"omitempty,boolean"`
}
