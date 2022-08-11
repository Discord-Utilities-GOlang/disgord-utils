package utilities

import (
	"errors"

	"github.com/andersfylling/disgord"
)

func SearchRole(guild *disgord.Guild, ID disgord.Snowflake) (error, *disgord.Role) {
	invalidRole := errors.New("The specified ID couldn't be used to find a role as no such role has that ID!")
	for _, v := range guild.Roles {
		if v.ID == ID {
			return nil, v
		}
	}

	return invalidRole, nil
}
