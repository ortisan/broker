package dto

import "github.com/go-playground/validator/v10"

type User struct {
	ID              string `json:"id,omitempty"`
	FederationId    string `json:"federation_id,omitempty"`
	Name            string `json:"name,omitempty"`
	Email           string `json:"email,omitempty"`
	Password        string `json:"password,omitempty"`
	ProfilePhotoUrl string `json:"profile_photo_url,omitempty"`
}

func (u User) ValidateSchema() error {
	var validate = validator.New()
	rules := map[string]string{
		"Name":  "min=1,max=100",
		"Email": "email",
	}
	validate.RegisterStructValidationMapRules(rules, RecoverLoginData{})
	err := validate.Struct(u)
	return err
}

func NewUser(
	ID string,
	FederationId string,
	Name string,
	Email string,
	ProfilePhotoUrl string,
) User {
	return User{
		ID:              ID,
		FederationId:    FederationId,
		Name:            Name,
		Email:           Email,
		ProfilePhotoUrl: ProfilePhotoUrl,
	}
}

type UserUpdate struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type UserResponse struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Token string `json:"auth_token,omitempty"`
}
