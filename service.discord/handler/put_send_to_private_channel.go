package handler

import (
	"context"
	"time"

	"github.com/monzo/terrors"

	"github.com/sashajdn/sasha/service.discord/client"
	"github.com/sashajdn/sasha/service.discord/dao"
	"github.com/sashajdn/sasha/service.discord/domain"
	discordproto "github.com/sashajdn/sasha/service.discord/proto"
)

// SendMsgToPrivateChannel ...
func (s *DiscordService) SendMsgToPrivateChannel(
	ctx context.Context, in *discordproto.SendMsgToPrivateChannelRequest,
) (*discordproto.SendMsgToPrivateChannelResponse, error) {
	errParams := map[string]string{
		"idempotency_key": in.IdempotencyKey,
		"user_id":         in.UserId,
		"sender_id":       in.SenderId,
	}

	// First lets check if the idempotency key exists in persistent storage.
	_, exists, err := dao.Exists(ctx, in.IdempotencyKey)
	if err != nil {
		return nil, terrors.Augment(err, "Failed to read existing; dao failed to read", errParams)
	}
	switch {
	case exists && !in.Force:
		return &discordproto.SendMsgToPrivateChannelResponse{}, nil
	}

	// Send message via discord.
	if err := client.SendPrivateMessage(ctx, in.Content, in.UserId); err != nil {
		return nil, terrors.Augment(err, "Failed to send message via discord.", errParams)
	}

	// If the touch doesn't exist or the sender wants to force through an update; then we set via the dao.
	switch {
	case !exists && in.IdempotencyKey != "":
		if _, err := (dao.Create(ctx, &domain.Touch{
			IdempotencyKey: in.IdempotencyKey,
			SenderID:       in.SenderId,
			Updated:        time.Now(),
		})); err != nil {
			// We do have the case whereby the write fails but we still send the message; this is preferable
			// to persisting the idempotency key, but failing to send.
			// We can take the hit of duplicate messages.
			return nil, terrors.Augment(err, "Failed to create touch.", errParams)
		}
	default:
		if _, err := (dao.Update(ctx, &domain.Touch{
			IdempotencyKey: in.IdempotencyKey,
			SenderID:       in.SenderId,
			Updated:        time.Now(),
		})); err != nil {
			// We have the same case as above here too.
			return nil, terrors.Augment(err, "Failed to update touch.", errParams)
		}
	}

	return &discordproto.SendMsgToPrivateChannelResponse{}, nil
}
