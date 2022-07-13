package testassets

// JSONAnnotatedModel
//
// @rest:{ "path": "/mngr/jsonannotatedmodel", "auth": "managerAuth", "include": ["ModelFields"] }
// @rest:{ "path": "/admin/base/{baseId}/json_annotated_models", "auth": "userAuth" }
// @rest:{ "path": "/admin/base/{baseId}/relations/{relationId}/json_annotated_models", "auth": "userAuth", "ignore": { "op_create": true } }
type JSONAnnotatedModel struct {
	// @description:This is a string field
	FieldOne string
	// @default:10.90
	FieldTwo float64
	// @maximum:10
	FieldThree *int
}
