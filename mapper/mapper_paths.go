package mapper

import (
	"fmt"
	"regexp"

	"github.com/kelvne/oam/v2/openapi"
	"github.com/kelvne/oam/v2/rnp"
)

func (m *mapper) readPath(r *rnp.ResourceRestConfig) *openapi.PathOperation {
	op := "read"
	operation := m.commonOperation(op, r)
	operation.AddPathParameter(m.res.IDParam())
	operation.SetOkResponse(refSchema("Model", m.res.Name))
	return operation
}

func (m *mapper) deletePath(r *rnp.ResourceRestConfig) *openapi.PathOperation {
	op := "delete"
	operation := m.commonOperation(op, r)
	operation.AddPathParameter(m.res.IDParam())
	schema := dataSchema()
	schema["properties"].(map[string]interface{})["data"] = map[string]interface{}{
		"type": "boolean",
	}
	operation.SetOkResponse(schema)
	return operation
}

func (m *mapper) updatePath(r *rnp.ResourceRestConfig) *openapi.PathOperation {
	op := "update"
	operation := m.commonOperation(op, r)
	operation.AddPathParameter(m.res.IDParam())
	operation.SetOkResponse(refSchema("Model", m.res.Name))
	operation.RequestBody = m.modelRequestBody(r)
	return operation
}

func (m *mapper) listPath(r *rnp.ResourceRestConfig) *openapi.PathOperation {
	op := "list"
	operation := m.commonOperation(op, r)
	operation.SetOkResponse(listSchema("Model", m.res.Name))
	operation.AddQueryParameter("limit", "integer", 20)
	operation.AddQueryParameter("page", "integer", 1)
	return operation
}

func (m *mapper) createPath(r *rnp.ResourceRestConfig) *openapi.PathOperation {
	op := "create"
	operation := m.commonOperation(op, r)
	operation.SetCreatedResponse(refSchema("Model", m.res.Name))
	operation.RequestBody = m.modelRequestBody(r)
	return operation
}

func (m *mapper) modelRequestBody(r *rnp.ResourceRestConfig) *openapi.RequestBody {
	body := &openapi.RequestBody{
		Required: true,
		Content: map[string]*openapi.MediaType{
			"application/json": {},
		},
	}
	if r.Include.ModelFields() {
		body.Content["application/json"].Schema = map[string]interface{}{
			"allOf": []map[string]interface{}{
				{
					"$ref": fmt.Sprintf("#/components/schemas/%s", m.res.Name),
				},
				{
					"$ref": fmt.Sprintf("#/components/schemas/Model"),
				},
			},
		}
	} else {
		body.Content["application/json"].Schema = map[string]interface{}{
			"$ref": fmt.Sprintf("#/components/schemas/%s", m.res.Name),
		}
	}
	return body
}

func (m *mapper) ensureRest(r *rnp.ResourceRestConfig) {
	paths := m.def.PathsFor(r.Path)
	pathsWithID := m.def.PathsFor(r.IDPath(m.res))

	if !r.Ignore.Create() {
		paths.Post = m.createPath(r)
	}

	if !r.Ignore.Delete() {
		pathsWithID.Delete = m.deletePath(r)
	}

	if !r.Ignore.List() {
		paths.Get = m.listPath(r)
	}

	if !r.Ignore.Read() {
		pathsWithID.Get = m.readPath(r)
	}

	if !r.Ignore.Update() {
		pathsWithID.Put = m.updatePath(r)
	}
}

func (m *mapper) commonOperation(op string, r *rnp.ResourceRestConfig) *openapi.PathOperation {
	operation := openapi.NewOperation()
	operation.AddDescriptionForResource(op, m.res.Name)
	operation.AddStdSecurity(r.Auth)
	operation.AddTag(m.res.Name)
	matcher := regexp.MustCompile("{([a-zA-Z_]+)}")
	matched := matcher.FindAllStringSubmatch(r.Path, -1)
	params := make([]string, 0)
	for _, match := range matched {
		if len(match) > 1 {
			params = append(params, match[1])
		}
	}
	operation.AddPathParameters(params, r.Path)
	return operation
}
