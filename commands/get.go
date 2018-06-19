package commands

import(
	"DiscordEventBot/db"
	"DiscordEventBot/session"
	"DiscordEventBot/config"
	"strings"
	"reflect"
	"regexp"
	"time"
	"strconv"
)

// Lists the people who have responded yes or no to an event.
func Get(server string, args []string) (string, error) {
	var blank db.Event

	serverLocationString := "US/Eastern" // TODO get this from server settings.

	alphanum := regexp.MustCompile("^[a-zA-Z0-9_]*$") // RegEx for checking if event name is alphanumeric w/ underscores


	// Formatting string for printing dates.
	dateLayout := "Monday, January 2 2006 3:04 PM MST"

	usageString := "**Usage:** `!e get <event name>`" // TODO get the command trigger

	noDateSetString := "No time scheduled for this event."

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
	if !alphanum.MatchString(args[0]) || len(args[0]) > config.Cfg.MaxEventNameLength { // Check event name argument
		return "**Error:** Invalid event name - event names are aplhanumeric and less than " + strconv.Itoa(config.Cfg.MaxEventNameLength) + " characters", nil
	}

	e, err := db.GetEventByName(server, args[0])

	if err != nil {
		return "", err
	}

	// FIXME I think this is occasionally coming back as true even when there's an event.
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

	// Parse the time
	var timeString string
	if e.Epoch != -1 {
	    if err != nil {
	        return "", err
	    }
	    t := time.Unix(e.Epoch, 0)
	    loc, err := time.LoadLocation(serverLocationString) // TODO Obviously that string changes per the server timezone
	    if err != nil {
	    	return "", err
	    }
	    timeString = t.In(loc).Format(dateLayout)
	} else {
		timeString = noDateSetString
	}
	ret += "**Time:** `" + timeString + "`\n"

	if len(e.Roster) != 0 {
		ret += "**Attendance Roster:**\n"
		yes := ""
		no := ""
		maybe := ""

		// TODO Do I want to do some sort of sorting here?
		for _, user := range e.Roster { // For each user that has responded we print out the status in the roster.
			userData, err := session.Session.User(user.Id) // We need to get the user's current username (usernames change on Discord, but IDs dont)
			if err != nil{
				return "", err
			}
			if user.Status == "yes" {
				yes += ":white_check_mark:\t`" + userData.Username + "#" + userData.Discriminator + "`\n"
			} else if user.Status == "no" {
				no += ":x:\t`" + userData.Username + "#" + userData.Discriminator + "`\n"
			} else { // It's "maybe"
				maybe += ":question:\t`" + userData.Username + "#" + userData.Discriminator + "`\n"
			}
		}
		ret += (yes + maybe + no) // Basically, sort them by putting the yes at the top, then maybe, then no
	} else {
		ret += "No **Attendance Roster** for this event, use `" + config.Cfg.Triggers[0] + " respond <yes/no> " + args[0] + "` to respond.\n"
	}

	return ret, nil
}
