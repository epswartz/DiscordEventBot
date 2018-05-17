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
	"github.com/epswartz/DiscordEventBot/pkg/commands"
	"flag"
	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
	"github.com/op/go-logging"
	"github.com/tkanos/gonfig"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var log = logging.MustGetLogger("DiscordEventBot") // Create logger
var format = logging.MustStringFormatter(          // Format string for the logger.
	// `%{color}%{time:15:04:05.000} %{shortfunc} %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	`%{color}[%{time:2006-01-02T15:04:05}] %{level:.4s} %{id:03x} > %{color:reset} %{message} `,
)

// Define cmd line options and read them in
var configFilePath = flag.String("config", "./config.json", "Configuration file for bot")

// Have to define a config type for gonfig to hold our config properties.
type Config struct {
	Token    string   // The auth token for connecting to discord.
	Triggers []string // Slice of command triggers.
}

// Config object needs to be global, so go ahead and declare it
var config = Config{}

// Returns error message for a bad command that the user tried to use
func invalidCommand(badCmd string) string {
	return ("`" + config.Triggers[0] + " " + badCmd + "` is not a valid command. Use `" + config.Triggers[0] + " help` for more information.")
}

// Command handler. Once we get in here, we know that the message started with the command trigger, and we have the remaining tokens.
// Returns the response to be printed out in the chat.
func handleCommand(tokens []string) string {

	// If there's no command, just print the help.
	if len(tokens) == 0 {
		return commands.Help(nil)
	}

	args := tokens[1:len(tokens)] // Remaining tokens. Command will get these. Guaranteed not empty by my previous if statement.

	// Now we go through our list of commands. when we find it, pass the rest of the tokens as the command's args.
	// Try to keep this in alphabetical order - it'll help
	switch tokens[0] {
	case "help":
		return commands.Help(args)
	case "status":
		return commands.Status(args)
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
	s.ChannelMessageSend(m.ChannelID, handleCommand(tokens[1:len(tokens)]))

	// TODO look for a command trigger (set in config file, defaults are "!event" and "!e")
	// TODO if command trigger found, send to other functions based on a list of mappings from command name to handler.
}

func main() {

	flag.Parse() // Parse args

	// Initialize the logger's backend so it has somewhere to send messages.
	loggingBackend := logging.NewLogBackend(os.Stdout, "", 0)                      // I personally just send everything to stdout.
	loggingBackendFormatter := logging.NewBackendFormatter(loggingBackend, format) // Format using the string we made earlier.
	logging.SetBackend(loggingBackendFormatter)                                    // Set the backends to be used by the logger.

	// Read the config file.
	log.Notice("Attempting to load config file: ", *configFilePath)
	err := gonfig.GetConf(*configFilePath, &config)
	if err != nil { // Check for stinkies
		log.Critical("Error reading config file: ", err)
		return
	}
	log.Notice("Successfully loaded read config file: ", *configFilePath)

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
