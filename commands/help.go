// Right now all this does is print a long ass string.

package commands

var helpString string = "`!e create <event name> [optional scheduled time (MM/DD/YYYY@HH:MM)]` - Create an event\n`!e delete <event name>` - Delete an event\n`!e get <event name>` - Get info and attendance roster for an event\n`!e list` - List events on this server\n`!e respond <yes/no/maybe> <event name>` - Respond with your status for an event\n`!e status` - Prints a string indicating that the bot is alive, and prints the status of the bot's database connection.\n`!e settings [setting] [optional value]` - Get or set some setting for the current server. Just plain old `!e settings` shows them all.\n`!e time <event name> <time (MM/DD/YYYY@HH:MM)>` - Schedule (or reschedule) a time for an event\n`!e version` - Prints information on the bot's current version.\n"

// TODO Currently all this does is print the static string. Would like if I could give some topic to it.
func Help(args []string) (string, error) {
	return helpString, nil
}
