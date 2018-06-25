package commands

// Prints current EventBot version.
func Version() (string, error) {
	// TODO check if args is nil
	return "**Version:** `1.0.0-beta`\n**Codename:** `Midly Useful`\n**Publish Date:** `06/25/2018`", nil
}
