package commands

// Prints current EventBot version.
func Version() (string, error) {
	// TODO check if args is nil
	return "**Version:** `1.1.0-beta`\n**Codename:** `Pretty Alright I Guess`\n**Publish Date:** `07/12/2018`", nil
}
