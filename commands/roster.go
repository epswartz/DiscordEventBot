package commands

import(
	"DiscordEventBot/db"
	"strings"
)

// Lists the people who have responded yes or no to an event.
func Roster(server string, args []string) (string, error) {
	var blank db.Event

	usageString := "**Usage:** `!e roster <event name>`" // TODO get the command trigger
	notFoundString := "**Error:** Event not found for this server"

	// Function for checking argument validity.
	argsValid := func(args []string) bool {
		if(len(args) != 1 || strings.ContainsAny(args[0], "/\\")){ // can't have more than one arg or any linux filename reserved chars.
			return false
		}
		return true
	}
	if !argsValid(args) {
		return usageString, nil
	}

	e, err := db.GetEventByName(server, args[0])

	if err != nil {
		return "", err
	}
	if e == blank {
		return notFoundString, nil
	}
	ret := ""

	// TODO Do I want to do some sort of sorting here?
	for _, user := range e.Roster { // For each user that has responded we print out the status in the roster.
		if(user.Status){
			ret += ":white_check_mark:" + user.Id + "\n" // TODO get the user's name instead of ID
		}else{
			ret += ":x:" + user.Id + "\n" // TODO get the user's name instead of ID
		}
	}

	return notFoundString, nil
}
