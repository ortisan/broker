package dto

type User struct {
	ID               string `json:"id,omitempty"`
	FederationId     string `json:"federation_id,omitempty"`
	Name             string `json:"name,omitempty" validate:"required,min=1,max=100"`
	Email            string `json:"email,omitempty" validate:"required,email"`
	Password         string `json:"password,omitempty"`
	ProfileAvatarUrl string `json:"profile_avatar_url,omitempty"`
}

type UserUpdate struct {
	ID               string `json:"id,omitempty" validate:"required"`
	Name             string `json:"name,omitempty" validate:"required,min=1,max=100"`
	Email            string `json:"email,omitempty" validate:"required,email"`
	ProfileAvatarUrl string `json:"profile_avatar_url,omitempty"`
}

type UserResponse struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Token string `json:"auth_token,omitempty"`
}
