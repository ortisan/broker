package dto

import "github.com/go-playground/validator/v10"

type LoginData struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (l LoginData) ValidateSchema() error {
	var validate = validator.New()
	rules := map[string]string{
		"Email":    "email",
		"Password": "min=6,max=20",
	}
	validate.RegisterStructValidationMapRules(rules, LoginData{})
	err := validate.Struct(l)
	return err
}

type RecoverLoginData struct {
	Email string `json:"email,omitempty"`
}

func (r RecoverLoginData) ValidateSchema() error {
	var validate = validator.New()
	rules := map[string]string{
		"Email": "min=4,max=6",
	}
	validate.RegisterStructValidationMapRules(rules, RecoverLoginData{})
	err := validate.Struct(r)
	return err
}

type RecoverLoginDataResponse struct {
	ID    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
}
