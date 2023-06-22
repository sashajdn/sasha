package marshaling

import (
	"github.com/sashajdn/sasha/service.discord/domain"
	discordproto "github.com/sashajdn/sasha/service.discord/proto"
)

// RolesToProto ...
func RolesToProto(roles []*domain.Role) []*discordproto.Role {
	protoRoles := []*discordproto.Role{}
	for _, role := range roles {
		protoRoles = append(protoRoles, &discordproto.Role{
			RoleId:   role.ID,
			RoleName: role.Name,
		})
	}

	return protoRoles
}

// RolesProtoToDomain ...
func RolesProtoToDomain(protos []*discordproto.Role) []*domain.Role {
	roles := []*domain.Role{}
	for _, p := range protos {
		roles = append(roles, &domain.Role{
			Name: p.RoleName,
			ID:   p.RoleId,
		})
	}

	return roles
}
