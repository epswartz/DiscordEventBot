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

import ("flag"
        "os"
        "os/signal"
        "syscall"
        "github.com/tkanos/gonfig"
        "github.com/fatih/color"
        "github.com/bwmarrin/discordgo"
        "github.com/op/go-logging")

var log = logging.MustGetLogger("DiscordEventBot") // Create logger
var format = logging.MustStringFormatter( // Format string for the logger.
    // `%{color}%{time:15:04:05.000} %{shortfunc} %{level:.4s} %{id:03x}%{color:reset} %{message}`,
    `%{color}[%{time:2006-01-02T15:04:05}] %{level:.4s} %{id:03x} > %{color:reset} %{message} `,
)

// Define cmd line options and read them in
var configFilePath = flag.String("config", "./config.json", "Configuration file for bot")

// Have to define a config type for gonfig to hold our config properties.
type Config struct {
    Token    string // The auth token for connecting to discord.
}

// Main message handler. Called for every message that the bot sees, in any channel it has access to.
func handleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

    // Ignore all messages created by the bot itself
    if m.Author.ID == s.State.User.ID {
        return
    }

    // Status update
    if m.Content == "!e status" {
        s.ChannelMessageSend(m.ChannelID, "**EventBot is running.**")
    }

    // TODO look for a command trigger (set in config file, defaults are "!event" and "!e")
    // TODO if command trigger found, send to other functions based on a list of mappings from command name to handler.
}

func main() {

    flag.Parse() // Parse args

    // Initialize the logger's backend so it has somewhere to send messages.
    loggingBackend := logging.NewLogBackend(os.Stdout, "", 0)   // I personally just send everything to stdout.
    loggingBackendFormatter := logging.NewBackendFormatter(loggingBackend, format) // Format using the string we made earlier.
    logging.SetBackend(loggingBackendFormatter) // Set the backends to be used by the logger.

    // Read the config file.
    log.Notice("Attempting to load config file: ", *configFilePath)
    config := Config{}
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
    if err != nil { // Check for stinkies
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






