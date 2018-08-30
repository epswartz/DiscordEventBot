package commands

import (
	"DiscordEventBot/config"
	"DiscordEventBot/db"
	"DiscordEventBot/session"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// Responds that you either are or are not going to an event.
func Respond(server string, sender string, args []string) (string, error) {
	var blank db.Event                                                   // Blank event for checking
	usageString := "**Usage:** `!e respond <yes/no/maybe> <event name>`" // TODO get the command trigger
	alphanum := regexp.MustCompile("^[a-zA-Z0-9_]*$")                    // RegEx for checking if event name is alphanumeric w/ underscores

	// Function for checking argument validity.
	argsValid := func(args []string) bool {
		// Need exactly 2 args, and the first one has to either be yes or no.
		if len(args) != 2 || (strings.ToLower(args[0]) != "yes" && strings.ToLower(args[0]) != "no" && strings.ToLower(args[0]) != "maybe") {
			return false
		}
		return true
	}
	if !argsValid(args) {
		return usageString, nil
	}
	if !alphanum.MatchString(args[1]) || len(args[1]) > config.Cfg.MaxEventNameLength { // Check event name argument
		return "**Error:** Invalid event name - event names are aplhanumeric and less than " + strconv.Itoa(config.Cfg.MaxEventNameLength) + " characters", nil
	}

	e, err := db.GetEventByName(server, args[1])
	if err != nil {
		return "", err
	}

	if reflect.DeepEqual(e, blank) { // TODO There's got to be a better way to figure out if there were no results.
		return "**Error:** Event `" + args[1] + "` not found", nil
	}

	senderStatus := strings.ToLower(args[0])

	add := true
	// Look for the sender in the roster. if they're there, change their status.
	for i := range e.Roster {
		if e.Roster[i].Id == sender {
			e.Roster[i].Status = senderStatus
			add = false
		}
	}

	// If they weren't already in the roster, add them to it.
	if add {
		e.Roster = append(e.Roster, db.EventUser{sender, senderStatus})
	}

	err = db.WriteEvent(e)
	if err != nil {
		return "", err
	}

	// Need username for the output
	userData, err := session.Session.User(sender)

	return "`" + userData.Username + "#" + userData.Discriminator + "` status for event `" + args[1] + "` is now `" + senderStatus + "`", nil

}
