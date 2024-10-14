package jwt

import (
	"crypto/hmac"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/MokhtarSMokhtar/online-wallet/comman/config"
	"strings"
	"time"
)

func ValidateToken(tokenString string) (*Claims, error) {
	configu := config.NewConfig()
	secretKey := []byte(configu.JWTSecret)
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid token")
	}
	headerEncoded := parts[0]
	claimsEncoded := parts[1]
	signatureEncoded := parts[2]

	unsignedToken := headerEncoded + "." + claimsEncoded
	// Verify the signature

	signature, err := base64.RawURLEncoding.DecodeString(signatureEncoded)
	if err != nil {
		return nil, errors.New("invalid signature encoding")
	}

	expectedSignature, err := sign(unsignedToken, secretKey)
	if err != nil {
		return nil, err
	}

	if !hmac.Equal(signature, expectedSignature) {
		return nil, errors.New("invalid token signature")
	}
	// Decode the claims
	claimsJSON, err := base64.RawURLEncoding.DecodeString(claimsEncoded)
	if err != nil {
		return nil, errors.New("invalid claims encoding")
	}
	var claims Claims
	err = json.Unmarshal(claimsJSON, &claims)
	if err != nil {
		return nil, errors.New("invalid claims")
	}

	if claims.Issuer != configu.ISSUER {
		return nil, errors.New("invalid token")
	}
	// Check expiration
	if time.Now().Unix() > claims.Exp {
		return nil, errors.New("token has expired")
	}

	return &claims, nil
}
