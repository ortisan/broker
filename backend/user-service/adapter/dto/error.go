package dto

type Error struct {
	Code       string `json:"code,omitempty"`
	Message    string `json:"message,omitempty"`
	Cause      string `json:"cause,omitempty"`
	StackTrace string `json:"stacktrace,omitempty"`
}

type Errors struct {
	errors *[]Error
}
