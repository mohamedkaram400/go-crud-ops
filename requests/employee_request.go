package requests

import (
	"encoding/json"
	"errors"
	"net/http"
	"fmt"
	"strings"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type CreateEmployeeRequest struct {
	Name       string `json:"name" validate:"required,min=3"`
	UserName   string `json:"username" validate:"required,min=4"`
	Password   string `json:"password" validate:"required,min=6"`
	Department string `json:"department" validate:"required,min=3"`
}

type UpdateEmployeeRequest struct {
	Name       string `json:"name" validate:"omitempty,min=3"`
	UserName   string `json:"username" validate:"omitempty,min=4"`
	Password   string `json:"password" validate:"omitempty,min=6"`
	Department string `json:"department" validate:"omitempty,min=3"`
}

type DeleteEmployeeRequest struct {
	EmployeeID string `json:"employeeId" validate:"required,uuid4"`
}


func ParseAndValidateCreateEmployee(r *http.Request) (*CreateEmployeeRequest, error) {
	var req CreateEmployeeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.New("invalid request body")
	}
	if err := validate.Struct(req); err != nil {
		return nil, formatValidationError(err)
	}
	return &req, nil
}

func ParseAndValidateUpdateEmployee(r *http.Request) (*UpdateEmployeeRequest, error) {
	var req UpdateEmployeeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.New("invalid request body")
	}

	if req.Name == "" && req.UserName == "" && req.Password == "" && req.Department == "" {
		return nil, errors.New("at least one field must be provided for update")
	}

	if err := validate.Struct(req); err != nil {
		return nil, formatValidationError(err)
	}
	return &req, nil
}

func ParseAndValidateDeleteEmployee(r *http.Request) (*DeleteEmployeeRequest, error) {
	var req DeleteEmployeeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.New("invalid request body")
	}
	if err := validate.Struct(req); err != nil {
		return nil, formatValidationError(err)
	}
	return &req, nil
}

func formatValidationError(err error) error {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		messages := make([]string, 0, len(validationErrors))
		for _, err := range validationErrors {
			field := err.Field()
			tag := err.Tag()

			switch field {
			case "Name":
				if tag == "min" {
					messages = append(messages, "Name must be at least 3 characters")
				}
			case "UserName":
				if tag == "min" {
					messages = append(messages, "Username must be at least 4 characters")
				}
			case "Password":
				if tag == "min" {
					messages = append(messages, "Password must be at least 6 characters")
				}
			case "Department":
				if tag == "min" {
					messages = append(messages, "Department must be at least 3 characters")
				}
			case "EmployeeID":
				if tag == "required" {
					messages = append(messages, "Employee ID is required")
				}
				if tag == "uuid4" {
					messages = append(messages, "Employee ID must be a valid UUID")
				}
			default:
				messages = append(messages, fmt.Sprintf("%s is invalid", field))
			}
		}
		return errors.New(strings.Join(messages, "; "))
	}
	return err
}