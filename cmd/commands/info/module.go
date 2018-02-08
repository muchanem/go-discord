package commands

import (
	dsg "github.com/bwmarrin/discordgo"
	f "github.com/skilstak/discord-public/lib"
	//	"strconv"
	// muchanem: only used within the "flags variabe (line 51)" and the commented help variable
	//"time"
)

var Commands = make(map[string]*f.Command)

func init() {
	Commands["info"] = &f.Command{
		Name: "Bot Info",
		Help: `Gets information about the bot, version number, so on.
		Options:
		 -e : Get info as embed (default)
		 -t : Get info as raw text
		 -m : Get info via direct message`,
		Action: Info(),
	}
}

func Info(s *dsg.Session, m *dsg.MessageCreate) {
	s.ChannelMessageSendEmbed(m.ChannelID, getBotInfo())
}
func getBotInfo() *dsg.MessageEmbed {
	return &dsg.MessageEmbed{
		Author:      &dsg.MessageEmbedAuthor{},
		Color:       0x073642,
		Title:       "Bot Information",
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
				Name:   "Link to bot framework:",
				Value:  "https://github.com/skilstak/discord-public",
				Inline: true,
			},
		},
	}
}
