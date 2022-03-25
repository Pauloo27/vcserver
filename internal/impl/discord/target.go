package discord

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/Pauloo27/logger"
	"github.com/Pauloo27/vcserver/internal/models"

	dc "github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/diamondburned/arikawa/v3/voice"
)

type DiscordData struct {
	*state.State
}

type DiscordTarget struct {
	data *DiscordData
}

var _ models.AbstractTarget = DiscordTarget{}

func NewDiscordTarget() models.AbstractTarget {
	return &DiscordTarget{data: &DiscordData{}}
}

func (t DiscordTarget) Start() error {
	logger.Info("Starting Discord target...")
	token := os.Getenv("DISCORD_TOKEN")
	s, err := state.New("Bot " + token)
	if err != nil {
		return err
	}
	t.data.State = s
	t.data.AddIntents(gateway.IntentGuildVoiceStates)
	t.data.AddIntents(gateway.IntentGuilds)
	return t.data.Open(context.Background())
}

func (t DiscordTarget) Connect(id models.AbstractConnectionID) error {
	discordID, ok := id.(DiscordConnectionID)
	if !ok {
		ids := strings.Split(id.String(), " ")
		if len(ids) != 2 {
			return errors.New("invalid id")
		}
		discordID = DiscordConnectionID{
			GuildID:   ids[0],
			ChannelID: ids[1],
		}
	}
	vs, err := voice.NewSession(t.data.State)
	if err != nil {
		return err
	}
	guildSf, err := dc.ParseSnowflake(discordID.GuildID)
	if err != nil {
		return err
	}
	channelSf, err := dc.ParseSnowflake(discordID.ChannelID)
	if err != nil {
		return err
	}
	return vs.JoinChannel(dc.GuildID(guildSf), dc.ChannelID(channelSf), false, false)
}

func (t DiscordTarget) Write([]byte) (int, error) { return 0, nil }
func (t DiscordTarget) Disconnect() error         { return nil }

func (t DiscordTarget) Stop() error {
	logger.Info("Closing Discord target...")
	return t.data.Close()
}
