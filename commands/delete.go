package commands

import (
	"DiscordEventBot/config"
	"DiscordEventBot/db"
	"DiscordEventBot/session"
	"reflect"
	"regexp"
)

// Deletes an event.
func Delete(server string, sender string, args []string) (string, error) {
	var blank db.Event // For checking whether the event already exists. If not, the returned event will match this new one.

	alphanum := regexp.MustCompile("^[a-zA-Z0-9_]*$")    // RegEx for checking if event name is alphanumeric w/ underscores
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

	// Before we delete the event we have to know that the person deleting it is allowed to.
	// You have to either have the admin role, be the server owner, or be the event's creator.
	// I do these checks in order by how fast they are, because we only have to do the others if the first ones fail, so we can spare a couple lookups sometimes.
	if sender != e.Creator { // If you're the creator, we're done
		guild, err := session.Session.Guild(server)
		if err != nil {
			return "", err
		}
		if sender != guild.OwnerID { // If you're the server owner, we're done
			memberInfo, err := session.Session.GuildMember(server, sender)
			if err != nil {
				return "", err
			}
			for i := range guild.Roles { // Stack Overflow says if I just use the index it's faster because using the elements copies them. We'll try it.
				if guild.Roles[i].Name == config.Cfg.AdminRole { // TODO Get admin role name from server specific settings. This is just the default
					// At this point i is the index of the admin role, so check if the command sender has that role.
					if !contains(memberInfo.Roles, guild.Roles[i].ID) { // If you don't have the role, you can't delete it
						return "**Error:** you do not have permission to delete event `" + args[0] + "`", nil
					}
				}
			}
		}
	}

	err = db.DeleteEvent(e)
	if err != nil {
		return "", err
	}
	return "Event `" + args[0] + "` deleted.", nil
}
