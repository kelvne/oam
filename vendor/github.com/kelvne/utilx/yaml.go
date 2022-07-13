package utilx

import (
	"os"

	"gopkg.in/yaml.v2"
)

// UnmarshalYamlFile reads a yaml file and tries to unmarshal the content to destination
func UnmarshalYamlFile(path string, dest interface{}) error {
	raw, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(raw, dest)
}
