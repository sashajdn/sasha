package handler

import (
	"context"

	"github.com/sashajdn/sasha/libraries/gerrors"
	"github.com/sashajdn/sasha/service.discord/client"
	"github.com/sashajdn/sasha/service.discord/marshaling"
	discordproto "github.com/sashajdn/sasha/service.discord/proto"
)

// ReadUserRoles ...
func (s *DiscordService) ReadUserRoles(
	ctx context.Context, in *discordproto.ReadUserRolesRequest,
) (*discordproto.ReadUserRolesResponse, error) {
	switch {
	case in.UserId == "":
		return nil, gerrors.BadParam("missing_param.user_id", nil)
	}

	errParams := map[string]string{
		"user_id": in.UserId,
	}

	roles, err := client.ReadRoles(ctx, in.UserId)
	if err != nil {
		return nil, gerrors.Augment(err, "failed_to_read_user_roles", errParams)
	}

	protoRoles := marshaling.RolesToProto(roles)

	return &discordproto.ReadUserRolesResponse{
		Roles: protoRoles,
	}, nil
}
