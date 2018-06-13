package commands

import(
	"DiscordEventBot/db"
	"strings"
	"reflect"
)

// Lists the people who have responded yes or no to an event.
func Roster(server string, args []string) (string, error) {
	var blank db.Event

	usageString := "**Usage:** `!e roster <event name>`" // TODO get the command trigger

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

	// FIXME I think this is coming back as true even when there's an event.
	if reflect.DeepEqual(e, blank) { // TODO There's got to be a better way to figure out if there were no results.
		return "**Error:** Event `" + args[0] + "` not found", nil
	}
	ret := ""

	// TODO Do I want to do some sort of sorting here?
	for _, user := range e.Roster { // For each user that has responded we print out the status in the roster.
		if(user.Status){
			ret += ":white_check_mark:\t" + user.Id + "\n" // TODO get the user's name instead of ID
		}else{
			ret += ":x:\t" + user.Id + "\n" // TODO get the user's name instead of ID
		}
	}

	return ret, nil
}
