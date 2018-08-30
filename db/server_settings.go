package db

import (
	"DiscordEventBot/log"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strconv"
)

// SMS settings for a single user
type SMSUser struct {
	Id       string `json:"id"` // Discord ID
	Server   string `json:"server"`
	Number   string `json:"number"`
	Provider string `json:"provider"`
	Enabled  bool   `json:"enabled"`
}

// A list of valid settings. Kept in here because it needs to match the properties of the ServerSettings object.

// Settings document for SMS. A list of people, their number/provider, and whether they have it on or off.
type ServerSettings struct {
	// SMS             []SMSUser `json: sms` // SMS Settings for server's users
	MessageAdminOnRSVP bool   `json: messageadminonrsvp`
	PrintListOnRSVP    bool   `json: printlistonrsvp`
	Server             string `json: server`
}

func (ss ServerSettings) DisplayString() string {
	ret := "**Settings for this server:**\n"
	ret += "\tMessageAdminOnRSVP: `" + strconv.FormatBool(ss.MessageAdminOnRSVP) + "`\n"
	ret += "\tPrintListOnRSVP: `" + strconv.FormatBool(ss.PrintListOnRSVP) + "`\n"
	return ret
}

// Returns a new ServerSettings containing the default for given server ID
func NewServerSettings(server string) ServerSettings {
	return ServerSettings{true, false, server}
}

// Writes server's settings to file.
func WriteServerSettings(s ServerSettings) error {
	log.Info("Changed Settings for Server " + s.Server)
	dirPath := "data/servers/" + s.Server + "/"
	sJson, err := json.Marshal(s)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(dirPath+"settings.json", sJson, 0664)
	if err != nil {
		return err
	}
	return nil
}

// Gets the settings for a given server.
func GetServerSettings(server string) (ServerSettings, error) {
	log.Debug("Getting server settings on server " + server)
	var s ServerSettings
	filePath := "data/servers/" + server + "/settings.json"
	if _, err := os.Stat(filePath); err == nil { // If true, we found the settings file
		rawEvent, err := ioutil.ReadFile(filePath)
		if err != nil {
			return s, err
		}
		json.Unmarshal(rawEvent, &s) // Stuff the unmarshalled data into e
		if err != nil {
			return s, err
		}
		return s, nil
	}
	return s, errors.New("No settings file found for server " + server)
}

/*
func GetSMS(server string, userid string) (SMSUser, error) {

}

func WriteSMS(server string, userid string, doc SMS) error {

}


*/
