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
	return "πΆπΆββοΈπΆπΆββοΈπΆββοΈπΆπΆββοΈπΆββοΈ"
}

func (QueueController) PostShuffle() string {
	return "π"
}
