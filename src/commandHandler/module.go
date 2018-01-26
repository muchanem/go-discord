package nil

import (
	f "../foundation"
	dsg "github.com/bwmarrin/discordgo"
	//	"strconv"
	a "../flagParser"
	"strings"
)

/* # MessageCreate
* The world's bigest switch statment
*
* This is a very big switch statment run commands. It reads all the messages in
* all the servers its in, determines which ones are commands, and then sees
* what in all the commands mean and then takes the appropriate action.
*
* Parameters:
* - s (type *discordgo.Session) | The current running discord session,
*     (discordgo needs that always apparently)
* - m (type *discordgo.Message) | The message thats to be acted upon.
*
* TODO: See if it can be made so it doesn't have to read every single message
*       ever.
*
* TODO: Break this one function up to smaller functions that only run if a user
*       has a certain role
*
* NOTE: Please delegate what the command actually does to a function. This
*       method should only be used to determine what the user is acutally
*       trying to do.
 */
func MessageCreate(s *dsg.Session, m *dsg.MessageCreate) {
	// The message is checked to see if its a command and can be run
	canRunCommand, err := canTriggerBot(s, m.Message)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "**FATAL. ERROR ENCOUNTERED IN PARSING MESSAGE. DETAILS FOLLOW:**\n"+err.Error()+"\n**CRASHING THE BOT.** *Have a good day!*")
		panic(-1)
	}
	if canRunCommand != true {
		return
	}

	// Removing case sensitivity:
	messageSanatized := strings.ToLower(m.Content)

	// The trailing > is cut off the message so the commands can be more easily handled.
	msg := strings.SplitAfterN(messageSanatized, f.MyBot.Prefs.Prefix, 2)
	message := strings.Split(msg[1], " ")
	flags := a.ParseFlags(message)

	// Now the message is run to see if its a valid command.
	switch message[0] {
	case "help":
		help(s, m, message)
	case "ping":
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	case "info":
		s.ChannelMessageSendEmbed(m.ChannelID, getBotInfo())
	case "rolelist":
		run, err := f.HasRole(s, m.Message, "")
		if err != nil {
			errorLong := "Error! Failed to retrieve  Details below:\n```" + err.Error() + "```"
			s.ChannelMessageSend(m.ChannelID, errorLong)
			return
		}
		if run == false && f.MyBot.Users.ReportPermFails == true {
			s.ChannelMessageSend(m.ChannelID, "You do not have permission to use that command.")
		} else if run == false && f.MyBot.Users.ReportPermFails == false {
			return
		} else if run == true {
			roles, err := getRoles(m.Message)
			if err != nil {
				errorLong := "Error! Failed to retrieve  Details below:\n```" + err.Error() + "```"
				s.ChannelMessageSend(m.ChannelID, errorLong)
				return
			}
			s.ChannelMessageSend(m.ChannelID, roles)
		}
	default:
		s.ChannelMessageSend(m.ChannelID, "Sorry, I don't understand.")
	}
}

/* # Check if user can run command
* This switch statment makes sure the bot runs when its triggered and the user has the perms to trigger it.
* Prevents:
* - Bot posted something that would trigger itself, possibly creating an infinite loop
* - Message posted doesn't have the bot's prefix
* - Command was posted in a channel where the bot shouldn't respond to commands
* - Bot whitelists channels and the command was run in a channel not on the whitelist.
* - Users with a blacklisted role from running the bot
*
* NOTE: Users who have "admin" roles (according to the bot's json data)
*       will have the abilityto run commands regardless of any other rules
*
* NOTE: IF THESE CONDITIONS ARE MET THEN NO ERROR WILL BE SENT TO EITHER DISCORD OR LOGGED.
* THIS IS BY DESIGN. DON'T CHANGE IT THINKING I WAS JUST LAZY.
 */
func canTriggerBot(s *dsg.Session, m *dsg.Message) (bool, error) {
	if m.Author.Bot {
		return false, nil
	}

	admin, err := f.HasRole(s, m, "")
	if err != nil {
		return false, err
	}

	switch true {
	case m.Author.ID == s.State.User.ID:
		return false, nil
	case !strings.HasPrefix(m.Content, f.MyBot.Prefs.Prefix):
		return false, nil
	case admin:
		return true, nil
	case f.Contains(f.MyBot.Perms.BlacklistedChannels, m.ChannelID) == true:
		return false, nil
	case f.MyBot.Perms.WhitelistChannels && !f.Contains(f.MyBot.Perms.WhitelistedChannels, m.ChannelID):
		return false, nil
	}
	for _, b := range f.MyBot.Users.BlacklistedRoles {
		blacklisted, err := f.HasRole(s, m, b)
		if err != nil {
			return false, err
		}
		if blacklisted {
			return false, nil
		}
	}
	return true, nil
}

/* # Get server roles
* A g-d impossibility.
*
* Parameters:
* - m (type *discordgo.Message) | The message used for data extraction about
*	the guild and its roles.
*
* Returns:
* - (type string) | A string list of all the roles.
* - (type error)  | Any errors that may have come up.
*
* NOTE: If you print this into a discord chat, it WILL mention @everyone
 */
func getRoles(m *dsg.Message) (string, error) {
	guild, err := f.GetGuild(f.DG, m)
	if err != nil {
		return "", err
	}
	roles, err := f.DG.GuildRoles(guild.ID)
	if err != nil {
		return "", err
	}
	role := "Server role list:\n```\n"
	for _, r := range roles {
		role += "Name: " + r.Name + "; ID: " + r.ID + ";\n"
	}
	role += "```"
	return role, nil
}

// Returns a messageEmbed about the bot; its a function because if it was a
// variable some of the data doesn't work properly.
func getBotInfo() *dsg.MessageEmbed {
	return &dsg.MessageEmbed{
		Author:      &dsg.MessageEmbedAuthor{},
		Color:       0x073642,
		Title:       "SkilBot Information",
		Description: "A list of commands can be brought up with `" + f.MyBot.Prefs.Prefix + "help`.",
		Thumbnail: &dsg.MessageEmbedThumbnail{
			URL:    "https://i.imgur.com/lPTAiFE.png",
			Width:  64,
			Height: 64,
		},
		Fields: []*dsg.MessageEmbedField{
			&dsg.MessageEmbedField{
				Name:   "Version",
				Value:  "Version " + f.MyBot.Prefs.Version + ".",
				Inline: true,
			},
			&dsg.MessageEmbedField{
				Name:   "Github Link",
				Value:  "https://github.com/skilstak/discord-public",
				Inline: true,
			},
		},
	}
}

/* # Get bot help
* Overcomplecated for little good reason
*
* Parameters:
* - s (type *discordgo.Session) The discord session, this function will manage
*	responding to users instead of a message handler.
* - m (type *discordgo.Session) The message (or a sanitized version of the
*	message) to be used if needed by certain options.
* - f (type []*flagParser.Flag) All the flags and modifiers used with the
*	command.
*
* Note that this function handles responding instead of returning a value to
* its parent to be sent out.
*
* Flags:
* -d | Sends the result via dm.
* -t | Sends the result as standard text (opposed to an embed)
* --command $COMMAND | gets help for the $COMMAND
*
* TODO: Make the help not hard coded. Move into json file? Massive refactor for
* v5.0-alpha probably.
 */
func help(s *dsg.Session, m *dsg.Message, f []*a.Flag) {

}
