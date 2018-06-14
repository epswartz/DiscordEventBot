# DiscordEventBot
![#f03c15](https://placehold.it/15/f03c15/000000?text=+) `Current Status: Useless WIP`

- [Features](#features)
- [Using the Publicly Deployed Bot](#using-the-publically-deployed-bot)
- [Deploying Your Own EventBot Instance](#deploying-your-own-eventbot-instance)

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
- `!e sms <on/off>` - Subscribe/Unsubscribe from SMS reminders (also requires it having your number in general) **Note: this feature currently unavailable on the public bot instance**
- `!e status` - Prints a string indicating that the bot is alive, and prints the status of the bot's database connection.
- `!e version` - Prints information on the bot's current version.

## Commands Available to Event Creators and Server Admins
- `!e timezone <timezone>` - Set the server timezone to each
- `!e delete <event name>` - Delete an event
- `!e remind <event name>` - Send a reminder for an event (in addition to the auto-reminder 1 hr before)
- `!e time <event name> <time (MM/DD/YYYY@HH:MM)>` - Schedule (or reschedule) a time for an event

## Other Features
- Reminder 1 hour before the event starts

# Using the Publicly Deployed Bot
`// TODO this section - talk about adding the public bot to servers and setting up the role`
- Admins are determined by the role "EventBotAdmin" - create a role with this name and give it to users if you want them to be able to use the admin commands. the actual server owner is also always considered an admin.

# Deploying Your Own EventBot Instance
`// TODO this section - talk about config file, DB setup (or lack thereof, if using local filesystem), etc`


`// TODO Add example usage for each command`