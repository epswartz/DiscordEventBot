// This package handles all the interaction with storage for the bot.
// Right now I'm just having it work with local files, but it will use a proper DB at some point.

package db

import (
	"DiscordEventBot/log"
	"encoding/json"
	"io/ioutil"
	"os"
)

// Settings document for SMS. A list of people, their number/provider, and whether they have it on or off.
type ServerSettings struct {
	Users []SMSUser `json: sms`
	Admins []string `json: admins`
}


// SMS settings for a single user
type SMSUser struct {
	Id string `json:"id"`
	Server string `json:"server"`
	Number string `json:"number"`
	Provider string `json:"provider"`
	Enabled bool `json:"enabled"`
}

// Data for a single event
type Event struct {
	Name string `json:"name"`
	Server string `json:"server"`
	TimeString string `json:"timeString"`
	Epoch string `json: "epoch"`
	Creator string `json: "creator"`
	Roster []EventUser `json:"roster"`
}

// A single user inside the Roster object.
type EventUser struct {
	Id string `json:"id"`
	Status bool	`json:"status"`
}

// TODO complete below functions. Names should explain what they do.

// Gets all events on all servers which begin at a given time
func GetEventsByTime(time string) ([]Event, error) {
	log.Debug("Getting all events at time " + time)
	var events []Event
	dirPath := "data/servers"

	if _, err := os.Stat(dirPath); os.IsNotExist(err) { // if true then path does not exist
		return events, nil // return empty slice.
	}

	serverDir, err := os.Open(dirPath)
	if err != nil { // if true then path does not exist
		return events, err
	}
	defer serverDir.Close()

	serverDirList,_ := serverDir.Readdirnames(0) // Read all the dir names
    for _, dirName := range serverDirList {
    	serverEvents, err := GetServerEventsByTime(dirName, time)
    	if err != nil {
    		return events, err
    	}
    	events = append(events, serverEvents...)
    }
    return events, nil
}

// Gets all events on a given server which begin at a given time
func GetServerEventsByTime(server string, time string) ([]Event, error) {
	log.Debug("Getting all events for server " + server + " at time " + time)
	var events []Event
	dirPath := "data/servers/" + server + "/events/" + time // Find the directory we need

	if _, err := os.Stat(dirPath); os.IsNotExist(err) { // if true then path does not exist
		return events, nil // return empty slice.
	}

	// Open the directory
	eventFiles, err := os.Open(dirPath)
	if err != nil {
        return events, err
    }
    defer eventFiles.Close()

    fileNameList,_ := eventFiles.Readdirnames(0) // Read all the file names in there
    for _, fileName := range fileNameList {
    	var e Event
    	rawEvent, err := ioutil.ReadFile(dirPath + "/" + fileName)
    	if err != nil {
        	return events, err
    	}
    	json.Unmarshal(rawEvent, &e) // Stuff the unmarshalled data into e
    	events = append(events, e)
    }
    return events, nil
}


func GetAllServerEvents(server string) ([]Event, error) {
	log.Debug("Getting all events for server " + server)
	var events []Event
	dirPath := "data/servers/" + server + "/events"// Find the directory we need

	if _, err := os.Stat(dirPath); os.IsNotExist(err) { // if true then path does not exist
		return events, nil // return empty slice.
	}
	// Open the directory
	eventDirs, err := os.Open(dirPath)
	if err != nil {
        return events, err
    }
	defer eventDirs.Close()

    eventDirNames,_ := eventDirs.Readdirnames(0) // Read all the dir names in the server
    for _, eventDirName := range eventDirNames {
		eventFiles, err := os.Open(dirPath + "/" + eventDirName)
		if err != nil {
        	return events, err
    	}
	 	eventFileNames,_ := eventFiles.Readdirnames(0) // Read all the file names in there
    	for _, eventFileName := range eventFileNames {
    		var e Event
	    	rawEvent, err := ioutil.ReadFile(dirPath + "/" + eventDirName + "/" + eventFileName)
	    	if err != nil {
	        	return events, err
	    	}
	    	json.Unmarshal(rawEvent, &e) // Stuff the unmarshalled data into e
	    	events = append(events, e)
    	}
    }
    return events, nil
}

func GetEventByName(server string, name string) (Event, error) {
	log.Debug("Getting event " + name + " on server " + server)
	var e Event
	dirPath := "data/servers/" + server + "/events"

	eventDirs, err := os.Open(dirPath)
		if err != nil {
        	return e, err
    	}
	 	eventDirNames,_ := eventDirs.Readdirnames(0) // Read all the file names in there
    	for _, eventDirName := range eventDirNames {
    		if _, err := os.Stat(dirPath + "/" + eventDirName + "/" + name + ".json"); err == nil { // If true, we found an event file with that name.
		    	rawEvent, err := ioutil.ReadFile(dirPath + "/" + eventDirName + "/" + name + ".json")
		    	if err != nil {
		        	return e, err
		    	}
		    	json.Unmarshal(rawEvent, &e) // Stuff the unmarshalled data into e
		    	return e, nil
			}
    	}
    	// If we get to the bottom and never found the file, we are done - there was no event with that name.
    	return e, nil
}

// Writes event by name to the proper location, creating the file if it does not exist or updating it if it does.
func WriteEventByName(server string, name string, doc Event) error {
	log.Debug("Writing event " + name + " on server " + server)
	dirPath := "data/servers/" + server + "/events"
}

/*
func GetSMS(server string, userid string) SMSUser {

}

func WriteSMS(server string, userid string, doc SMS){

}

func GetAdmins(server string) Admins {

}
*/

