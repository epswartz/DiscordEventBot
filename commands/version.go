package commands

// Prints current EventBot version.
func Version() (string, error) {
	// TODO check if args is nil
	return "**Version:** `1.0.0`\n**Codename:** `Noice`\n**Publish Date:** `08/28/2018`", nil
}
