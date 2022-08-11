package utilities

import "github.com/andersfylling/disgord"

type IListRole struct {
	// Display for each role: <Mention> (`<ID>`)
	MentionAndID []string
	// Display for each role: <Name> (`<ID>`)
	NameAndID []string
}

// NOTE: THIS USES DISCORD TEXT MARKDOWN FORMATTING
//
// NOTE: This ONLY displays 10 roles per "page".
func ListRoles(guild *disgord.Guild) {
	var (
		mentionAndID string
		nameAndID    string

		mentionAndIDPages []string
		nameAndIDPages    []string
		idx               int
	)
	for i, r := range guild.Roles {
		idx = i
		if idx == 0 {
			idx = idx + 1
		}

		mentionAndID = mentionAndID + "\n" + r.Mention() + " (`" + DiscordIDToString(r.ID) + "`)"
		nameAndID = nameAndID + "\n**" + r.Name + "** (`" + DiscordIDToString(r.ID) + "`)"

		switch idx {
		case 10:

		}
	}

}
