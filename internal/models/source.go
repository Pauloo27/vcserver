package models

type AbstractSource interface {
	Search(query string) ([]*AbstractSearchResultItem, error)
}
