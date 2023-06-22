package client

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/monzo/slog"
	"github.com/monzo/terrors"
	"github.com/opentracing/opentracing-go"

	"github.com/sashajdn/sasha/libraries/util"
	"github.com/sashajdn/sasha/service.discord/domain"
)

var (
	// TODO: change implementation to use own defined mocks
	DiscordClientID = "discord-client-id"

	isActiveFlag          bool
	discordTestingChannel = "817513133274824715"

	client      DiscordClient
	clientToken string
)

type EmojiID string

func init() {
	clientToken = fmt.Sprintf("%v", util.EnvGetOrDefault("SATOSHI_DISCORD_API_TOKEN", ""))
	v := util.EnvGetOrDefault("DISCORD_TESTING_OVERRIDE", "0")
	if v != "1" {
		isActiveFlag = true
	}
}

// DiscordClient ...
type DiscordClient interface {
	// Send ...
	Send(ctx context.Context, message, channelID string) (*discordgo.Message, error)
	// SendPrivateMessage ...
	SendPrivateMessage(ctx context.Context, message, userID string) error
	// AddHandler ...
	AddHandler(handler func(s *discordgo.Session, m *discordgo.MessageCreate))
	// AddHandlerGuildMemberAdd ...
	AddHandlerGuildMemberAdd(handler func(s *discordgo.Session, u *discordgo.GuildMemberAdd))
	// ReadRoles ...
	ReadRoles(ctx context.Context, userID string) ([]*domain.Role, error)
	// SetRoles set the users roles to the roles passed. It replaces all the roles the user currently has.
	SetRoles(ctx context.Context, userID string, roles []*domain.Role) error
	// ReadMessageReactions returns a map of reactions by emoji id and the list of users who made that reaction.
	ReadMessageReactions(ctx context.Context, messageID, channelID string) (map[EmojiID][]string, error)

	// TODO: Remove below
	Close()
	Ping(ctx context.Context) error
}

// Init initialises the internal discord client.
func Init(ctx context.Context) error {
	c := New(DiscordClientID, clientToken, true)

	if err := c.Ping(ctx); err != nil {
		return terrors.Augment(err, "Failed to establish connection with discord client", nil)
	}

	slog.Info(ctx, "Discord client initialized", nil)

	// Register guild member add handlers to client.
	for id, guildMemberAddHandler := range guildMemberAddRegistry {
		c.AddHandlerGuildMemberAdd(guildMemberAddHandler)
		slog.Info(ctx, "Guild member handler: %s adding to client", id)
	}

	client = c
	return nil
}

// Send sends a message to a given channel`channel_id` via discord.
func Send(ctx context.Context, message, channelID string) (*discordgo.Message, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Send discord message via channel")
	defer span.Finish()
	return client.Send(ctx, message, channelID)
}

// Send sends a private message to a given user `user_id` via discord.
func SendPrivateMessage(ctx context.Context, message, userID string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Send discord message via private channel")
	defer span.Finish()
	return client.SendPrivateMessage(ctx, message, userID)
}

// ReadRoles ...
func ReadRoles(ctx context.Context, userID string) ([]*domain.Role, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Read user roles")
	defer span.Finish()
	return client.ReadRoles(ctx, userID)
}

// SetRoles ...
func SetRoles(ctx context.Context, userID string, roles []*domain.Role) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Set user roles")
	defer span.Finish()
	return client.SetRoles(ctx, userID, roles)
}

// ReadMessage ...
func ReadMessageReactions(ctx context.Context, messageID, channelID string) (map[EmojiID][]string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Read message")
	defer span.Finish()
	return client.ReadMessageReactions(ctx, messageID, channelID)
}
