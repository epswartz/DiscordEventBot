// Keeps the discordgo session so that we can get to it from inside commands without passing it around.
// Discordgo doesn't really have a way to get the current session , so I made this to keep it so I don't have to make the bot log in again.

package session

import(
	"github.com/bwmarrin/discordgo"
)

// We keep the session.
var Session *discordgo.Session

// An alias for discordgo.New which saves what we get back.
func New(initString string) (*discordgo.Session, error){
	Session, err := discordgo.New(initString)
	if(err != nil){
		return Session, err
	}
	return Session, nil
}