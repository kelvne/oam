package rnp

import (
	"fmt"
	"strings"
)

// LowerName returns the lowercased resource name
func (r *Resource) LowerName() string {
	return strings.ToLower(r.Name)
}

// IDParam returns the id parameter for resource
func (r *Resource) IDParam() string {
	return fmt.Sprintf("%s_id", r.LowerName())
}

// IDPath returns the path with ID
func (r *ResourceRestConfig) IDPath(res *Resource) string {
	return fmt.Sprintf("%s/{%s}", strings.ToLower(r.Path), res.IDParam())
}
