package db

import (
	"DiscordEventBot/config"
	"DiscordEventBot/log"
	"github.com/rhinoman/couchdb-go"
	"time"
)
/*
// Define several types so that we can work with the SMS doc in cloudant
type SMSDocument struct{
	doctype string
	Contacts
}

type Contacts struct {
	contacts []Contact
}

type Contact struct {
	id string
	number string
	enabled bool
}
*/

type TestDocument struct {
	Title string
	Note string
}

func Wack(){
}

func init(){

	log.Notice("Connecting to DB")
	var timeout = time.Duration(config.Cfg.DBTimeout) * time.Millisecond
	conn, err := couchdb.NewConnection(config.Cfg.DBHost,config.Cfg.DBPort,timeout)
	if(err != nil){
		log.Critical("Failed to connect to couchdb", err)
	}
	auth := couchdb.BasicAuth{Username: config.Cfg.DBUser, Password: config.Cfg.DBPass }
	db := conn.SelectDB(config.Cfg.DBName, &auth)

	theDoc := TestDocument{
		Title: "My Document",
		Note: "This is a note",
	}

	theId := "abc123" //use whatever method you like to generate a uuid
	//The third argument here would be a revision, if you were updating an existing document
	_ , err = db.Save(theDoc, theId, "")
	if(err != nil){
		log.Critical("NOPE", err)
	}
	//If all is well, rev should contain the revision of the newly created
	//or updated Document
}