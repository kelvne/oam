package astutils

import "strings"

// Single returns an annotation as single string
func (a Annotations) Single(key string) string {
	if asSlice, ok := a[key].([]string); ok {
		if len(asSlice) > 0 {
			return asSlice[0]
		}
	}
	return ""
}

// Multiple returns an annotation as multiple values
func (a Annotations) Multiple(key string) []interface{} {
	if asSlice, ok := a[key].([]interface{}); ok {
		if len(asSlice) > 0 {
			return asSlice
		}
	}
	return make([]interface{}, 0)
}

// Map returns an annotation as a map (from a JSONString)
func (a Annotations) Map(key string) map[string]interface{} {
	if asMap, ok := a[key].(map[string]interface{}); ok {
		return asMap
	}
	return make(map[string]interface{})
}

func (a Annotations) addLine(l string) {
	key, value := a.parseLine(strings.TrimSpace(strings.TrimPrefix(l, "//")))
	if key != "" {
		if _, ok := a[key]; !ok {
			a[key] = value
		} else {
			if _, ok := a[key].([]interface{}); ok {
				a[key] = append(a.Multiple(key), value)
			} else {
				a[key] = append([]interface{}{
					a[key],
				}, value)
			}
		}
	}
}

func (a Annotations) parseLine(l string) (key string, value interface{}) {
	if strings.Contains(l, "@") {
		kv := strings.SplitN(l, ":", 2)
		if len(kv) > 0 {
			key = strings.TrimPrefix(kv[0], "@")
		}
		if len(kv) > 1 {
			if m, ok := mapIfJSON(kv[1]); ok {
				value = m
			} else {
				value = kv[1]
			}
		}
	}
	return
}
