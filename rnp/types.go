package rnp

// Resource resource extracted from a package
type Resource struct {
	Name        string
	Description string
	Rest        []*ResourceRestConfig
	Properties  []*ResourceProperty
}

// ResourceProperty resource's property
type ResourceProperty struct {
	Name   string
	Type   string
	Fields map[string]interface{}
}

// ResourceRestConfig settings for each REST path
type ResourceRestConfig struct {
	Path    string        `json:"path"`
	Auth    string        `json:"auth"`
	Include IncludeConfig `json:"include"`
	Ignore  IgnoreConfig  `json:"ignore"`
}

// IncludeConfig default false flags
type IncludeConfig []string

// IgnoreConfig default true flags
type IgnoreConfig []string

// Parser parser of AST to extract resources
type Parser struct {
	Path      string
	Resources []*Resource
}
