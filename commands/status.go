package commands

import "DiscordEventBot/db"

// Status doesn't actually use the args, but it does take them, as does every other handler
/*
func Status() string {
	return ":white_check_mark: **EventBot is Online.**"
}
*/


// I just use this version for testing shit.
func Status() string {
	e := db.GetEventsByTime("2018-06-07@14:30")

	if len(e) == 0 {
		return "No events at the time"
	}

	ret := ""
	for _, ev := range e {
		ret += ev.Name + " "
    }
    return ret
}