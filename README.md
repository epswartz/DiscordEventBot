# DiscordEventBot
![#f03c15](https://placehold.it/15/f03c15/000000?text=+) `Current Status: Useless WIP`

A bot for scheduling and handling events on a discord server. Will allow people to manage (create, update, delete, etc) events, post their status, query for details of the event later, and will post reminders about the event via discord, and via SMS if you have one registered with the bot.

# Features
The bot features some commands which are available to everyone, some commands which are available only to the event creator and the server admins, and some passive features which aren't commands at all (like reminders an hour before start).

`<>` means required parameters. For example `<event name>` could be `ethanbday`

`[]` means optional parameters. For example `[optional scheduled time]` could be `05/30/2018@17:30`

## Commands Available to Everyone
- `!e create <event name> [optional scheduled time (MM/DD/YYYY@HH:MM)]` - Create an event
- `!e help` - Get link to instructions to use bot
- `!e list` - List events on this server
- `!e respond <yes/no> <event name>` - Respond with your status for an event
- `!e roster <event name>` - Get current attendees status for an event
- `!e sms <on/off>` - Subscribe/Unsubscribe from SMS reminders (also requires it having your number in general)
- `!e status` - Prints a string indicating that the bot is alive, and prints the status of the bot's database connection.
- `!e version` - Prints information on the bot's current version.

## Commands Available to Event Creators and Admins
- `!e delete <event name>` - Delete an event
- `!e remind <event name>` - Send a reminder for an event (in addition to the auto-reminder 1 hr before)
- `!e time <event name> <time (MM/DD/YYYY@HH:MM)>` - Schedule (or reschedule) a time for an event

## Other Features
- Reminder 1 hour before the event starts

## Things I am unsure about
- Descriptions for events
- Using the bot for larger events in the main channel (not sure I want to have those events, though the bot certainly could do it)
- Automatic SMS signup that isn't manual by me (you PM the bot a command w/ your phone number)
- The real-world feasibility of a satisfyingly moral existence
