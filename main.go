/*
        _.---._    /\\
      ./'       "--`\//
    ./     Ethan    o \
   /./\  )______   \__ \
  ./  / /\ \   | \ \  \ \
     / /  \ \  | |\ \  \7
      "     "    "  "
*/

// Author: Ethan Swartzentruber - eswartzen@gmail.com

package main

import (
	"DiscordEventBot/commands"
	"DiscordEventBot/log"
	"flag"
	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
	"github.com/tkanos/gonfig"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)



// Define cmd line options and read them in
var configFilePath = flag.String("c", "./config.json", "Configuration file for bot") // -c example.json
// var verboseLogging = flag.Bool("v", false, "Use for more verbose logging") // -v to turn on debug logLevel


// Have to define a config type for gonfig to hold our config properties.
type Config struct {
	Token string   // The auth token for connecting to discord.
	Triggers []string // Slice of command triggers.
	DBHost string
	DBPort int
	DBTimeout int
	DBUser string
	DBPass string
}

// Config object needs to be global, so go ahead and declare it
var config = Config{}

// Returns error message for a bad command that the user tried to use
func invalidCommand(badCmd string) string {
	return ("`" + config.Triggers[0] + " " + badCmd + "` is not a valid command. Use `" + config.Triggers[0] + " help` for more information.")
}

// Command handler. Once we get in here, we know that the message started with the command trigger, and we have the remaining tokens.
// Returns the response to be printed out in the chat.
func handleCommand(sender string, tokens []string) string {

	// If there's no command, just print the help.
	if len(tokens) == 0 {
		return commands.Help(nil)
	}

	args := tokens[1:len(tokens)] // Remaining tokens. Command will get these. Guaranteed not empty by my previous if statement.

	// Now we go through our list of commands. when we find it, pass the rest of the tokens as the command's args, and if needed, the sender.
	// Try to keep this in alphabetical order - it'll help
	// Pass the commands some subset of the sender, the args
	switch tokens[0] {
	case "create":
		return commands.Create(sender, args)
	case "delete":
		return commands.Delete(sender, args)
	case "list":
		return commands.List(args)
	case "help":
		return commands.Help(args)
	case "remind":
		return commands.Remind(sender, args)
	case "respond":
		return commands.Respond(sender, args)
	case "roster":
		return commands.Roster(args)
	case "sms":
		return commands.Sms(sender, args)
	case "status":
		return commands.Status() // Status needs no args. :)
	case "time":
		return commands.Time(sender, args)
	case "version":
		return commands.Version()
	}
	return invalidCommand(tokens[0]) // Print err message if command not defined
}

// Main message handler. Called for every message that the bot sees, in any channel it has access to. returns the
func handleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// The entire thing is case insensitive.
	content := strings.ToLower(m.Content)

	// First split the string by whitespace.
	tokens := strings.Fields(content)

	doCommand := false                        // For tracking whether we are doing anything with this message.
	for _, trigger := range config.Triggers { // see if the first token is a command trigger, meaning that it's meant for the bot
		if trigger == tokens[0] {
			doCommand = true
			break
		}
	}
	if !doCommand { // If we didn't see the command trigger then we're done, this message isn't a command.
		return
	}

	// From here, we pass the remaining tokens to a command handler which sorts out what command (if any) they are, and executes.
	// TODO add sender id to handleCommand call
	s.ChannelMessageSend(m.ChannelID, handleCommand(m.Author.ID, tokens[1:len(tokens)]))

	// TODO look for a command trigger (set in config file, defaults are "!event" and "!e")
	// TODO if command trigger found, send to other functions based on a list of mappings from command name to handler.
}


// TODO
func checkEvents() {
	log.Debug("Checking for events")
}

// Starts up the event notifier. A function which fires every minute on the minute checking for events for which reminders and things need to be sent out.
func startEventNotifier() {
	log.Notice("EventBot notification watcher started")
	for {
	    nextTime := time.Now().Truncate(time.Minute)
	    nextTime = nextTime.Add(time.Minute)
	    time.Sleep(time.Until(nextTime))
	    checkEvents()
	}
}

func main() {

	flag.Parse() // Parse args

	// Read the config file.
	log.Notice("Attempting to load config file: ", *configFilePath)
	err := gonfig.GetConf(*configFilePath, &config)
	if err != nil { // Check for stinkies
		log.Critical("Error reading config file: ", err)
		return
	}
	log.Notice("Successfully loaded config file: ", *configFilePath)

	// All good programs start their log with some word art.
	color.Cyan("  ___             _   ___      _   ")
	color.Cyan(" | __|_ _____ _ _| |_| _ ) ___| |_ ")
	color.Cyan(" | _|\\ V / -_) ' \\  _| _ \\/ _ \\  _|")
	color.Cyan(" |___|\\_/\\___|_||_\\__|___/\\___/\\__|")

	// Start Bot
	log.Notice("EventBot attempting to start ...")
	dg, err := discordgo.New("Bot " + config.Token) // dg is the DiscordGo object
	if err != nil {                                 // Check for stinkies
		log.Critical("Error initializing bot: ", err)
		return
	}

	// Register callback for MessageCreate events
	dg.AddHandler(handleMessage)

	// Open a websocket connection to Discord and begin listening
	err = dg.Open()
	if err != nil {
		log.Critical("Error connecting to Discord, check internet connection and token: ", err)
		return
	}
	log.Info("EventBot successfully connected to Discord")

	// Start up the other part of the bot - the part that monitors for when events start and puts out reeminders
	go startEventNotifier()

	// Wait here until CTRL-C or other term signal is received.
	// Full disclosure I know nothing about channels as this is my first time using Go
	// Copied from here: https://github.com/bwmarrin/discordgo/blob/master/examples/pingpong/main.go#L43
	log.Notice("EventBot is now running. Press CTRL+C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
