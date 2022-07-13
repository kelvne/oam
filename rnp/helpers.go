package rnp

import (
	"encoding/json"
	"strconv"

	"github.com/kelvne/oam/v2/astutils"
)

func unmarshalRestConfig(m map[string]interface{}) *ResourceRestConfig {
	rrc := new(ResourceRestConfig)
	if b, err := json.Marshal(m); err == nil {
		json.Unmarshal(b, rrc)
	}
	return rrc
}

func resourceFromStruct(s *astutils.Struct) *Resource {
	r := NewResource(s.Name, s.Annotations.Single("description"))
	etlRest(r, s)
	etlFields(r, s)
	return r
}

func etlRest(r *Resource, s *astutils.Struct) {
	if m := s.Annotations.Map("rest"); len(m) > 0 {
		etlRestMap(m, r)
	} else if sl := s.Annotations.Multiple("rest"); len(sl) > 0 {
		for _, item := range sl {
			if m, ok := item.(map[string]interface{}); ok {
				etlRestMap(m, r)
			}
		}
	}
}

func etlRestMap(m map[string]interface{}, r *Resource) {
	restConfig := unmarshalRestConfig(m)
	r.Rest = append(r.Rest, restConfig)
}

func etlFields(r *Resource, s *astutils.Struct) {
	for _, field := range s.Fields {
		fieldType, fieldFormat := field.SchemaType()
		prop := NewResourceProperty(field.Name, fieldType)

		if fieldFormat != "" {
			prop.Fields["format"] = fieldFormat
		}

		if field.Pointer {
			prop.Fields["nullable"] = true
		}

		description := field.Annotations.Single("description")
		maximum, _ := strconv.Atoi(field.Annotations.Single("maximum"))
		minimum, _ := strconv.Atoi(field.Annotations.Single("minimum"))

		if description != "" {
			prop.Fields["description"] = description
		}

		if maximum != 0 {
			prop.Fields["maximum"] = maximum
		}

		if minimum != 0 {
			prop.Fields["minimum"] = minimum
		}

		def := field.Annotations.Single("default")
		if def != "" {
			switch fieldType {
			case "string":
				prop.Fields["default"] = def
			case "integer":
				prop.Fields["default"], _ = strconv.Atoi(def)
			case "float":
				prop.Fields["default"], _ = strconv.ParseFloat(def, 64)
			default:
				prop.Fields["default"], _ = strconv.ParseBool(def)
			}
		}

		r.Properties = append(r.Properties, prop)
	}
}
