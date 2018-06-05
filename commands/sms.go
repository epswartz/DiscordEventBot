package commands

import(
)

// Turns SMS notifications on or off for a user.
func Sms(sender string, args []string) string {


	usageString := "**Usage:** `!e sms <on/off>`\nUse `!e help sms` for more information."

	// Function for checking argument validity.
	argsValid := func(args []string) bool {
		// This cmd needs exactly 1 arg.
		if(len(args) != 1){
			return false
		}
		if(args[0] != "on" && args[0] != "off"){
			return false
		}
		return true
	}

	if(!argsValid(args)){
		return usageString;
	}

	// TODO Fetch SMS doc for this server

	// enable/disable user's sms
	// Replace doc in couchdb
	return "SMS args were valid."
}
