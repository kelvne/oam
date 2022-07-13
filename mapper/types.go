package mapper

import (
	"github.com/kelvne/oam/v2/openapi"
	"github.com/kelvne/oam/v2/rnp"
)

type mapper struct {
	def *openapi.Definition
	res *rnp.Resource
}
