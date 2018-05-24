package commands

/*
// Define several types so that we can work with the SMS doc in cloudant
type SMSDocument struct{
	doctype string
	Contacts
}

type Contacts struct {
	contacts []Contact
}

type Contact struct {
	id string
	number string
	enabled bool
}
*/

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

	// Fetch SMS doc for this server out of couchdb

	// enable/disable user's sms
	// Replace doc in couchdb
	return "SMS args were valid."
}
