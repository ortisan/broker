package schema_validator

import "github.com/go-playground/validator/v10"

type SchemaValidator interface {
	ValidateSchema(any) error
}

type schemaValidator struct {
	validator *validator.Validate
}

func NewSchemaValidator() SchemaValidator {
	validator := validator.New()
	return &schemaValidator{
		validator: validator,
	}
}

func (sv schemaValidator) ValidateSchema(v any) error {
	return sv.validator.Struct(v)
}
