package commands

// Prints current EventBot version.
func Version() (string, error) {
	// TODO check if args is nil
	return "**Current EventBot version: 0.0.1 (Codename: Useless)**\nPublish Date: 05/24/2018", nil
}
