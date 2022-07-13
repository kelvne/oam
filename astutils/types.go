package astutils

// Annotations annotations for an object
type Annotations map[string]interface{}

// Struct extracted struct from an AST
type Struct struct {
	Name        string
	Fields      []*StructField
	Annotations Annotations
}

// StructField mapped fields from a struct
type StructField struct {
	Name        string
	Type        string
	Pointer     bool
	Annotations Annotations
}
