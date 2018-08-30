// Run StartEventNotifier() as a goroutine to start the event monitor that sends DMs to people when things are starting.

package monitor

import (
	"DiscordEventBot/db"
	"DiscordEventBot/log"
	"DiscordEventBot/session"
	"time"
)

// We pass the currentTime as well as the time we want to check for here.
// This is to distinguish between situations where we say "Event begins now" vs. "Event begins in an hour".
func checkEvents(searchTime int64, currentTime int64) {
	log.Debug("Checking for events")

	events, err := db.GetEventsByTime(searchTime)

	// If this doesn't work we are boned
	if err != nil {
		log.Error("Could not check events:\n", err)
		return
	}

	// Notice that this never runs if there are no events.
	for i := range events {
		var dmChannelIDs []string

		// First add the creator
		creatorChannel, err := session.Session.UserChannelCreate(events[i].Creator)
		if err != nil {
			log.Error("Could not get event creator DM Channel:\n", err)
			return
		}
		dmChannelIDs = append(dmChannelIDs, creatorChannel.ID)
		for _, r := range events[i].Roster {
			if (r.Status == "yes" || r.Status == "maybe") && r.Id != events[i].Creator {
				dmChannel, err := session.Session.UserChannelCreate(r.Id)
				if err != nil {
					log.Error("Could not get event attendee DM Channel:\n", err)
					return
				}
				dmChannelIDs = append(dmChannelIDs, dmChannel.ID)
			}
		}
		log.Info("Sending", len(dmChannelIDs), "reminder(s) for event:", events[i].Name)

		for _, ch := range dmChannelIDs {

			var m string
			// If it starts now, the message is different.
			if currentTime == events[i].Epoch {
				m = "**Reminder:** Event `" + events[i].Name + "` has begun."
			} else {
				// TODO obviously if it isn't always an hour, this needs to change.
				m = "**Reminder:** Event `" + events[i].Name + "` begins in one hour."
			}
			session.Session.ChannelMessageSend(ch, m)
		}
	}
}

// Starts up the event notifier. A function which fires every minute on the minute checking for events for which reminders and things need to be sent out.
func StartEventNotifier() {
	log.Notice("EventBot notification watcher started")
	for {
		nextTime := time.Now().Truncate(time.Minute)
		nextTime = nextTime.Add(time.Minute)
		time.Sleep(time.Until(nextTime))
		currentTime := time.Now().Truncate(time.Minute)
		epoch := currentTime.Unix()
		checkEvents(epoch, epoch) // Look for events that start now

		// TODO should this always be an hour? Server specific setting? Event specific setting?
		// TODO if you change this, you need to show the starting date instead of just saying "in one hour"
		checkEvents(epoch+3600, epoch) // Look for events that start in an hour
	}
}
