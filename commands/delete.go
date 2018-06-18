package commands

import(
	"DiscordEventBot/db"
	"DiscordEventBot/config"
	"DiscordEventBot/session"
	"DiscordEventBot/log"
	"reflect"
	"regexp"
)

// Deletes an event.
func Delete(server string, sender string, args []string) (string, error) {
	var blank db.Event // For checking whether the event already exists. If not, the returned event will match this new one.

	alphanum := regexp.MustCompile("^[a-zA-Z0-9_]*$") // RegEx for checking if event name is alphanumeric w/ underscores
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
		return "**Error:** Event not found", nil
	}

	e, err := db.GetEventByName(server, args[0])

	if err != nil {
		return "", err
	}
	if reflect.DeepEqual(e, blank) {
		return "**Error:** Event `" + args[0] + "` not found", nil
	}
	// Before we delete the event we have to know that the person deleting it is allowed to. You have to either be an admin, or the event's creator.


	memberInfo, err := session.Session.GuildMember(server, sender)
	if err != nil {
		return "", err
	}
	log.Error(memberInfo.Roles)

	err = db.DeleteEvent(e)
	if err != nil {
		return "", err
	}
	return "Event `" + args[0] + "` deleted.", nil




	return ":x: **This function not yet implemented.**", nil
}
