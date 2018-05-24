package commands

import(
	"DiscordEventBot/db"
)

// Turns SMS notifications on or off for a user.
func Sms(sender string, args []string) string {

	db.Wack()

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

	// Fetch SMS doc for this server out of couchdb

	// enable/disable user's sms
	// Replace doc in couchdb
	return "SMS args were valid."
}
