package db

import "os"

// Set up data needed when a server is first added. Does nothing if it already exists.
// Handles both events dir and server settings file.
func InitServer(server string) error {
	dirPath := "data/servers/" + server

	// Check the events dir for existence
	if _, err := os.Stat(dirPath); os.IsNotExist(err) { // If true then path does not exist
		err = os.MkdirAll(dirPath+"/events", os.ModePerm) // Create events dir
		if err != nil {
			return err
		}

		return nil
	} else if err != nil { // Some other wacko err that doesn't mean it doesn't exist
		return err
	}

	// Check the settings file for existence
	if _, err := os.Stat(dirPath + "/settings.json"); os.IsNotExist(err) { // If true then path does not exist
		WriteServerSettings(NewServerSettings(server))
	} else if err != nil { // Some other wacko err that doesn't mean it doesn't exist
		return err
	}

	return nil
}
