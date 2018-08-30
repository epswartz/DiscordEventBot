package commands

import (
	"DiscordEventBot/config"
	"DiscordEventBot/db"
	"DiscordEventBot/session"
	"strconv"
)

// Sets and/or changes the time for an event to the time given.
func Settings(server string, sender string, args []string) (string, error) {
	// TODO check if args is nil

	usageString := "**Usage:** `!e settings [setting name] [value]`"
	validSettings := []string{"messageadminonrsvp", "printlistonrsvp"}
	invalidSettingString := "**Invalid setting. Valid settings are:**\n"
	for _, setting := range validSettings {
		invalidSettingString += "\t`" + setting + "`\n"
	}

	guildRoleNames := []string{} // Incidentally we have to loop across all the roles to check admin privileges - we'll form this while we're in there

	// To interact with server settings at all, you have to either be the server owner or have the admin role.
	guild, err := session.Session.Guild(server)
	if err != nil {
		return "", err
	}
	if sender != guild.OwnerID { // If you're the server owner, we're done
		memberInfo, err := session.Session.GuildMember(server, sender)
		if err != nil {
			return "", err
		}
		for i := range guild.Roles { // Stack Overflow says if I just use the index it's faster because using the elements copies them. We'll try it.
			guildRoleNames = append(guildRoleNames, guild.Roles[i].Name) // This is unrelated to the admin check - we will need it later, so form it now
			if guild.Roles[i].Name == config.Cfg.AdminRole {             // TODO Get admin role name from server specific settings. This is just the default
				// At this point i is the index of the admin role, so check if the command sender has that role.
				if !contains(memberInfo.Roles, guild.Roles[i].ID) { // If you don't have the role, you can't touch settings
					return "**Error:** you do not have permission to use settings on this server", nil
				}
			}
		}
	}

	if len(args) == 0 {
		settings, err := db.GetServerSettings(server)
		if err != nil {
			return "", err
		}
		return settings.DisplayString(), nil
	}

	if len(args) == 1 {
		if !contains(validSettings, args[0]) {
			return invalidSettingString, nil
		}

		settings, err := db.GetServerSettings(server)
		if err != nil {
			return "", err
		}

		switch args[0] {
		case "messageadminonrsvp":
			return "Current value: `" + strconv.FormatBool(settings.MessageAdminOnRSVP) + "`", nil
		case "printlistonrsvp":
			return "Current value: `" + strconv.FormatBool(settings.PrintListOnRSVP) + "`", nil
		default:
			return invalidSettingString, nil
		}
	}
	// Check number of args
	if len(args) != 2 { // Check number of args
		return usageString, nil
	}

	settings, err := db.GetServerSettings(server)
	if err != nil {
		return "", err
	}

	// Validate the actual args, the key and value given.
	// Form invalidSettingString from the list of valid settings above
	switch args[0] {
	case "messageadminonrsvp":
		if args[1] != "true" && args[1] != "false" {
			return "Valid values for this setting are: `true` and `false`", nil
		}
		b, err := strconv.ParseBool(args[1])
		if err != nil {
			return "", err
		}
		settings.MessageAdminOnRSVP = b
	case "printlistonrsvp":
		if args[1] != "true" && args[1] != "false" {
			return "Valid values for this setting are: `true` and `false`", nil
		}
		b, err := strconv.ParseBool(args[1])
		if err != nil {
			return "", err
		}
		settings.PrintListOnRSVP = b
	default:
		return invalidSettingString, nil
	}

	err = db.WriteServerSettings(settings)
	if err != nil {
		return "", nil
	}

	return "Setting `" + args[0] + "` changed to `" + args[1] + "`", nil
}
