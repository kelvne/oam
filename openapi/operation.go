package openapi

import (
	"fmt"
	"strings"
)

// NewOperation creates a new *PathOperation with default settings
func NewOperation() *PathOperation {
	return &PathOperation{
		Security:   make([]map[string][]string, 0),
		Parameters: make([]*PathParameter, 0),
		Tags:       make([]string, 0),
		Responses:  make(Responses),
	}
}

// SetRestDescription sets the description for a restful resource
func (p *PathOperation) AddDescriptionForResource(operation string, resource string) {
	switch strings.ToLower(operation) {
	case "read":
		p.Description = fmt.Sprintf("Retrieve a single %s.", resource)
	case "create":
		p.Description = fmt.Sprintf("Create a %s.", resource)
	case "update":
		p.Description = fmt.Sprintf("Updates a %s.", resource)
	case "list":
		p.Description = fmt.Sprintf("Retrieve multiple %s.", resource)
	case "delete":
		p.Description = fmt.Sprintf("Delete a single %s.", resource)
	default:
		p.Description = fmt.Sprintf("%s endpoint for %s.", operation, resource)
	}
}

// AddPathParameter adds a new path parameter
func (p *PathOperation) AddPathParameter(name string) {
	p.addParameter(name, "path", true, false, "integer", 0)
}

// AddPathParameters adds multiple path parameters
func (p *PathOperation) AddPathParameters(params []string, path string) {
	if len(params) > 0 {
		for _, param := range params {
			if strings.Contains(path, param) {
				p.AddPathParameter(param)
			}
		}
	}
}

// AddQueryParameter adds a new query parameter
func (p *PathOperation) AddQueryParameter(name, t string, def interface{}) {
	p.addParameter(name, "query", false, true, t, def)
}

func (p *PathOperation) addParameter(name, in string, required, allowEmptyValue bool, t string, def interface{}) {
	p.Parameters = append(p.Parameters, &PathParameter{
		Name:            name,
		In:              in,
		Required:        required,
		AllowEmptyValue: allowEmptyValue,
		Schema: Schema{
			"type":    t,
			"default": def,
		},
	})
}

// AddTag adds a new tag
func (p *PathOperation) AddTag(tag string) {
	p.Tags = append(p.Tags, tag)
}

// AddStdSecurity adds a new standard security requirement
func (p *PathOperation) AddStdSecurity(security string) {
	if security != "" {
		p.Security = []map[string][]string{
			{
				security: make([]string, 0),
			},
		}
	}
}

// SetOkResponse sets a new http.StatusOK (200) response
func (p *PathOperation) SetOkResponse(schema Schema) {
	p.setSuccessfulResponse("200", schema)
}

// SetCreatedResponse sets a new http.StatusCreated (201) response
func (p *PathOperation) SetCreatedResponse(schema Schema) {
	p.setSuccessfulResponse("201", schema)
}

func (p *PathOperation) setSuccessfulResponse(code string, schema Schema) {
	p.Responses[code] = &Response{
		Description: "Successful response",
		Content: map[string]*MediaType{
			"application/json": {
				Schema: schema,
			},
		},
	}
}
