package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"github.com/MokhtarSMokhtar/online-wallet/comman/config"
	"time"
)

// Header represents the JWT header
type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

// Claims represents the JWT payload
type Claims struct {
	UserId   string `json:"user-id"`
	Name     string `json:"name"`     // User's name
	Email    string `json:"email"`    // User's email
	Phone    string `json:"phone"`    // User's email
	Verified string `json:"verified"` // User's email
	Exp      int64  `json:"exp"`      // Expiration time
	Issuer   string `json:"iss"`
}

// GenerateToken creates a JWT token string
func GenerateToken(claims Claims) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	configu := config.NewConfig()
	// Set standard claims
	claims.Exp = time.Now().Add(time.Hour * 40).Unix()
	claims.Issuer = configu.ISSUER
	// Marshal header and claims
	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	claimsJSON, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	// Encode header and claims
	headerEncoded := base64.RawURLEncoding.EncodeToString(headerJSON)
	claimsEncoded := base64.RawURLEncoding.EncodeToString(claimsJSON)

	// Create the unsigned token
	unsignedToken := headerEncoded + "." + claimsEncoded

	// Sign the token
	signature, err := sign(unsignedToken, []byte(configu.JWTSecret))
	if err != nil {
		return "", err
	}

	signatureEncoded := base64.RawURLEncoding.EncodeToString(signature)

	// Combine all parts
	token := unsignedToken + "." + signatureEncoded

	return token, nil
}

// sign creates an HMAC SHA256 signature
func sign(data string, secretKey []byte) ([]byte, error) {
	h := hmac.New(sha256.New, secretKey)
	_, err := h.Write([]byte(data))
	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}
