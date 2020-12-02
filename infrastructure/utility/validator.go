package utility

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/go-ozzo/ozzo-validation/v4"
)


func ValidateBody(data io.ReadCloser, request interface{}, fields ...*validation.FieldRules) []ResponseError{
	
	var errResponses []ResponseError
	err := json.NewDecoder(data).Decode(request)

	if err != nil {

		errResponses = append(errResponses, ResponseError{
			Code: "4000",
			Description: "Invalid JSON body",
		})

		return errResponses
	}

	verr := validation.ValidateStruct(request, fields...)

	if verr != nil {
		verrStr := strings.Split(verr.Error(), ";")

		for _, val := range verrStr{
			errResponses = append(errResponses, ResponseError{
				Code: "40001",
				Description: val,
			})
		}
		return errResponses
	}

	return errResponses
}

func ValidateParams(request interface{}, fields ...*validation.FieldRules) []ResponseError {
	var errResponses []ResponseError
	verr := validation.ValidateStruct(request, fields...)

	if verr != nil {
		verrStr := strings.Split(verr.Error(), ";")

		for _, val := range verrStr{
			errResponses = append(errResponses, ResponseError{
				Code: "40002",
				Description: val,
			})
		}
		return errResponses
	}

	return errResponses
}