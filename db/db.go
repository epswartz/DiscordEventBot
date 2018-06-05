// This package handles all the interaction with storage for the bot.
// Right now I'm just having it work with local files, but it will use a proper DB at some point.

package db

import (
	// "DiscordEventBot/config"
	// "DiscordEventBot/log"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"time"
)

// Settings document for SMS. A list of people, their number/provider, and whether they have it on or off.
type SMS struct {
	Users []SMSUser
}


// SMS settings for a single user
type SMSUser struct {
	Id string `json:"id"`
	Server string `json:"server"`
	Number string `json:"number"`
	Provider string `json:"provider"`
	Enabled bool `json:"enabled"`
}


// List of admin discord ids.
type Admins struct {
	Ids []string `json: "ids"`
}

// Data for a single event
type Event struct {
	Name string `json:"name"`
	Server string `json:"server"`
	StartTime Time `json:"startTime"`
	Creator string `json:""`
	Server string `json:""`
	Roster []EventUser `json:"roster"`
}

// A single user inside the Roster object.
type EventUser struct {
	Id string `json:"id"`
	Status bool	`json:"status"`
}

func init(){

	// TODO Check that the data directory exists and that we can read it.

	// log.Notice("Connecting to DB")
	theDoc := TestDocument{
		Id: 0,
		Title: "ethan_birthday",
	}

	jsonDoc, _ := json.Marshal(theDoc)
	ioutil.WriteFile("output.json", jsonDoc, 0664)
}


// TODO complete below functions. Names should explain what they do.

func getEventsByTime(server string, time string){

}

func getEventByName(server string, name string){

}

func writeEventByName(server string, name string, doc Event){

}

func getSMS(server string, userid string){

}

func writeSMS(server string, userid string, doc SMS){

}

