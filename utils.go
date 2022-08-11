package utilities

import (
	"fmt"

	"github.com/andersfylling/disgord"
)

func Contains(arr []interface{}, value interface{}) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}

	return false
}

func DiscordIDToString(ID disgord.Snowflake) string {
	return fmt.Sprint(ID)
}
