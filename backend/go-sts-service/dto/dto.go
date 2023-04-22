package dto

type ClientCredentialsRequest struct {
	ClientName string `json:"client_name,omitempty"`
}

type ClientCredentials struct {
	ClientName   string `json:"client_name,omitempty"`
	ClientId     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
}

type TokenRequest struct {
	ClientId     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
}

type TokenResponse struct {
	Token string `json:"token,omitempty"`
}

type ValidateTokenRequest struct {
	Token string `json:"token,omitempty"`
}

type ValidateTokenResponse struct {
	Token      string `json:"token,omitempty"`
	ClientId   string `json:"client_id,omitempty"`
	ClientName string `json:"client_name,omitempty"`
}
