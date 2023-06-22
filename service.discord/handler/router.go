package handler

import (
	discordproto "github.com/sashajdn/sasha/service.discord/proto"
)

// DiscordService implements the service for discord.
type DiscordService struct {
	discordproto.UnimplementedDiscordServer
}
