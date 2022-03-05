package controller

import (
	"github.com/Pauloo27/vcserver/internal/models"
	"github.com/Pauloo27/vcserver/internal/vcserver"
)

func Connect(id models.AbstractConnectionID) error {
	return vcserver.Instance.Target.Connect(id)
}
