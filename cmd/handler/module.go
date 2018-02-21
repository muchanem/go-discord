package cmd

import (
	dsg "github.com/bwmarrin/discordgo"
	f "github.com/skilstak/go-discord"
	"github.com/skilstak/go-discord/cmd/commands/info"
	"github.com/skilstak/go-discord/cmd/commands/ping"
	"github.com/skilstak/go-discord/dat"
	"strings"
)

var Cmd = map[string]*f.Command{}

/* FOR THE PERSON RUNNING THIS BOT: Adding packages to the command list
* As of now, the bot has no commands set to it so while it may boot up, it
* won't actually do anything. You will need to add the maps of the command
* modules you have imported or made into the main Cmd map. To do this, add
* each of the command's public map[string]*f.Command type into the following
* init statment. 2 commands, `info` and `ping` have already been added to help
* show what you need to do:
 */

func init() {
	for key, value := range ping.Commands {
		Cmd[key] = value
	}
	for key, value := range info.Commands {
		Cmd[key] = value
	}
	//for key, value := range IMPORTNAMEHERE.Commands {
	//        Cmd[key] = value
	//}
}

//----------------------------------------------------------------------------------//

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
		dat.Log.Println(err.Error())
		dat.Log.Println("YO I FOUND SOME REALLY WILD CHANNEL DATA (dms): " + m.ChannelID)
		s.ChannelMessageSend(m.ChannelID, "Hi there, <@"+m.Author.ID+">. I don't know if you meant to trigger me but I'm always in the background reading messages and yours was a little... weird to me. Could you please inform a server mod or admin about this? I'll regurgitate the error I got here:\n```"+err.Error()+"```\n. Hopefully someone can make something of it because I sure can't.")
		return
	}
	if canRunCommand != true {
		return
	}

	// Removing case sensitivity:
	messageSanatized := strings.ToLower(m.Content)

	// The trailing > is cut off the message so the commands can be more easily handled.
	msg := strings.SplitAfterN(messageSanatized, f.MyBot.Prefs.Prefix, 2)
	message := strings.Split(msg[1], " ")

	// Now the message is run to see if its a valid command and acted upon.
	didAThing := false
	for command, action := range Cmd {
		if message[0] == command {
			action.Action(s, m)
			didAThing = true
		}
	}
	if didAThing == false {
		if strings.Contains(m.Message.Content, "@everyone") {
			s.ChannelMessageSend(m.ChannelID, "Sorry <@"+m.Message.Author.ID+">, but I don't understand what you're saying.\nI'm also not going to @'ing everyone over it so don't bother trying.")
		} else if strings.Contains(m.Message.Content, "@here") {
			s.ChannelMessageSend(m.ChannelID, "Sorry <@"+m.Message.Author.ID+">, but I don't understand what you're saying.\nI'm also not going to be tricked into @'ing here over it so don't bother trying.")
		} else {
			s.ChannelMessageSend(m.ChannelID, "Sorry <@"+m.Message.Author.ID+">, but I don't know what you mean by \"`"+m.Message.Content+"`\".")
		}
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
* NOTE: Users who have "admin" roles (according to the bot's json data) or
*       permissions will have the ability to run commands regardless of any
*       other rules
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
		dat.Log.Println(err.Error())
		//return true, err
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
//func help(s *dsg.Session, m *dsg.Message, f []*a.Flag) {

//}
