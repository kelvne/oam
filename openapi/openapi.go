package openapi

import (
	"encoding/json"
	"os"

	"github.com/kelvne/utilx"
)

// =====================
// Definition
// =====================

// PathsFor returns a *Path for given key
func (d *Definition) PathsFor(key string) *Path {
	if _, ok := d.Paths[key]; !ok {
		d.Paths[key] = &Path{}
	}
	return d.Paths[key]
}

// Schemas returns schemas object on components property as Schemas type
func (d *Definition) Schemas() Schemas {
	return d.Components.Schemas
}

// WriteToFile writes the definition to a yaml file
func (d *Definition) WriteToFile(path string) error {
	marshaled, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, marshaled, os.ModePerm)
}

// =====================
// Schemas
// =====================

// Schema returns a schema as a Schema type
func (s Schemas) Schema(name string) Schema {
	if _, ok := s[name]; !ok {
		s[name] = make(Schema)
	}
	return s[name]
}

// =====================
// Schema
// =====================

// Type returns the getter and setter for the schema type
func (s Schema) Type() *utilx.GetNSet {
	return utilx.GetNSetFrom(s, "type")
}

// Description returns the getter and setter for the schema description
func (s Schema) Description() *utilx.GetNSet {
	return utilx.GetNSetFrom(s, "description")
}

// Properties returns the getter and setter for the schema properties
func (s Schema) Properties() SchemaProperties {
	if _, ok := s["properties"]; !ok {
		s["properties"] = make(map[string]interface{})
	}
	return SchemaProperties(s["properties"].(map[string]interface{}))
}

// =====================
// SchemaProperties
// =====================

// Property returns a property as a SchemaProperty type
func (sp SchemaProperties) Property(name string) SchemaProperty {
	if _, ok := sp[name]; !ok {
		sp[name] = make(map[string]interface{})
	}
	return SchemaProperty(sp[name].(map[string]interface{}))
}

// =====================
// SchemaProperty
// =====================

// Type returns the getter and setter for the property type
func (sp SchemaProperty) Type() *utilx.GetNSet {
	return utilx.GetNSetFrom(sp, "type")
}

// Description returns the getter and setter for the property description
func (sp SchemaProperty) Description() *utilx.GetNSet {
	return utilx.GetNSetFrom(sp, "description")
}

// Default returns the getter and setter for the property default
func (sp SchemaProperty) Default() *utilx.GetNSet {
	return utilx.GetNSetFrom(sp, "default")
}

// Maximum returns the getter and setter for the property maximum
func (sp SchemaProperty) Maximum() *utilx.GetNSet {
	return utilx.GetNSetFrom(sp, "maximum")
}

// Minimum returns the getter and setter for the property minimum
func (sp SchemaProperty) Minimum() *utilx.GetNSet {
	return utilx.GetNSetFrom(sp, "minimum")
}

// Nullable returns the getter and setter for the property nullable
func (sp SchemaProperty) Nullable() *utilx.GetNSet {
	return utilx.GetNSetFrom(sp, "nullable")
}

// Format returns the getter and setter for the property format
func (sp SchemaProperty) Format() *utilx.GetNSet {
	return utilx.GetNSetFrom(sp, "format")
}
