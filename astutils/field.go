package astutils

import "strings"

// SchemaType returns the JSON type and format
func (sf *StructField) SchemaType() (string, string) {
	if strings.Contains(sf.Type, "float") {
		return "number", ""
	} else if strings.Contains(sf.Type, "int") {
		return "integer", ""
	} else if sf.Type == "string" {
		return sf.Type, ""
	} else if sf.Type == "time.Time" {
		return "string", "date"
	} else if sf.Type == "bool" {
		return "boolean", ""
	}
	return "object", ""
}
