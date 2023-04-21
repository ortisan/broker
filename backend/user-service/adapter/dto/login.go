package dto

type LoginData struct {
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"min=6,max=20"`
}

type RecoverLoginData struct {
	Email string `json:"email,omitempty" validate:"required,email"`
}

type RecoverLoginDataResponse struct {
	ID    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
}
