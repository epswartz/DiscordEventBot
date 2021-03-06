// This package handles all the interaction with storage for the bot.
// Right now I'm just having it work with local files, but it will use a proper DB at some point.

package db

import (
	"DiscordEventBot/log"
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
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
	Epoch int64 `json: "epoch"`
	Creator string `json: "creator"`
	Roster []EventUser `json:"roster"`
}

// A single user inside the Roster object.
type EventUser struct {
	Id string `json:"id"`
	Status string	`json:"status"`
}

// TODO complete below functions. Names should explain what they do.

// Gets all events on all servers which begin at a given time
func GetEventsByTime(epoch int64) ([]Event, error) {
	log.Debug("Getting all events at time " + strconv.FormatInt(epoch, 10))
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
    	serverEvents, err := GetServerEventsByTime(dirName, epoch)
    	if err != nil {
    		return events, err
    	}
    	events = append(events, serverEvents...)
    }
    return events, nil
}

// Gets all events on a given server which begin at a given time
// FIXME This hasn't worked since I changed the time formatting stuff, don't need it until I add reminders.
func GetServerEventsByTime(server string, epoch int64) ([]Event, error) {
	log.Debug("Getting all events for server " + server + " at time " + strconv.FormatInt(epoch, 10))
	var events []Event
	dirPath := "data/servers/" + server + "/events/" // Find the directory we need

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
    	if e.Epoch == epoch { // If the event starts at the requested time
    		events = append(events, e)
    	}
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
	eventFiles, err := os.Open(dirPath)
	if err != nil {
        return events, err
    }
    defer eventFiles.Close()

 	eventFileNames,_ := eventFiles.Readdirnames(0) // Read all the file names in there
	for _, eventFileName := range eventFileNames {
		var e Event
    	rawEvent, err := ioutil.ReadFile(dirPath + "/" + eventFileName)
    	if err != nil {
        	return events, err
    	}
    	json.Unmarshal(rawEvent, &e) // Stuff the unmarshalled data into e
    	events = append(events, e)
	}

    return events, nil
}

func GetEventByName(server string, name string) (Event, error) {
	log.Debug("Getting event " + name + " on server " + server)
	var e Event
	dirPath := "data/servers/" + server + "/events"
	if _, err := os.Stat(dirPath + "/" + name + ".json"); err == nil { // If true, we found an event file with that name.
    	rawEvent, err := ioutil.ReadFile(dirPath+ "/" + name + ".json")
    	if err != nil {
        	return e, err
    	}
    	json.Unmarshal(rawEvent, &e) // Stuff the unmarshalled data into e
    	return e, nil
	}
	// If we never found the file, we are done - there was no event with that name.
	return e, nil
}


// Writes event to the proper location, creating the file if it does not exist or updating it if it does.
func WriteEvent(e Event) error {
	log.Debug("Writing event " + e.Name + " on server " + e.Server)
	dirPath := "data/servers/" + e.Server + "/events/"
	eJson, err := json.Marshal(e)
	if err != nil {
    	return err
	}
	err = ioutil.WriteFile(dirPath + e.Name + ".json", eJson, 0664)
	if err != nil {
		return err
	}
	return nil
}


// Delete an event if it exists.
// Standard use here is to  call GetEventByName, and then call delete, that's why it takes an event.
func DeleteEvent(e Event) error {
	log.Debug("Deleting event " + e.Name + " on server " + e.Server)
	dirPath := "data/servers/" + e.Server + "/events/"
	err := os.Remove(dirPath + e.Name + ".json")
	if err != nil {
		return err
	}
	return nil
}

// Set up data needed when a server is first added. Does nothing if it already exists.
func InitServer(server string) error {
	dirPath := "data/servers/" + server
	if _, err := os.Stat(dirPath); os.IsNotExist(err) { // If true then path does not exist
		err = os.MkdirAll(dirPath + "/events", os.ModePerm); // Create dirs
		if err != nil {
			return err
		}
		return nil
	} else if err != nil { // Some other wacko err that doesn't mean it doesn't exist
		return err
	}
	return nil
}


/*
func GetSMS(server string, userid string) (SMSUser, error) {

}

func WriteSMS(server string, userid string, doc SMS) error {

}

func GetServerSettings(server string) (ServerSettings, error) {

}

func WriteServerSettings(s ServerSettings) error {

}
*/

