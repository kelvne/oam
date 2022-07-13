package utilx

import (
	"strings"

	"golang.org/x/exp/constraints"
)

// GetOrDefault returns valid provided value or default value
func GetOrDefault[K any](defaultValue K, value ...K) K {
	if len(value) > 0 {
		return value[0]
	}
	return defaultValue
}

// GetStringOrDefault returns valid provided or default string
func GetStringOrDefault(defaultString string, value ...string) string {
	v := GetOrDefault(defaultString, value...)
	if strings.TrimSpace(v) != "" {
		return v
	}
	return defaultString
}

// GetNumberOrDefault returns valid provided or default number
func GetNumberOrDefault[K constraints.Ordered](defaultNumber K, value ...K) K {
	return GetOrDefault(defaultNumber, value...)
}

// GetBoolOrDefault returns valid provided or default bool
func GetBoolOrDefault(defaultBool bool, value ...bool) bool {
	return GetOrDefault(defaultBool, value...)
}
