package requests

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type CreateEmployeeRequest struct {
	Name       string `json:"name" validate:"required,min=3"`
	Department string `json:"department" validate:"required"`
}

func ParseAndValidateCreateEmployee(r *http.Request) (*CreateEmployeeRequest, error) {
	var req CreateEmployeeRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.New("invalid request body")
	}

	err = validate.Struct(req)
	if err != nil {
		return nil, err
	}

	return &req, nil
}

type UpdateEmployeeRequest struct {
	Name       string `json:"name" validate:"omitempty,min=3"`
	Department string `json:"department" validate:"omitempty,min=3"`
}

func ParseAndValidateUpdateEmployee(r *http.Request) (*UpdateEmployeeRequest, error) {
	var req UpdateEmployeeRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.New("invalid request body")
	}

	// check at least one field is present
	if req.Name == "" && req.Department == "" {
		return nil, errors.New("at least one of name or department must be provided")
	}

	err = validate.Struct(req)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			// Convert to a custom message
			return nil, errors.New(buildValidationErrorMessage(validationErrors))
		}
		return nil, err
	}

	return &req, nil
}

func buildValidationErrorMessage(errs validator.ValidationErrors) string {
	for _, err := range errs {
		switch err.Field() {
		case "Name":
			if err.Tag() == "min" {
				return "Name must be at least 3 characters long"
			}
		case "Department":
			if err.Tag() == "min" {
				return "Department must be at least 3 characters long"
			}
		}
	}
	return "Invalid input data"
}