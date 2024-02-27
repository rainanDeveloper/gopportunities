package handler

import "fmt"

func errParamIsRequired(name string, typ string) error {
	return fmt.Errorf("param \"%s\" (type: %s) is required", name, typ)
}

type CreateOpportunityRequest struct {
	Role     string  `json:"role"`
	Company  string  `json:"company"`
	Location string  `json:"location"`
	Remote   *bool   `json:"remote"`
	Link     string  `json:"link"`
	Salary   float64 `json:"salary"`
}

func (request *CreateOpportunityRequest) Validate() error {
	if request.Role == "" && request.Company == "" && request.Location == "" && request.Remote == nil && request.Link == "" && request.Salary <= 0 {
		return fmt.Errorf("request is empty or malformed")
	}

	if request.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if request.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if request.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if request.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}
	if request.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if request.Salary <= 0 {
		return errParamIsRequired("salary", "float64")
	}

	return nil
}

type UpdateOpportunityRequest struct {
	Role     string  `json:"role"`
	Company  string  `json:"company"`
	Location string  `json:"location"`
	Remote   *bool   `json:"remote"`
	Link     string  `json:"link"`
	Salary   float64 `json:"salary"`
}

func (request *UpdateOpportunityRequest) Validate() error {
	if request.Role != "" || request.Company != "" || request.Location != "" || request.Remote != nil || request.Link != "" || request.Salary > 0 {
		return nil
	}

	return fmt.Errorf("at least one field must be informed")
}
