package dummy

import (
	"errors"

	"github.com/Pauloo27/vcserver/internal/models"
)

type DummySource struct {
	SearchFn func(query string) ([]*models.AbstractSearchResultItem, error)
}

var _ models.AbstractSource = DummySource{}

func (d DummySource) Search(query string) ([]*models.AbstractSearchResultItem, error) {
	if d.SearchFn == nil {
		return nil, errors.New("not implemented")
	}
	return d.SearchFn(query)
}
