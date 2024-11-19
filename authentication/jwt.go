package authentication

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5" // MIT
	"github.com/mbarreca/gosix/client"
	"github.com/mbarreca/gosix/consumer"
	"github.com/mbarreca/gosix/models"
)

type JWT struct {
	client *client.Client
	c      *consumer.Consumer
}

// Constructor - You *shouldn't* be using this
func NewJWT(c *client.Client, consumer *consumer.Consumer) *JWT {
	return &JWT{client: c, c: consumer}
}

// This method will add JWT if it doesn't exist or if it does, generate a new key
// It is highly recommended you store the secret-key using APISix's Secret Methods
// Which is Environment Variables, GCP, AWS or Hashicorps Vault
// username -> Consumers username
// key -> The JWT key to generate the token with
func (j *JWT) Get(username string, key string) (string, error) {
	user, err := j.c.Get(username)
	if err != nil {
		return "", err
	}
	// Create EXP Time
	exp := 86400
	if os.Getenv("GOSIX_APISIX_PLUGIN_JWT_EXP") != "" {
		exp, err = strconv.Atoi(os.Getenv("GOSIX_APISIX_PLUGIN_JWT_EXP"))
		if err != nil {
			return "", err
		}
	}
	expTime := time.Now().UTC().Add(time.Second * time.Duration(exp))
	var jwtAuth *models.JwtAuth
	if user.Plugins != nil && user.Plugins.JwtAuth != nil {
		// This means JWT Auth is already added, generate a new key and return
		token, err := generateJwt(key, expTime)
		if err != nil {
			return "", err
		}
		return token, nil
	} else {
		// This means we haven't added JWT yet, let's add it
		// You can't have Key and JWT Auth in APISix, prevent this
		if user.Plugins != nil && user.Plugins.KeyAuth != nil {
			return "", errors.New("You can't have JWT and Key Auth on the same consumer")
		}
		// Get an new key object with a new key
		jwtAuth, err = createJwtObject(key)
		if err != nil {
			return "", err
		}
	}
	// If there are no plugins, add them
	if user.Plugins == nil {
		user.Plugins = new(models.Plugins)
	}
	// Re-create the modified consumer
	jwtAuth.Exp = expTime.Unix()
	user.Plugins.JwtAuth = jwtAuth
	if err := j.c.Update(user); err != nil {
		return "", err
	}
	// Generate JWT and return
	token, err := generateJwt(key, expTime)
	if err != nil {
		return "", err
	}
	return token, nil
}

// Delete the JWT Plugin from the Consumer
// username -> Consumers username
func (j *JWT) Delete(username string) error {
	user, err := j.c.Get(username)
	if err != nil {
		return err
	}
	if user.Plugins == nil || user.Plugins.JwtAuth == nil {
		return errors.New("No Key Auth on this consumer")
	}
	user.Plugins.JwtAuth = nil
	if err := j.c.Update(user); err != nil {
		return err
	}
	return nil
}

// Set the enabled/disabled sate of Basic Auth for this consumer
// enabled - True if enabled, false if disabled - default state is Enabled
// username -> Consumers username
func (j *JWT) Enabled(enabled bool, username string) error {
	// Get the consumer
	user, err := j.c.Get(username)
	if err != nil {
		return err
	}
	if user.Plugins != nil && user.Plugins.JwtAuth != nil {
		// Check to see if Meta Exists
		if user.Plugins.JwtAuth.Meta == nil {
			user.Plugins.JwtAuth.Meta = new(models.Meta)
		}
		// Disable the key
		user.Plugins.JwtAuth.Meta.Disable = !enabled
	} else {
		return errors.New("User doesn't have a Basic Auth plugin")
	}
	// Update the consumer
	if err := j.c.Update(user); err != nil {
		return err
	}
	return nil

}

func createJwtObject(key string) (*models.JwtAuth, error) {

	// Validate ENV Variables
	secret, algo, err := validateSecretAndAlgorithm()
	if err != nil {
		return nil, err
	}
	// Build the Key Auth Object, make a new key and update
	var jwtAuth models.JwtAuth
	jwtAuth.Key = key
	jwtAuth.Algorithm = algo
	jwtAuth.Secret = secret
	return &jwtAuth, nil
}

// Generates a JWT based on Environment Variables and the provided information
func generateJwt(key string, exp time.Time) (string, error) {
	// Setup Secret and Algorithm
	secret, algo, err := validateSecretAndAlgorithm()
	if err != nil {
		return "", err
	}
	// Setup Signing Method
	var method jwt.SigningMethod
	if algo == "HS256" {
		method = jwt.SigningMethodHS256
	} else if algo == "HS512" {
		method = jwt.SigningMethodHS512
	} else if algo == "RS256" {
		method = jwt.SigningMethodRS256
	} else if algo == "ES256" {
		method = jwt.SigningMethodES256
	} else {
		return "", errors.New("Unsupported or Invalid Signing Method")
	}
	token := jwt.NewWithClaims(method, jwt.MapClaims{
		"key": key,
		"exp": exp.Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Generates a JWT based on Environment Variables and the provided information
func validateSecretAndAlgorithm() (string, models.JwtEncryptionAlgorithm, error) {
	secret := os.Getenv("GOSIX_APISIX_PLUGIN_JWT_SECRET")
	if len(secret) < 25 {
		return "", "", errors.New("Jwt Secret is invalid, must be 25 characters+")
	}
	// Make sure algorithm is valid
	algoRaw := os.Getenv("GOSIX_APISIX_PLUGIN_JWT_ALGORITHM")
	if algoRaw != "HS256" && algoRaw != "HS512" && algoRaw != "RS256" && algoRaw != "ES256" {
		return "", "", errors.New("Jwt Algorithm is invalid, must be HS256, HS512, RS256 or ES256")
	}
	return secret, models.JwtEncryptionAlgorithm(algoRaw), nil
}
