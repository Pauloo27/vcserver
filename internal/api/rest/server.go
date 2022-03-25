package rest

import (
	"errors"
	"os"

	"github.com/Pauloo27/aravia"
	"github.com/Pauloo27/aravia/impl"
	"github.com/Pauloo27/logger"
	"github.com/Pauloo27/vcserver/internal/api/rest/controllers"
)

type RestServer struct {
}

func NewRestServer() RestServer {
	return RestServer{}
}

func (RestServer) Start() error {
	logger.Info("Starting REST API server...")
	bindAddr := os.Getenv("REST_SERVER_BIND")
	if bindAddr == "" {
		return errors.New("invalid value for env REST_SERVER_BIND")
	}

	app := aravia.App{
		Server: impl.NewFiberServer(),
	}
	app.With(
		controllers.ConnectionController{},
		controllers.QueueController{},
		controllers.TrackController{},
	)

	return app.Listen(bindAddr)
}
