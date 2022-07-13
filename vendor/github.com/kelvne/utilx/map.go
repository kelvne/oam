package utilx

import "strconv"

// Map is a more complete map
type Map map[string]interface{}

// String returns property value as string
func (m Map) String(key string, or ...string) string {
	if prop, ok := m[key]; ok {
		if value, ok := prop.(string); ok {
			return GetStringOrDefault("", value)
		}
	}
	return GetStringOrDefault("", or...)
}

// Int returns property value as int
func (m Map) Int(key string, or ...int) int {
	if prop, ok := m[key]; ok {
		switch value := prop.(type) {
		case string:
			if v, err := strconv.Atoi(value); err == nil {
				return GetNumberOrDefault(0, v)
			}
		case int:
			return value
		case uint:
		case uint32:
		case uint64:
		case int64:
		case int32:
		case float64:
		case float32:
			return int(value)
		}
	}
	return GetNumberOrDefault(0, or...)
}

// Bool returns property value as bool
func (m Map) Bool(key string, or ...bool) bool {
	if prop, ok := m[key]; ok {
		if value, ok := prop.(bool); ok {
			return GetBoolOrDefault(false, value)
		}
	}
	return GetBoolOrDefault(false, or...)
}

// MapStr utilities for map[string]string
type MapStr map[string]string

// Str returns the value from key as a string
func (m MapStr) Str(k string) string {
	if v, ok := m[k]; ok {
		return v
	}
	return ""
}

// Int returns the value from key as a int
func (m MapStr) Int(k string) int {
	if v, ok := m[k]; ok {
		if i, err := strconv.Atoi(v); err != nil {
			return i
		}
	}
	return 0
}

// Float returns the value from key as a float64
func (m MapStr) Float(k string) float64 {
	if v, ok := m[k]; ok {
		if f, err := strconv.ParseFloat(v, 64); err != nil {
			return f
		}
	}
	return 0
}

// Bool returns the value from key as a bool
func (m MapStr) Bool(k string) bool {
	if v, ok := m[k]; ok {
		if b, err := strconv.ParseBool(v); err != nil {
			return b
		}
	}
	return false
}
