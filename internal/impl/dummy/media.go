package dummy

import "github.com/Pauloo27/vcserver/internal/models"

type DummyMedia struct {
	DummySearchResultItem
}

var _ models.AbstractMedia = DummyMedia{}
