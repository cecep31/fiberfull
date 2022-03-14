package validation

import (
	"fiberfull/exception"
	"fiberfull/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func Validate(request model.CreateProductRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Id, validation.Required),
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Price, validation.Required, validation.Min(1)),
		validation.Field(&request.Quantity, validation.Required, validation.Min(0)),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
