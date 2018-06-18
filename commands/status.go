package commands

// import "DiscordEventBot/db"

// Status just prints out the fact that the thing is running. Eventually it will print a DB connection status, too.
func Status() (string, error) {
	return ":white_check_mark: **EventBot is Online.**", nil
}
