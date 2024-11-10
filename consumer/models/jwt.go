package models

type JWTAuth struct {
	Key                 string                 `json:"key,omitempty" validate:"required,alphanum,min=20,max=200"`         // Unique Key for the Consumer
	Secret              string                 `json:"secret,omitempty" validate:"omitempty,alphanum,min=10,max=200"`     // The encryption key. If unspecified, auto generated in the background. This field supports saving the value in Secret Manager using the APISIX Secret resource.
	PublicKey           string                 `json:"public_key,omitempty" validate:"omitempty,alphanum,min=10,max=200"` // Required if RS256 or ES256 is set for the algorithm attribute. RSA or ECDSA public key. This field supports saving the value in Secret Manager using the APISIX Secret resource.
	Algorithm           JWTEncryptionAlgorithm `json:"algorithm,omitempty" validate:"omitempty"`                          // Default is HS256 - ["HS256", "HS512", "RS256", "ES256"] - Encryption algorithm.
	Exp                 int                    `json:"exp,omitempty" validate:"omitempty,number"`                         // Default - 86400 - Expiry time of the token in seconds.
	Base64Secret        bool                   `json:"base64_secret,omitempty" validate:"omitempty,boolean"`              // Default - False - Set to true if the secret is base64 encoded.
	LifetimeGracePeriod int                    `json:"lifetime_grace_period,omitempty" validate:"omitempty,number"`       // Default - Don't touch unless you know what you're doing - 0 - Define the leeway in seconds to account for clock skew between the server that generated the jwt and the server validating it. Value should be zero (0) or a positive integer.
	Meta                *Meta                  `json:"_meta,omitempty" validate:"omitempty"`
}

type JWTEncryptionAlgorithm string

const (
	HS256 JWTEncryptionAlgorithm = "HS256"
	HS512 JWTEncryptionAlgorithm = "HS512"
	RS256 JWTEncryptionAlgorithm = "RS256"
	ES256 JWTEncryptionAlgorithm = "ES256"
)
