package consumer

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5" // MIT
	"github.com/mbarreca/gosix"
	"github.com/mbarreca/gosix/consumer/models"
)

// This method will add JWT if it doesn't exist or if it does, generate a new key
// It is highly recommended you store the secret-key using APISix's Secret Methods
// Which is Environment Variables, GCP, AWS or Hashicorps Vault
func JWTAuthGetKey(username string, key string, client *gosix.Client) (string, error) {
	// Get the consumer
	origConsumer, err := GetByUsername(username, client)
	if err != nil {
		return "", err
	}
	// Pull relevant fields
	var modConsumer models.ConsumerRequest
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
	// This is safe because Value needs to be included otherwise the validator will throw an error
	plugins := origConsumer.Value.Plugins
	if plugins != nil && plugins.JwtAuth != nil {
		// This means JWT Auth is already added, generate a new key and return
		token, err := generateJwt(key, expTime)
		if err != nil {
			return "", err
		}
		return token, nil
	} else {
		// This means we haven't added JWT yet, let's add it
		// You can't have Key and JWT Auth in APISix, prevent this
		if plugins != nil && plugins.KeyAuth != nil {
			return "", errors.New("You can't have JWT and Key Auth on the same consumer")
		}
		// Get an new key object with a new key
		jwtAuth, err = createJwtObject(key)
		if err != nil {
			return "", err
		}
	}
	// If there are no plugins, add them
	if plugins == nil {
		plugins = new(models.Plugins)
	}
	// Re-create the modified consumer
	jwtAuth.Exp = expTime.Unix()
	plugins.JwtAuth = jwtAuth
	modConsumer.Username = origConsumer.Value.Username
	modConsumer.Desc = origConsumer.Value.Desc
	modConsumer.Plugins = plugins
	_, err = Put(modConsumer, client)
	if err != nil {
		return "", err
	}
	// Generate JWT and return
	token, err := generateJwt(key, expTime)
	if err != nil {
		return "", err
	}
	return token, nil
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
