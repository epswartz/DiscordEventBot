package commands

import "strings"

func argsValid(args []string) bool{
	if(len(args) == 0){
		return false
	}
	// TODO check if valid phone number?
	return true
}

var usageString = "**Usage:** `!e register <Phone Number (XXX-XXX-XXXX)>`\nUse `!e help register` for more info." // TODO get the command trigger

// Registers a user for sms.
func Register(sender string, args []string) string {
	// If not valid args, give usage
	if(!argsValid(args)){
		return usageString
	}



	return strings.Join(args, " ")
}
