package nil

import (
	f "../foundation"
	dsg "github.com/bwmarrin/discordgo"
	"strconv"
	"strings"
)

func MessageCreate(s *dsg.Session, m *dsg.MessageCreate) ***REMOVED***
	// This switch statment makes sure the bot runs when its triggered and the user has the perms to trigger it.
	// Prevents:
	// - Bot posted something that would trigger itself, possibly creating an infinite loop
	// - Message posted doesn't have the bot's prefix
	// - Command was posted in a channel where the bot shouldn't respond to commands
	// - Bot whitelists channels and the command was run in a channel not on the whitelist.
	//
	// Allows:
	// - Priority "Alpha" users to run commands regardless of any other rules
	// - Checks if priority "Bravo" users can override, lets command run if so.
	//
	// IF THESE CONDITIONS ARE MET THEN NO ERROR WILL BE SENT TO EITHER DISCORD OR LOGGED.
	// THIS IS BY DESIGN. DON'T CHANGE IT THINKING I WAS JUST LAZY.

	switch true ***REMOVED***
	case !strings.HasPrefix(m.Content, f.MyBot.Prefs.Prefix):
		return
	case f.Contains(f.MyBot.Users.AlphaUsers, m.Author.ID):
		break
	case f.MyBot.Users.BravoOverride && f.Contains(f.MyBot.Users.BravoUsers, m.Author.ID):
		break
	case !strings.HasPrefix(m.Content, f.MyBot.Prefs.Prefix):
		return
	case f.Contains(f.MyBot.Perms.BlacklistedChannels, m.ChannelID):
		return
	case f.MyBot.Perms.WhitelistChannels && !f.Contains(f.MyBot.Perms.WhitelistedChannels, m.ChannelID):
		return
	***REMOVED***

	// The trailing > is cut off the message so the commands can be more easily handled.
	message := strings.SplitAfterN(m.Content, ">", 2)
	message = strings.SplitAfterN(message[1], " ", -1)

	// Now the message is run to see if its a valid command.
	switch message[0] ***REMOVED***
	case "help":
		s.ChannelMessageSend(m.ChannelID, "Still working on this--- whoops.")
	case "ping":
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	case "info":
		s.ChannelMessageSendEmbed(m.ChannelID, getBotInfo())
	case "perms":
		s.ChannelMessageSendEmbed(m.ChannelID, getPerms(m.Author.ID, m.Author.AvatarURL("9"), m.Author.Mention()))
	default:
		s.ChannelMessageSend(m.ChannelID, "Sorry, I don't understand.")
	***REMOVED***
***REMOVED***

func getBotInfo() *dsg.MessageEmbed ***REMOVED***
	return &dsg.MessageEmbed***REMOVED***
		Author:      &dsg.MessageEmbedAuthor***REMOVED******REMOVED***,
		Color:       0x073642,
		Title:       "MyBot Information",
		Description: "A list of commands can be brought up with `" + f.MyBot.Prefs.Prefix + "help`.",
		Thumbnail: &dsg.MessageEmbedThumbnail***REMOVED***
			URL:    "https://pbs.twimg.com/media/CLI0Kd5UcAAkpiM.png",
			Width:  64,
			Height: 64,
		***REMOVED***,
		Fields: []*dsg.MessageEmbedField***REMOVED***
			&dsg.MessageEmbedField***REMOVED***
				Name:   "Version",
				Value:  "Version " + f.MyBot.Prefs.Version + ".",
				Inline: true,
			***REMOVED***,
			&dsg.MessageEmbedField***REMOVED***
				Name:   "Github Link",
				Value:  "https://github.com/skilstak/discord-public",
				Inline: true,
			***REMOVED***,
		***REMOVED***,
	***REMOVED***
***REMOVED***

func getPerms(user string, icon string, mention string) *dsg.MessageEmbed ***REMOVED***
	a := f.Contains(f.MyBot.Users.AlphaUsers, user)
	b := f.Contains(f.MyBot.Users.BravoUsers, user)
	c := f.Contains(f.MyBot.Users.CharlieUsers, user)

	return &dsg.MessageEmbed***REMOVED***
		Author:      &dsg.MessageEmbedAuthor***REMOVED******REMOVED***,
		Color:       0x1FD36F,
		Title:       "User Permission information",
		Description: "Report for " + mention + "'s permissions under skilbot.",
		Thumbnail: &dsg.MessageEmbedThumbnail***REMOVED***
			URL:    icon,
			Width:  64,
			Height: 64,
		***REMOVED***,
		Fields: []*dsg.MessageEmbedField***REMOVED***
			&dsg.MessageEmbedField***REMOVED***
				Name:   "Alpha?",
				Value:  strconv.FormatBool(a),
				Inline: true,
			***REMOVED***,
			&dsg.MessageEmbedField***REMOVED***
				Name:   "Bravo Permissions:",
				Value:  strconv.FormatBool(b),
				Inline: true,
			***REMOVED***,
			&dsg.MessageEmbedField***REMOVED***
				Name:   "Charlie Permissions:",
				Value:  strconv.FormatBool(c),
				Inline: true,
			***REMOVED***,
		***REMOVED***,
	***REMOVED***
***REMOVED***
