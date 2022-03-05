package bootstrap

import (
	"github.com/Pauloo27/vcserver/internal/impl/discord"
	"github.com/Pauloo27/vcserver/internal/vcserver"
)

func loadTarget() {
	vcserver.Instance = vcserver.NewVoiceServer(discord.NewDiscordTarget())
}
