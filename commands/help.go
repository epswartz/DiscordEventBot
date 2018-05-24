package commands

import "strings"

func argsValid(args []string) bool{
	// TODO not sure how it would be valid.
	return true
}

var usageString = "**Usage:** `!e help [topic (optional)]`\n" // TODO get the command trigger

// Help gives help on the specific command given to it, or it can just print out all of them.
func Help(sender string, args []string) string {
	// If not valid args, give usage
	if(!argsValid(args)){
		return usageString
	}



	return strings.Join("Args recieved: " + len(args), args, " ")
}
