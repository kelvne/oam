package utilx

// GetNSet getter and setter for map property
type GetNSet struct {
	m map[string]interface{}
	k string
}

// GetNSetFrom returns a new GetNSet for given map and key
func GetNSetFrom(m map[string]interface{}, key string) *GetNSet {
	return &GetNSet{
		m: m,
		k: key,
	}
}

// Get returns the value of a map property
func (gs *GetNSet) Get() interface{} {
	return gs.m[gs.k]
}

// Set sets the value of a map property
func (gs *GetNSet) Set(v interface{}) {
	if v != nil {
		gs.m[gs.k] = v
	}
}
