package mapper

import (
	"fmt"

	"github.com/kelvne/oam/v2/openapi"
)

func dataSchema() openapi.Schema {
	return openapi.Schema{
		"type": "object",
		"properties": map[string]interface{}{
			"data": make(map[string]interface{}),
		},
	}
}

func refSchema(schemas ...string) openapi.Schema {
	schema := dataSchema()
	properties := schema["properties"].(map[string]interface{})
	properties["data"] = allOfOr(schemas...)
	return schema
}

func allOfOr(schemas ...string) map[string]interface{} {
	if len(schemas) == 1 {
		return map[string]interface{}{
			"$ref": fmt.Sprintf("#/components/schemas/%s", schemas[0]),
		}
	} else if len(schemas) > 1 {
		sl := make([]map[string]interface{}, 0)
		for _, schema := range schemas {
			sl = append(sl, allOfOr(schema))
		}
		return map[string]interface{}{
			"allOf": sl,
		}
	}
	return nil
}

func listSchema(schemas ...string) openapi.Schema {
	return map[string]interface{}{
		"allOf": []map[string]interface{}{
			{
				"$ref": "#/components/schemas/PaginatedResponseAttributes",
			},
			{
				"type": "object",
				"properties": map[string]interface{}{
					"data": map[string]interface{}{
						"items": allOfOr(schemas...),
					},
				},
			},
		},
	}
}
