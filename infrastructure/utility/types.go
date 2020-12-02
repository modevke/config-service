package utility

import (

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type EnvironmentVariables struct {
	Port        string
	Host        string
	Scheme      string
	ProjectName string
	Environment string
}

func (e EnvironmentVariables) Validate() error {

	return validation.ValidateStruct(&e,
		validation.Field(&e.Port, validation.Required),
		validation.Field(&e.Host, validation.Required),
		validation.Field(&e.Scheme, validation.Required),
		validation.Field(&e.ProjectName, validation.Required),
		validation.Field(&e.Environment, validation.Required),
	)
}


type IDVariable struct{
	ID   string	`json:"id"`
}