# DiscordEventBot 1.1.0-beta
A bot for scheduling and handling events on a discord server. Allows people to manage (create, update, delete, etc) events, post their status, query for details of the event later, and sends reminders about the event via discord DMs.

- [Features](#features)
- [Using the Publicly Deployed Bot](#using-the-publicly-deployed-bot)
- [Deploying Your Own EventBot Instance](#deploying-your-own-eventbot-instance)
- [Upcoming Development](#upcoming-development)

# Features
The bot features some commands which are available to everyone, some commands which are available only to the event creator and the server admins, and some passive features which aren't commands at all (like reminders an hour before start).

`<>` means required parameters. For example `<event name>` could be `ethanbday`

`[]` means optional parameters. For example `[optional scheduled time]` could be `05/30/2018@17:30`

## Commands Available to Everyone
- `!e create <event name> [optional scheduled time (MM/DD/YYYY@HH:MM)]` - Create an event
- `!e get <event name>` - Get info and attendance roster for an event
- `!e help` - Get link to instructions to use bot **Note: this feature is not yet finished**
- `!e list` - List events on this server
- `!e mention <event name>` - Tag everyone who is in the event's attendance roster as a yes or maybe. **Note: this feature is not yet finished**
- `!e respond <yes/no/maybe> <event name>` - Respond with your status for an event
- `!e sms <on/off>` - Subscribe/Unsubscribe from SMS reminders (also requires it having your number in general) **Note: this feature is not yet finished**
- `!e status` - Prints a string indicating that the bot is alive, and prints the status of the bot's database connection.
- `!e version` - Prints information on the bot's current version.

## Commands Available to Event Creators and Server Admins
- `!e delete <event name>` - Delete an event
- `!e remind <event name>` - Send a reminder for an event (in addition to the auto-reminder 1 hr before) **Note: this feature is not yet finished**
- `!e settings [setting] [optional value]` - Get or Set some setting for the current server. just plain old `!e settings` shows them all. **Note: this feature is not yet finished**
- `!e time <event name> <time (MM/DD/YYYY@HH:MM)>` - Schedule (or reschedule) a time for an event

## Other Features
- Reminder when the event starts
- Reminder 1 hour before the event starts

# Using the Publicly Deployed Bot
The bot is not yet publicly deployed. When it is, I will keep a link here that you can use to add it to your server.
- Admins are determined by the role "EventBotAdmin" - create a role with this name and give it to users if you want them to be able to use the admin commands. The actual server owner is also always considered an admin.

# Deploying Your Own EventBot Instance
Haven't written this section yet. For now, just message @Exnur#0001 on discord.

# Upcoming Development
- Refactor some arg checking into a utils package
- Bar creation of events at times which already passed
- Complete commands listed above as incomplete
- Use embeds for better formatting
- Bot info command that shows a beautiful embed similar to [this one](https://cdn.discordapp.com/attachments/460847996431761428/460848388573888541/unknown.png)
- Delete events some amount of time after they happen
- SMS
- Server specific settings
- Delete multiple events in one go by passing more args
- Auto-setup DB for new servers
