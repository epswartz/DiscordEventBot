// The one file in the commands package that isn't itself a command. It just has utility functions it it, for use by the other commands.

package commands

// contains method for slices of string.
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
