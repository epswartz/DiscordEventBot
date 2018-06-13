package commands

import (
	"strings"
	"strconv"
)



// Help gives help on the specific command given to it, or it can just print out all of them.
func Help(args []string) (string, error) {

	usageString := "**Usage:** `!e help [topic (optional)]`" // TODO get the command trigger

	// Function for checking argument validity.
	argsValid := func(args []string) bool {
		// TODO There may eventually be some way for them to be invalid.
		return true
	}
	if(!argsValid(args)){
		return usageString, nil
	}

	if(len(args) == 0){
		return "No args.", nil
	}


	return ("Args recieved: " + strconv.Itoa(len(args)) + "\n" + strings.Join(args, " ")), nil
}
