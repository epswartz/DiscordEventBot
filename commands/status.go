// A command which says the bot is online (or doesn't, if it turns out the bot isn't online). Will also eventually print database status.

package commands

// Status doesn't actually use the args, but it does take them, as does every other handler
func Status(args []string) string {
	return ":white_check_mark: **EventBot is Online.**"
}