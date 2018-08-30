// The init function here is the very first thing that runs. This package reads cmd line args and stores the config.
// You can't use the log package in here, because it hasn't been initialized yet, because it depends on something in the config, so this one has to come first.

package config

import (
	"flag"
	"fmt"
	"github.com/tkanos/gonfig"
	"os"
)

// Define a config type for gonfig to hold our config properties.
type Config struct {
	Token              string   // The auth token for connecting to discord.
	Triggers           []string // Slice of command triggers.
	VerboseLogging     bool     // enable/disable debug log level
	MaxEventNameLength int      // How long event names can be.
	AdminRole          string   // Default name of the EventBot admin role
}

// in main.go, we pass a pointer to this to fill it up. then outside this package, we can get at it with config.Cfg once we import the package.
var Cfg = Config{}
var FilePath = flag.String("c", "./config.json", "Configuration file for bot") // -c example.json

/*
func validateConfigFile(){
	// TODO
}
*/

func init() {
	flag.Parse()
	err := gonfig.GetConf(*FilePath, &Cfg)
	if err != nil { // Check for stinkies
		fmt.Println("Error reading config file: \n", err)
		os.Exit(1)
	}
	// validateConfigFile() // TODO
}
