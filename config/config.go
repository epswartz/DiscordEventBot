// The init function here is the very first thing that runs. This package reads cmd line args and stores the config.
// You can't use the log package in here, because it hasn't been initialized yet, because it depends on something in the config, so this one has to come first.

package config

import 	(
	"github.com/tkanos/gonfig"
	"flag"
	"fmt"
	"os"
)

// Define a config type for gonfig to hold our config properties.
type Config struct {
	DBHost string // Host of couchdb
	DBName string // Name of db to use on couchdb
	DBPass string // Pwd for couchdb
	DBPort int // Port of couchdb
	DBTimeout int // Timeout in ms for couchdb calls
	DBUser string // username for couchdb
	Token string   // The auth token for connecting to discord.
	Triggers []string // Slice of command triggers.
	VerboseLogging bool // enable/disable debug log level
}

// in main.go, we pass a pointer to this to fill it up. then outside this package, we can get at it with config.Cfg once we import the package.
var Cfg = Config{}
var FilePath = flag.String("c", "./config.json", "Configuration file for bot") // -c example.json

/*
func validateConfigFile(){
	// TODO
}
*/

func init(){
	flag.Parse()
	err := gonfig.GetConf(*FilePath, &Cfg)
	if err != nil { // Check for stinkies
		fmt.Println("Error reading config file: \n", err)
		os.Exit(1)
	}
	// validateConfigFile() // TODO
}