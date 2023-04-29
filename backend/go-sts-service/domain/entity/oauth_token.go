package entity

import (
	"fmt"
	errApp "ortisan-broker/go-commons/error"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type OauthToken interface {
	ClientCredentials() ClientCredentials
	Value() string
	Claims() map[string]any
	ExpirationTime() time.Time
}

type oauthToken struct {
	clientCredentials ClientCredentials
	value             string
	claims            map[string]any
	expirationTime    time.Time
}

func (o *oauthToken) ClientCredentials() ClientCredentials {
	return o.clientCredentials
}

func (o *oauthToken) Value() string {
	return o.value
}

func (o *oauthToken) Claims() map[string]any {
	return o.claims
}

func (o *oauthToken) ExpirationTime() time.Time {
	return o.expirationTime
}

func NewOauthTokenFromToken(credentials ClientCredentials, token string) (OauthToken, error) {
	if credentials == nil {
		return nil, errApp.NewBadArgumentError("client credentials is required")
	}
	if token == "" {
		return nil, errApp.NewBadArgumentError("token is required")
	}

	claims := jwt.MapClaims{}
	jwtToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errApp.NewBadArgumentError(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
		}
		if _, ok := token.Header["client_id"]; !ok {
			return nil, errApp.NewBadArgumentError("Invalid token. client_id was not found.")
		}

		if token.Header["client_id"] != credentials.ClientId() {
			return nil, errApp.NewBadArgumentError("Invalid token. Divergence of client_id.")
		}

		return []byte(credentials.ClientSecret().Value()), nil
	})

	if err != nil || !jwtToken.Valid {
		errApp.NewBadArgumentError("token is invalid")
	}

	return &oauthToken{clientCredentials: credentials, value: token}, nil
}

func NewOauthToken(credentials ClientCredentials) (OauthToken, error) {
	if credentials == nil {
		return nil, errApp.NewBadArgumentError("client credentials is required")
	}

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["client_id"] = credentials.ClientId().Value()
	atClaims["client_name"] = credentials.ClientName().Value()
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	at.Header["client_id"] = credentials.ClientId().Value()
	token, err := at.SignedString([]byte(credentials.ClientSecret().Value()))
	if err != nil {
		return nil, errApp.NewBaseErrorWithCause("Error to generate token.", err)
	}

	return &oauthToken{clientCredentials: credentials, value: token}, nil
}
