package dummy

import (
	"time"

	"github.com/Pauloo27/vcserver/internal/models"
)

type DummySearchResultItem struct {
	Title        string
	Artist       string
	Description  string
	Duration     time.Duration
	Source       *models.AbstractSource
	MediaFactory func(DummySearchResultItem) *models.AbstractMedia
}

var _ models.AbstractSearchResultItem = DummySearchResultItem{}

func (d DummySearchResultItem) GetTitle() string {
	return d.Title
}

func (d DummySearchResultItem) GetArtist() string {
	return d.Artist
}

func (d DummySearchResultItem) GetDescription() string {
	return d.Description
}

func (d DummySearchResultItem) GetDuration() time.Duration {
	return d.Duration
}

func (d DummySearchResultItem) GetSource() *models.AbstractSource {
	return d.Source
}

func (d DummySearchResultItem) AsMedia() *models.AbstractMedia {
	if d.MediaFactory == nil {
		return nil
	}
	return d.MediaFactory(d)
}
