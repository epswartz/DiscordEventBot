package commands

// Status doesn't actually use the args, but it does take them, as does every other handler
func Status() string {
	return ":white_check_mark: **EventBot is Online.**"
}