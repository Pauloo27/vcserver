package controllers

import (
	"github.com/Pauloo27/aravia"
	"github.com/Pauloo27/vcserver/internal/controller"
	"github.com/Pauloo27/vcserver/internal/impl/dummy"
)

type ConnectionController struct {
}

func (ConnectionController) Init() *aravia.ControllerInfo {
	return nil
}

type ConnectBodyInput struct {
	ID string `json:"id" validate:"required"`
}

func (ConnectionController) Post(body ConnectBodyInput) aravia.HttpStatus {
	id := dummy.DummyConnectionID{ID: body.ID}
	if err := controller.Connect(id); err != nil {
		return aravia.StatusInternalServerError
	}
	return aravia.StatusCreated
}
