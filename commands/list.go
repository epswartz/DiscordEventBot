package commands

import (
	"DiscordEventBot/db"
	"time"
)

// Lists events for this server.
func List(server string) (string, error) {
	// Currently it is impossible to have invalid args for list cmd.
	serverLocationString := "US/Eastern" // TODO get this from server settings.

	dateLayout := "Monday, January 2 2006 3:04 PM MST"
	noDateSetString := "No time scheduled for this event."

	events, err := db.GetAllServerEvents(server)

	if err != nil {
		return "", err
	}

	ret := ""

	for _, e := range events {
		var timeString string
		if e.Epoch != -1 {
			t := time.Unix(e.Epoch, 0)
			loc, err := time.LoadLocation(serverLocationString) // TODO Obviously that string changes per the server timezone
			if err != nil {
				return "", err
			}
			timeString = t.In(loc).Format(dateLayout)
		} else {
			timeString = noDateSetString
		}
		ret += "**" + e.Name + ":** `" + timeString + "`\n"
	}

	if ret == "" {
		ret = "No events found for this server. Use `!e create` to create one."
	}

	return ret, nil
}
