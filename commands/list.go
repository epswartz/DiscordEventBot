package commands

import (
	"DiscordEventBot/db"
)

// Lists events for this server.
func List(server string) (string, error) {

	// Currently it is impossible to have invalid args for list cmd.

	events, err := db.GetAllServerEvents(server)

	if err != nil {
		return "", err
	}

	ret := ""
	for _, e := range events {
		ret += "**" + e.Name + "** " + e.TimeString + "\n"
    }
	return ret, nil
}
