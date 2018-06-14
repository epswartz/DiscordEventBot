package commands

import(
	"DiscordEventBot/db"
	"DiscordEventBot/config"
)

// Deletes an event.
func Delete(server string, sender string, args []string) (string, error) {

	/*
	usageString := "**Usage:** `!e delete <event name>`" // TODO get the command trigger

	// TODO check if args is nil
	// Function for checking argument validity.
	argsValid := func(args []string) bool {
		if len(args) != 1 { // Check number of args
			return false
		}
		return true
	}
	if !argsValid(args) {
		return usageString, nil
	}
	if !alphanum.MatchString(args[0]) || len(args[0]) > config.Cfg.MaxEventNameLength { // Check event name argument
		return "**Error:** Event not found"
	}

	// TODO admin check

	e, err := db.GetEventByName(server, args[0])
	*/

	return ":x: **This function not yet implemented.**", nil
}
