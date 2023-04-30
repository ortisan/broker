package dto

type ClientCredentialsRequest struct {
	ClientName string `json:"client_name,omitempty" validate:"required"`
}

type ClientCredentials struct {
	ClientName   string `json:"client_name,omitempty" validate:"required"`
	ClientId     string `json:"client_id,omitempty" validate:"required"`
	ClientSecret string `json:"client_secret,omitempty" validate:"required"`
}

type OauthTokenRequest struct {
	ClientId     string `json:"client_id,omitempty" validate:"required"`
	ClientSecret string `json:"client_secret,omitempty" validate:"required"`
}

type OauthTokenResponse struct {
	ClientId     string `json:"client_id,omitempty" validate:"required"`
	ClientSecret string `json:"client_secret,omitempty" validate:"required"`
	Token        string `json:"token,omitempty"`
}

type ValidateTokenRequest struct {
	Token string `json:"token,omitempty" validate:"required"`
}

type ValidateTokenResponse struct {
	Token      string `json:"token,omitempty" validate:"required"`
	ClientId   string `json:"client_id,omitempty" validate:"required"`
	ClientName string `json:"client_name,omitempty" validate:"required"`
}
