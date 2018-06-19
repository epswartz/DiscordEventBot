package commands

import(
	"DiscordEventBot/db"
	"DiscordEventBot/config"
	"regexp"
	"time"
	"strconv"
	"reflect"
)

// Creates a new event.
func Create(server string, sender string, args []string) (string, error) {
	var blank db.Event

	serverLocationString := "US/Eastern" // TODO get this from server settings.
	serverLocation, err := time.LoadLocation(serverLocationString)
	if err != nil {
		return "", err
	}
	alphanum := regexp.MustCompile("^[a-zA-Z0-9_]*$") // RegEx for checking if event name is alphanumeric w/ underscores
	dateRegEx := regexp.MustCompile("^[0-9][0-9]/[0-9][0-9]/[0-9][0-9][0-9][0-9]@[0-9][0-9]:[0-9][0-9]$")

	// Different date formats because one looks good when printed and one works well in the cmd
	dateCommandLayout := "01/02/2006@15:04" // TODO Read the actual timezone for the server and concat it.
	datePrintLayout := "Monday, January 2 2006 3:04 PM MST"

	usageString := "**Usage:** `!e create <event_name> [optional scheduled time (MM/DD/YYYY@HH:MM)]`\n Note: Event names are one word." // TODO get the command trigger
	incorrectDateString := "**Error:** Incorrect Date format. Use `MM/DD/YYYY@HH:MM` with 24 hour time and include any leading 0s." // Needed a more intuive err for this one.

	// Function for checking argument validity.
	argsValid := func(args []string) bool {
		if len(args) > 2 || len(args) == 0 { // Check number of args
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
	if len(args) == 2 { // If there's a date, check that too
		if !dateRegEx.MatchString(args[1]) {
			return incorrectDateString, nil
		}
	}

	ev, err := db.GetEventByName(server, args[0])
	if err != nil {
		return "", err
	}
	if !reflect.DeepEqual(ev, blank) {
		return "**Error:** Event `" + args[0] + "` already exists", nil
	}

	// Once we get here we know that the args are valid and there isn't already an event with that name, so we actually create an event.
	var e db.Event

	// Fill out the event struct
	e.Name = args[0]
	e.Server = server



	var t time.Time
	if len(args) == 2 {
		// Find epoch.
		t, err = time.ParseInLocation(dateCommandLayout, args[1], serverLocation) // TODO get actual timezone
		epoch := t.Unix()
		// epoch += 18000 // TODO this number depends on the server timezone. This is for EST.
		e.Epoch = epoch
	} else {
		e.Epoch = -1
	}
	e.Creator = sender
	e.Roster = []db.EventUser{} // Empty slice of EventUser

	err = db.WriteEvent(e)
	if err != nil {
		return "", err
	}

	if len(args) == 1 {
		return "Event `" + args[0] + "` created.", nil
	}
	return "Event `" + args[0] + "` created for `" + t.Format(datePrintLayout) + "`", nil
}
