package mapper

import "github.com/kelvne/oam/v2/openapi"

func ensureModelDefaultAttributes(def *openapi.Definition) {
	schema := openapi.Schemas(def.Components.Schemas).Schema("Model")
	schema.Description().Set("Default attributes for every model")
	schema.Type().Set("object")

	schema.Properties().Property("ID").Type().Set("integer")
	schema.Properties().Property("ID").Description().Set("Primary key")

	schema.Properties().Property("CreatedAt").Type().Set("string")
	schema.Properties().Property("CreatedAt").Description().Set("Timestamp representing when the entity was created")
	schema.Properties().Property("CreatedAt").Format().Set("date")

	schema.Properties().Property("UpdatedAt").Type().Set("string")
	schema.Properties().Property("UpdatedAt").Description().Set("Timestamp representing when the entity was updated")
	schema.Properties().Property("UpdatedAt").Format().Set("date")

	schema.Properties().Property("DeletedAt").Type().Set("string")
	schema.Properties().Property("DeletedAt").Description().Set("Timestamp representing when the entity was deleted")
	schema.Properties().Property("DeletedAt").Format().Set("date")
}

func ensurePaginatedAttributes(def *openapi.Definition) {
	schema := openapi.Schemas(def.Components.Schemas).Schema("PaginatedResponseAttributes")
	schema.Description().Set("Common attributes for a paginated response")
	schema.Type().Set("object")

	schema.Properties().Property("page").Type().Set("integer")
	schema.Properties().Property("page").Description().Set("Page of the listed response objects.")
	schema.Properties().Property("page").Default().Set(1)

	schema.Properties().Property("limit").Type().Set("integer")
	schema.Properties().Property("limit").Description().Set("Maximum amount of returned objects.")
	schema.Properties().Property("limit").Default().Set(20)

	schema.Properties().Property("total").Type().Set("integer")
	schema.Properties().Property("total").Description().Set("Total number of objects found on the database.")
}
