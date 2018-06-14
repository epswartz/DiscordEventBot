package commands

import (
	"DiscordEventBot/db"
	"time"
	"strconv"
)

// Lists events for this server.
func List(server string) (string, error) {

	// Currently it is impossible to have invalid args for list cmd.

	dateLayout := "Monday, January 2 2006 3:04 PM MST"
	noDateSetString := "No time scheduled for this event."

	events, err := db.GetAllServerEvents(server)

	if err != nil {
		return "", err
	}

	ret := ""
	for _, e := range events {
		var timeString string
		if e.Epoch != "" {
			unixTime, err := strconv.ParseInt(e.Epoch, 10, 64)
		    if err != nil {
		        return "", err
		    }
		    t := time.Unix(unixTime, 0)
    	    loc, err := time.LoadLocation("EST") // TODO Obviously that string changes per the server timezone
		    if err != nil {
		    	return "", nil
		    }
		    timeString = t.In(loc).Format(dateLayout)
		} else {
			timeString = noDateSetString
		}
		ret += "**" + e.Name + ":** `" + timeString + "`\n"
    }
	return ret, nil
}
