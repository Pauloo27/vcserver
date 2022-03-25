package bootstrap

import (
	"github.com/Pauloo27/vcserver/internal/api/rest"
	"github.com/Pauloo27/vcserver/internal/impl/discord"
	"github.com/Pauloo27/vcserver/internal/models"
	"github.com/Pauloo27/vcserver/internal/vcserver"
)

func loadTarget() {
	// TODO: dynamic load this
	vcserver.Instance = vcserver.NewVoiceServer(
		discord.NewDiscordTarget(),
		[]models.AbstractAPIServer{rest.NewRestServer()},
	)
}
