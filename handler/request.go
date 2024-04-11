package handler

import "fmt"

func errParamIsRequired(param string, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", param, typ)
}

// Create Opening

type CreateOpeningRequest struct {
	Role     string `json:"role" binding:"required"`
	Company  string `json:"company" binding:"required"`
	Location string `json:"location" binding:"required"`
	Remote   *bool  `json:"remote" binding:"required"`
	Link     string `json:"link" binding:"required"`
	Salary   int64  `json:"salary" binding:"required"`
}

func (r *CreateOpeningRequest) Validate() error {
	fields := map[string]interface{}{
		"role":     r.Role,
		"company":  r.Company,
		"location": r.Location,
		"link":     r.Link,
		"remote":   r.Remote,
		"salary":   r.Salary,
	}

	types := map[string]string{
		"role":     "string",
		"company":  "string",
		"location": "string",
		"link":     "string",
		"remote":   "bool",
		"salary":   "int",
	}

	for field, value := range fields {
		switch v := value.(type) {
		case string:
			if v == "" {
				return errParamIsRequired(field, types[field])
			}
		case *bool:
			if v == nil {
				return errParamIsRequired(field, types[field])
			}
		case int64:
			if v <= 0 {
				return errParamIsRequired(field, types[field])
			}
		}
	}

	return nil
}

type UpdateOpeningRequest struct {
	Role     string `json:"role" binding:"required"`
	Company  string `json:"company" binding:"required"`
	Location string `json:"location" binding:"required"`
	Remote   *bool  `json:"remote" binding:"required"`
	Link     string `json:"link" binding:"required"`
	Salary   int64  `json:"salary" binding:"required"`
}

func (r *UpdateOpeningRequest) Validate() error {
	// If any field is provided, validation is true
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Link != "" || r.Remote != nil || r.Salary > 0 {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
