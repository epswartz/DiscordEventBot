// This package handles all the interaction with storage for the bot.
// Right now I'm just having it work with local files, but it will use a proper DB at some point.

package db

import (
	// "DiscordEventBot/config"
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
	StartTime string `json:"startTime"`
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
func GetEventsByTime(time string) []Event {
	var events []Event
	dirPath := "data/servers"

	if _, err := os.Stat(dirPath); os.IsNotExist(err) { // if true then path does not exist
		return events // return empty slice.
	}

	serverDir, err := os.Open(dirPath)
	if err != nil { // if true then path does not exist
		log.Critical("Failed to open data directory: %s", err)
		os.Exit(1)

	}
	defer serverDir.Close()

	serverDirList,_ := serverDir.Readdirnames(0) // Read all the dir names
    for _, dirName := range serverDirList {
    	events = append(events,GetServerEventsByTime(dirName, time)...)
    }
    return events
}

// Gets all events on a given server which begin at a given time
func GetServerEventsByTime(server string, time string) []Event {
	var events []Event
	dirPath := "data/servers/" + server + "/events/" + time // Find the directory we need

	if _, err := os.Stat(dirPath); os.IsNotExist(err) { // if true then path does not exist
		return events // return empty slice.
	}

	// Open the directory
	eventFiles, err := os.Open(dirPath)
	if err != nil {
        log.Critical("Failed to open data directory: %s", err)
        os.Exit(1)
    }
    defer eventFiles.Close()

    fileNameList,_ := eventFiles.Readdirnames(0) // Read all the file names in there
    for _, fileName := range fileNameList {
    	var e Event
    	rawEvent, err := ioutil.ReadFile(dirPath + "/" + fileName)
    	if err != nil {
        	log.Critical("Failed to open event file: %s", err)
        	os.Exit(1)
    	}
    	json.Unmarshal(rawEvent, &e) // Stuff the unmarshalled data into e
    	events = append(events, e)
    }
    return events
}


func GetAllServerEvents(server string) []Event {
	var events []Event
	dirPath := "data/servers/" + server + "/events"// Find the directory we need

	if _, err := os.Stat(dirPath); os.IsNotExist(err) { // if true then path does not exist
		return events // return empty slice.
	}
	// Open the directory
	eventDirs, err := os.Open(dirPath)
	if err != nil {
        log.Critical("Failed to open data directory: %s", err)
        os.Exit(1)
    }
	defer eventDirs.Close()

    eventDirNames,_ := eventDirs.Readdirnames(0) // Read all the dir names in the server
    for _, eventDirName := range eventDirNames {
		eventFiles, err := os.Open(dirPath + "/" + eventDirName)
		if err != nil {
        	log.Critical("Failed to open data directory: %s", err)
    	}
	 	eventFileNames,_ := eventFiles.Readdirnames(0) // Read all the file names in there
    	for _, eventFileName := range eventFileNames {
    		var e Event
	    	rawEvent, err := ioutil.ReadFile(dirPath + "/" + eventDirName + "/" + eventFileName)
	    	if err != nil {
	        	log.Critical("Failed to open event file: %s", err)
	        	os.Exit(1)
	    	}
	    	json.Unmarshal(rawEvent, &e) // Stuff the unmarshalled data into e
	    	events = append(events, e)
    	}
    }
    return events
}

func GetEventByName(server string, name string) Event {
	var e Event
	dirPath := "data/servers/" + server + "/events"

	eventDirs, err := os.Open(dirPath)
		if err != nil {
        	log.Critical("Failed to open data directory: %s", err)
    	}
	 	eventDirNames,_ := eventDirs.Readdirnames(0) // Read all the file names in there
    	for _, eventDirName := range eventDirNames {
    		if _, err := os.Stat(dirPath + "/" + eventDirName + "/" + name + ".json"); err == nil { // If true, we found an event file with that name.
		    	rawEvent, err := ioutil.ReadFile(dirPath + "/" + eventDirName + "/" + name + ".json")
		    	if err != nil {
		        	log.Critical("Failed to open event file: %s", err)
		        	os.Exit(1)
		    	}
		    	json.Unmarshal(rawEvent, &e) // Stuff the unmarshalled data into e
		    	return e
			}
    	}
    	// If we get to the bottom and never found the file, we are done - there was no event with that name.
    	return e
}

/*
func WriteEventByName(server string, name string, doc Event){

}

func GetSMS(server string, userid string) SMSUser {

}

func WriteSMS(server string, userid string, doc SMS){

}

func GetAdmins(server string) Admins {

}
*/

