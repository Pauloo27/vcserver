package controllers

import (
	"strings"

	"github.com/Pauloo27/aravia"
)

type TrackController struct {
}

func (TrackController) Init() *aravia.ControllerInfo {
	return &aravia.ControllerInfo{
		Path: "/connections/:id/queue/tracks",
	}
}

func (TrackController) Get() []string {
	return strings.Split(strings.Repeat("🎶", 10), "")
}

func (TrackController) Delete() aravia.HttpStatus {
	return aravia.StatusNoContent
}

func (TrackController) Post() []string {
	return strings.Split(strings.Repeat("🎶", 11), "")
}

func (TrackController) Get_Idx() string {
	return "🎶"
}

func (TrackController) Delete_Idx() aravia.HttpStatus {
	return aravia.StatusNoContent
}
