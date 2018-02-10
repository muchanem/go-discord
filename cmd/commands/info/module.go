package commands

import (
	"errors"
	dsg "github.com/bwmarrin/discordgo"
	"github.com/skilstak/discord-public/dat"
	"github.com/skilstak/discord-public/flags"
	f "github.com/skilstak/discord-public/lib"
	"strings"
)

type config struct {
	version string `json:"version"`
	embed   bool   `json:"embedDefault"`
	useDMs  bool   `json:"dmDefault"`
}

var (
	Commands = make(map[string]*f.Command)
	cfg      *config
)

func init() {
	dat.Load("info/config.json", &cfg)
	Commands["info"] = &f.Command{
		Name: "Bot Info",
		Help: `Gets information about the bot, version number, so on.
		Options:
		 -e : Get info as embed
		 -t : Get info as raw text
		 -m : Get info via direct message
		 -c : Post info in channel
		Version : ` + cfg.version + `.
		Github  : https://github.com/skilstak/discord-public/cmd/commands/info`,
		Action: info,
	}
}

func info(session *dsg.Session, message *dsg.MessageCreate) {
	f1 := strings.ToLower(message.Content)
	f2 := strings.SplitAfterN(f1, f.MyBot.Prefs.Prefix+"info", 2)
	dat.Log(errors.New("Recived f2. As follows:"+f2[1]+"."), 0)
	f3 := strings.Split(f2[1], " ")

	e := false
	t := false
	m := false
	c := false

	if len(f3) > 2 {
		f := flags.Parse(f3)
		for _, myflags := range f {
			if myflags.Type == flags.Dash && myflags.Name == "e" {
				e = true
			} else if myflags.Type == flags.Dash && myflags.Name == "t" {
				t = true
			} else if myflags.Type == flags.Dash && myflags.Name == "m" {
				m = true
			} else if myflags.Type == flags.Dash && myflags.Name == "c" {
				c = true
			}
		}
	}
	if m {
		if t {
			session.ChannelMessageSend(message.Author.ID, getBotInfoAsText())
		} else if e {
			session.ChannelMessageSendEmbed(message.Author.ID, getBotInfoAsEmbed())
		} else if cfg.embed {
			session.ChannelMessageSendEmbed(message.Author.ID, getBotInfoAsEmbed())
		} else {
			session.ChannelMessageSend(message.Author.ID, getBotInfoAsText())
		}
	} else if c {
		if t {
			session.ChannelMessageSend(message.ChannelID, getBotInfoAsText())
		} else if e {
			session.ChannelMessageSendEmbed(message.ChannelID, getBotInfoAsEmbed())
		} else if cfg.embed {
			session.ChannelMessageSendEmbed(message.ChannelID, getBotInfoAsEmbed())
		} else {
			session.ChannelMessageSend(message.ChannelID, getBotInfoAsText())
		}
	} else if !cfg.useDMs {
		if t {
			session.ChannelMessageSend(message.ChannelID, getBotInfoAsText())
		} else if e {
			session.ChannelMessageSendEmbed(message.ChannelID, getBotInfoAsEmbed())
		} else if cfg.embed {
			session.ChannelMessageSendEmbed(message.ChannelID, getBotInfoAsEmbed())
		} else {
			session.ChannelMessageSend(message.ChannelID, getBotInfoAsText())
		}
	} else {
		if t {
			session.ChannelMessageSend(message.Author.ID, getBotInfoAsText())
		} else if e {
			session.ChannelMessageSendEmbed(message.Author.ID, getBotInfoAsEmbed())
		} else if cfg.embed {
			session.ChannelMessageSendEmbed(message.Author.ID, getBotInfoAsEmbed())
		} else {
			session.ChannelMessageSend(message.Author.ID, getBotInfoAsText())
		}
	}
}

func getBotInfoAsEmbed() *dsg.MessageEmbed {
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
				Name:   "Bot Version",
				Value:  "Version " + f.MyBot.Prefs.Version + ".",
				Inline: true,
			},
			&dsg.MessageEmbedField{
				Name:   "Link to bot framework:",
				Value:  "https://github.com/skilstak/discord-public",
				Inline: true,
			},
			&dsg.MessageEmbedField{
				Name:   "Command info:",
				Value:  "Version " + cfg.version + ".",
				Inline: true,
			},
		},
	}
}

func getBotInfoAsText() string {
	return ` Bot information:
	A list of commands can be brought up with ` + "`" + f.MyBot.Prefs.Prefix + "help`" + `

	Bot github link: https://github.com/skilstak/discord-public
	Bot version    : ` + f.MyBot.Prefs.Version + `
	Command version: ` + cfg.version + `
	`
}
