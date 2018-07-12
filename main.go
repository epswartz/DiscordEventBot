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
	"DiscordEventBot/config"
	"DiscordEventBot/db"
	"DiscordEventBot/log"
	"DiscordEventBot/monitor"
	"DiscordEventBot/session"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
)

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
		if err == nil {
			return msg
		} else {
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
	case "get":
		msg, err = commands.Get(server, args)
	case "help":
		msg, err = commands.Help(args)
	case "list":
		msg, err = commands.List(server)
	case "remind":
		msg, err = commands.Remind(server, sender, args)
	case "respond":
		msg, err = commands.Respond(server, sender, args)
	case "settings":
		msg, err = commands.Settings(server, sender, args)
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

	if err == nil {
		return msg
	} else {
		return handleError(err)
	}
}

// Main message handler. Called for every message that the bot sees, in any channel it has access to. returns the
func handleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	// TODO Support fancier permissions where only people with X role can use the bot at all. This should be in server settings, checked here.
	// TODO Possibly also support even finer level of control where each command has permissions linked to it - that seems like a bitch tho

	// First thing we do is faff around in the disgordgo object hierarchy to get the server's ID
	c, err := s.Channel(m.ChannelID)
	if err != nil {
		log.Error("Error getting channel instance from message object: %s", err)
		return
	}
	server := c.GuildID

	// Ignore all messages created by the bot itself, and ones that are DMs to the bot (that's what having empty string for server means)
	// Blank content typically means image message but also happens with some integrations stuff I think. We ignore those.
	if m.Author.ID == s.State.User.ID || server == "" || m.Content == "" {
		return
	}

	// The entire thing is case insensitive.
	content := strings.ToLower(m.Content)

	// First split the string by whitespace.
	tokens := strings.Fields(content)

	doCommand := false                            // For tracking whether we are doing anything as a result of this message.
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
	s.ChannelMessageSend(m.ChannelID, handleCommand(server, m.Author.ID, tokens[1:len(tokens)]))

	// TODO look for a command trigger (set in config file, defaults are "!event" and "!e")
	// TODO if command trigger found, send to other functions based on a list of mappings from command name to handler.
}

// A bunch of these are fired when the bot starts, because it "sees" the guilds again.
// This is also fired when a guild goes from unavailable to available from some outage.
// The actual use here is when the bot is added to the server orignally.
// This also handles the case where someone adds the bot user to their server while it's offline. When it comes back online, it will init their data.
func handleGuildCreate(s *discordgo.Session, c *discordgo.GuildCreate) {
	log.Debug("Discovered server:", c.ID)
	err := db.InitServer(c.ID) // This will init the new server in the database, or will do nothing if it's already there.
	if err != nil {
		log.Error("Could not initialize guild: ")
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
	dg, err := session.New("Bot " + config.Cfg.Token) // dg is the DiscordGo object
	if err != nil {                                   // Check for stinkies
		log.Critical("Error initializing bot:\n", err)
		os.Exit(1)
	}

	// Register callback for MessageCreate events
	dg.AddHandler(handleMessage)
	dg.AddHandler(handleGuildCreate)

	// Open a websocket connection to Discord and begin listening
	err = dg.Open()
	if err != nil {
		log.Critical("Error connecting to Discord, check internet connection and token: ", err)
		os.Exit(1)
	}
	log.Info("EventBot successfully connected to Discord")

	// Start up the other part of the bot - the part that monitors for when events start and puts out reeminders
	go monitor.StartEventNotifier()

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
