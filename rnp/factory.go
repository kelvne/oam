package rnp

// NewParser returns a new *Parser
func NewParser(path string) *Parser {
	return &Parser{
		Path:      path,
		Resources: make([]*Resource, 0),
	}
}

// NewResource returns a new *Resource
func NewResource(name, description string) *Resource {
	return &Resource{
		Name:        name,
		Description: description,
		Properties:  make([]*ResourceProperty, 0),
		Rest:        make([]*ResourceRestConfig, 0),
	}
}

// NewResourceProperty returns a new *ResourceProperty
func NewResourceProperty(name, propType string) *ResourceProperty {
	return &ResourceProperty{
		Name:   name,
		Type:   propType,
		Fields: make(map[string]interface{}),
	}
}
