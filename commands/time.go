package commands

import (
	"DiscordEventBot/config"
	"DiscordEventBot/db"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Sets and/or changes the time for an event to the time given.
func Time(server string, sender string, args []string) (string, error) {
	var blank db.Event

	serverLocationString := "US/Eastern" // TODO get this from server settings.
	serverLocation, err := time.LoadLocation(serverLocationString)
	if err != nil {
		return "", err
	}

	alphanum := regexp.MustCompile("^[a-zA-Z0-9_]*$") // RegEx for checking if event name is alphanumeric w/ underscores
	dateRegEx := regexp.MustCompile("^[0-9][0-9]/[0-9][0-9]/[0-9][0-9][0-9][0-9]@[0-9][0-9]:[0-9][0-9]$")

	// Formatting strings for dates.
	datePrintLayout := "Monday, January 2 2006 3:04 PM MST"
	dateCommandLayout := "01/02/2006@15:04"

	usageString := "**Usage:** `!e time <event name> <time (MM/DD/YYYY@HH:MM)>`"                                                    // TODO get the command trigger
	incorrectDateString := "**Error:** Incorrect Date format. Use `MM/DD/YYYY@HH:MM` with 24 hour time and include any leading 0s." // Needed a more intuive err for this one.

	// Function for checking argument validity.
	argsValid := func(args []string) bool {
		if len(args) != 2 || strings.ContainsAny(args[0], "/\\") { // can't have more than one arg or any linux filename reserved chars.
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
	if !dateRegEx.MatchString(args[1]) { // Check the date.
		return incorrectDateString, nil
	}

	e, err := db.GetEventByName(server, args[0])

	if err != nil {
		return "", err
	}

	if reflect.DeepEqual(e, blank) { // TODO There's got to be a better way to figure out if there were no results.
		return "**Error:** Event `" + args[0] + "` not found", nil
	}

	var t time.Time
	// Find epoch time.
	t, err = time.ParseInLocation(dateCommandLayout, args[1], serverLocation) // TODO get actual timezone
	epoch := t.Unix()
	// epoch += 18000 // TODO this number depends on the server timezone. This is for EST.
	e.Epoch = epoch // Right here is where we actually change the time. :)

	err = db.WriteEvent(e)

	if err != nil {
		return "", err
	}

	return "Time for event `" + args[0] + "` changed to `" + t.Format(datePrintLayout) + "`", nil
}
