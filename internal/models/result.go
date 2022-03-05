package models

import "time"

type AbstractSearchResultItem interface {
	GetTitle() string
	GetArtist() string
	GetDescription() string
	GetDuration() time.Duration
	GetSource() *AbstractSource
	AsMedia() *AbstractMedia
}
