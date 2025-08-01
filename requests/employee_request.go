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

func ParseAndValidateEmployee(r *http.Request) (*CreateEmployeeRequest, error) {
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