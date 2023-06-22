package handler

import (
	"context"

	"github.com/sashajdn/sasha/libraries/gerrors"
	"github.com/sashajdn/sasha/service.discord/client"
	discordproto "github.com/sashajdn/sasha/service.discord/proto"
)

// ReadMessageReactions ...
func (s *DiscordService) ReadMessageReactions(
	ctx context.Context, in *discordproto.ReadMessageReactionsRequest,
) (*discordproto.ReadMessageReactionsResponse, error) {
	switch {
	case in.MessageId == "":
		return nil, gerrors.BadParam("missing_param.message_id", nil)
	case in.ChannelId == "":
		return nil, gerrors.BadParam("missing_param.channel_id", nil)
	}

	errParams := map[string]string{
		"message_id": in.MessageId,
		"channel_id": in.ChannelId,
	}

	reactions, err := client.ReadMessageReactions(ctx, in.MessageId, in.ChannelId)
	if err != nil {
		return nil, gerrors.Augment(err, "failed_to_read_message_reactions", errParams)
	}

	protoReactions := []*discordproto.Reaction{}
	for id, users := range reactions {
		protoReactions = append(protoReactions, &discordproto.Reaction{
			ReactionId: string(id),
			UserIds:    users,
		})
	}

	return &discordproto.ReadMessageReactionsResponse{
		Reactions: protoReactions,
	}, nil
}
