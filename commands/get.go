package commands

import(
	"DiscordEventBot/db"
	"DiscordEventBot/session"
	"DiscordEventBot/config"
	"strings"
	"reflect"
)

// Lists the people who have responded yes or no to an event.
func Get(server string, args []string) (string, error) {
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

	// Get current data for crator so that we can print their current username
	creatorData, err := session.Session.User(e.Creator) // We need to get the user's current username (usernames change on Discord, but IDs dont)
	if err != nil{
		return "", err
	}

	ret := "**Event:** `" + args[0] + "`\n"
	ret += "**Creator:** `" + creatorData.Username + "#" + creatorData.Discriminator + "`\n"
	ret += "**Time:** `" + e.TimeString + "`\n"

	if len(e.Roster) != 0 {
		ret += "**Attendance Roster:**\n"

		// TODO Do I want to do some sort of sorting here?
		for _, user := range e.Roster { // For each user that has responded we print out the status in the roster.
			userData, err := session.Session.User(user.Id) // We need to get the user's current username (usernames change on Discord, but IDs dont)
			if err != nil{
				return "", err
			}
			if(user.Status){
				ret += ":white_check_mark:\t" + userData.Username + "#" + userData.Discriminator + "\n"
			}else{
				ret += ":x:\t" + userData.Username + "#" + userData.Discriminator + "\n"
			}
		}
	} else {
		ret += "No **Attendance Roster** for this event, use `" + config.Cfg.Triggers[0] + " respond <yes/no> " + args[0] + "` to respond.\n"
	}

	return ret, nil
}
