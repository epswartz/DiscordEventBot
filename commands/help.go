// Right now all this does is print a long ass string.

package commands

var helpString string = "`!e create <event name> [optional scheduled time (MM/DD/YYYY@HH:MM)]` - Create an event\n`!e delete <event name>` - Delete an event\n`!e get <event name>` - Get info and attendance roster for an event\n`!e list` - List events on this server\n`!e respond <yes/no/maybe> <event name>` - Respond with your status for an event\n`!e status` - Prints a string indicating that the bot is alive, and prints the status of the bot's database connection.\n`!e time <event name> <time (MM/DD/YYYY@HH:MM)>` - Schedule (or reschedule) a time for an event\n`!e version` - Prints information on the bot's current version.\n"

// Help gives help on the specific command given to it, or it can just print out all of them.
func Help(args []string) (string, error) {

	return helpString, nil
}
