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

	alphanum := regexp.MustCompile("^[a-zA-Z0-9_]*$") // RegEx for checking if event name is alphanumeric w/ underscores
	dateRegEx := regexp.MustCompile("^[0-9][0-9]/[0-9][0-9]/[0-9][0-9][0-9][0-9]@[0-9][0-9]:[0-9][0-9]$")

	// Different date formats because one looks good when printed and one works well in the cmd
	dateCommandLayout := "01/02/2006@15:04"
	datePrintLayout := "Monday, January 2 2006 3:04 PM"

	usageString := "**Usage:** `!e create <event name> [optional scheduled time (MM/DD/YYYY@HH:MM)]`" // TODO get the command trigger
	incorrectDateString := "**Error:** Incorrect Date format. Use `MM/DD/YYYY@HH:MM` with 24 hour time and include any leading 0s." // Needed a more intuive err for this one.

	// Function for checking argument validity.
	argsValid := func(args []string) bool {
		if len(args) > 2 || len(args) == 0 { // Check number of args
			return false
		}
		if !alphanum.MatchString(args[0]) || len(args[0]) > config.Cfg.MaxEventNameLength { // Check event name argument
			return false
		}
		return true
	}
	if !argsValid(args) {
		return usageString, nil
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
		t, err = time.Parse(dateCommandLayout, args[1])
		epoch := t.UTC().Unix()
		e.Epoch = strconv.FormatInt(epoch, 10)
	} else {
		e.Epoch = ""
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
