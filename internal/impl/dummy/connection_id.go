package dummy

import "github.com/Pauloo27/vcserver/internal/models"

type DummyConnectionID struct {
	ID string
}

var _ models.AbstractConnectionID = DummyConnectionID{}

func (d DummyConnectionID) String() string {
	return d.ID
}
