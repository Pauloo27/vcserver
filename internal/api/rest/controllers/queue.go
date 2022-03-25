package controllers

import "github.com/Pauloo27/aravia"

type QueueController struct {
}

func (QueueController) Init() *aravia.ControllerInfo {
	return &aravia.ControllerInfo{
		Path: "/connections/:id/queue/",
	}
}

func (QueueController) Get() string {
	return "🚶🚶‍♀️🚶🚶‍♂️🚶‍♀️🚶🚶‍♂️🚶‍♀️"
}

func (QueueController) PostShuffle() string {
	return "🔀"
}
