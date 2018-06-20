package commands

// Prints current EventBot version.
func Version() (string, error) {
	// TODO check if args is nil
	return "**Version:** `0.0.1`\n**Codename:** `Useless`\n**Publish Date:** `05/24/2018`", nil
}
