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
	"DiscordEventBot/config"
	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)



// Define cmd line options and read them in

// Returns error message for a bad command that the user tried to use
func invalidCommand(badCmd string) string {
	return ("`" + config.Cfg.Triggers[0] + " " + badCmd + "` is not a valid command. Use `" + config.Cfg.Triggers[0] + " help` for more information.")
}

func handleError(e error) string {
	log.Error(e)
	return "Internal Error. Check log or contact `@Exnur#0001` for assistance."
}

// Command handler. Once we get in here, we know that the message started with the command trigger, and we have the remaining tokens.
// Returns the response to be printed out in the chat.
func handleCommand(server string, sender string, tokens []string) string {

	var err error
	var msg string

	// If there's no command, just print the help.
	if len(tokens) == 0 {
		msg, err = commands.Help(nil)
		if(err == nil){
			return msg
		}else{
			return handleError(err)
		}
	}

	log.Info("Attempting Command: " + tokens[0] + "\n\t issuer: " + sender + "\n\t server: " + server)

	args := tokens[1:len(tokens)] // Remaining tokens. Command will get these. Guaranteed not empty by my previous if statement.

	// Now we go through our list of commands. when we find it, pass the rest of the tokens as the command's args, and if needed, the sender.
	// Try to keep this in alphabetical order - it'll help
	// Pass the commands some subset of the sender, the args
	switch tokens[0] {
	case "create":
		msg, err = commands.Create(server, sender, args)
	case "delete":
		msg, err = commands.Delete(server, sender, args)
	case "help":
		msg, err = commands.Help(args)
	case "list":
		msg, err = commands.List(server)
	case "remind":
		msg, err = commands.Remind(server, sender, args)
	case "respond":
		msg, err = commands.Respond(server, sender, args)
	case "roster":
		msg, err = commands.Roster(server, args)
	case "sms":
		msg, err = commands.Sms(server, sender, args)
	case "status":
		msg, err = commands.Status()
	case "time":
		msg, err = commands.Time(server, sender, args)
	case "version":
		msg, err = commands.Version()
	default:
		return invalidCommand(tokens[0]) // Print err message if command not defined
	}

	if(err == nil){
		return msg
	}else{
		return handleError(err)
	}
}

// Main message handler. Called for every message that the bot sees, in any channel it has access to. returns the
func handleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	// First thing we do is faff around in the disgordgo object hierarchy to get the server's ID
	c, err := s.Channel(m.ChannelID)
	if err != nil {
		log.Critical("Error getting channel instance from message object: %s", err)
		os.Exit(1)
	}
	server := c.GuildID

	// Ignore all messages created by the bot itself, and ones that are DMs to the bot (that's what not having a server means)
	if m.Author.ID == s.State.User.ID || server == "" {
		return
	}

	// The entire thing is case insensitive.
	content := strings.ToLower(m.Content)

	// First split the string by whitespace.
	tokens := strings.Fields(content)

	doCommand := false                        // For tracking whether we are doing anything with this message.
	for _, trigger := range config.Cfg.Triggers { // see if the first token is a command trigger, meaning that it's meant for the bot
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
	s.ChannelMessageSend(m.ChannelID, handleCommand(server, m.Author.ID, tokens[1:len(tokens)]))

	// TODO look for a command trigger (set in config file, defaults are "!event" and "!e")
	// TODO if command trigger found, send to other functions based on a list of mappings from command name to handler.
}


// TODO
func checkEvents() {
	// e := db.GetEventsByTime("166015498469769216", "2018-06-07@14:30")
	log.Debug("Checking for events")

	// Current time to the minute
	//currentTime := time.Now().Truncate(time.Minute)
	//dateString := currentTime.Format("2006-01-02@15:04")

	// Form time string in the same syntax as the user puts it in
	//dateString := currentTime.Year() + "-" + currentTime.Month() + "-" + currentTime.Day() + "@" + currentTime.Hour() + ":" + currentTime.Minute()
	//log.Debug(dateString)



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


	// All good programs start their log with some word art.
	color.Cyan("  ___             _   ___      _   ")
	color.Cyan(" | __|_ _____ _ _| |_| _ ) ___| |_ ")
	color.Cyan(" | _|\\ V / -_) ' \\  _| _ \\/ _ \\  _|")
	color.Cyan(" |___|\\_/\\___|_||_\\__|___/\\___/\\__|")

	log.Notice("Successfully loaded config file: ", *config.FilePath)

	// Start Bot
	log.Notice("EventBot attempting to start ...")
	dg, err := discordgo.New("Bot " + config.Cfg.Token) // dg is the DiscordGo object
	if err != nil {                                 // Check for stinkies
		log.Critical("Error initializing bot: ", err)
		os.Exit(1)
	}

	// Register callback for MessageCreate events
	dg.AddHandler(handleMessage)

	// Open a websocket connection to Discord and begin listening
	err = dg.Open()
	if err != nil {
		log.Critical("Error connecting to Discord, check internet connection and token: ", err)
		os.Exit(1)
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
