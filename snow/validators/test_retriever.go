package validators

import (
	"github.com/flare-foundation/flare/ids"
)

type RetrieverMock struct {
	GetValidatorsByBlockIDFunc func(blockID ids.ID) (Set, error)
}

func (m *RetrieverMock) GetValidatorsByBlockID(blockID ids.ID) (Set, error) {
	return m.GetValidatorsByBlockIDFunc(blockID)
}
