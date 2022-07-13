package astutils

// NewStruct returns a new *Struct
func NewStruct(name string) *Struct {
	return &Struct{
		Name:        name,
		Fields:      make([]*StructField, 0),
		Annotations: make(Annotations),
	}
}

// NewStructField returns a new *StructField
func NewStructField(name, fieldType string) *StructField {
	return &StructField{
		Name:        name,
		Type:        fieldType,
		Annotations: make(Annotations),
	}
}
