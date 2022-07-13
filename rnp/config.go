package rnp

import "golang.org/x/exp/slices"

func (i IgnoreConfig) Create() bool {
	return slices.Contains(i, "create")
}

func (i IgnoreConfig) Update() bool {
	return slices.Contains(i, "Update")
}

func (i IgnoreConfig) List() bool {
	return slices.Contains(i, "List")
}

func (i IgnoreConfig) Read() bool {
	return slices.Contains(i, "Read")
}

func (i IgnoreConfig) Delete() bool {
	return slices.Contains(i, "Delete")
}

func (i IncludeConfig) ModelFields() bool {
	return slices.Contains(i, "ModelFields")
}
