package mapper

import (
	"encoding/json"
	"io/ioutil"

	"github.com/kelvne/oam/v2/openapi"
	"github.com/kelvne/oam/v2/rnp"
)

// MapResourcesToDefinitionFile maps resources to OpenAPI file
func MapResourcesToDefinitionFile(resources []*rnp.Resource, inPath, outPath string) error {
	definition, err := readDefinitionFile(inPath)
	if err != nil {
		return err
	}
	for _, r := range resources {
		resourceToDefinition(definition, r)
	}
	ensurePaginatedAttributes(definition)
	ensureModelDefaultAttributes(definition)
	return definition.WriteToFile(outPath)
}

func newMapper(def *openapi.Definition, res *rnp.Resource) *mapper {
	return &mapper{
		def: def,
		res: res,
	}
}

func readDefinitionFile(path string) (*openapi.Definition, error) {
	baseRaw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	base := new(openapi.Definition)
	if err := json.Unmarshal(baseRaw, base); err != nil {
		return nil, err
	}
	return base, err
}

func resourceToDefinition(d *openapi.Definition, r *rnp.Resource) {
	schema := d.Schemas().Schema(r.Name)
	schema.Type().Set("object")
	schema.Description().Set(r.Description)

	for _, p := range r.Properties {
		prop := schema.Properties().Property(p.Name)
		prop.Default().Set(p.Fields["default"])
		prop.Description().Set(p.Fields["description"])
		prop.Format().Set(p.Fields["format"])
		prop.Maximum().Set(p.Fields["maximum"])
		prop.Minimum().Set(p.Fields["minimum"])
		prop.Nullable().Set(p.Fields["nullable"])
		prop.Type().Set(p.Type)
	}

	m := newMapper(d, r)

	for _, rest := range r.Rest {
		m.ensureRest(rest)
	}
}
