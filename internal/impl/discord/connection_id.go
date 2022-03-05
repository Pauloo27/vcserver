package discord

import (
	"fmt"

	"github.com/Pauloo27/vcserver/internal/models"
)

type DiscordConnectionID struct {
	GuildID   string
	ChannelID string
}

var _ models.AbstractConnectionID = DiscordConnectionID{}

func NewDiscordConnectionID(guildID, channelID string) DiscordConnectionID {
	return DiscordConnectionID{
		GuildID:   guildID,
		ChannelID: channelID,
	}
}

func (d DiscordConnectionID) String() string {
	return fmt.Sprintf("%s %s", d.GuildID, d.ChannelID)
}
