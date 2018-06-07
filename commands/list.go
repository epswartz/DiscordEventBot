package commands

import "DiscordEventBot/db"

// Lists events for this server.
func List(server string) string {
	events := db.GetAllServerEvents(server)

	ret := ""
	for _, e := range events {
		ret += "**" + e.Name + "** - " + e.StartTime + "\n"
    }
	return ret
}
