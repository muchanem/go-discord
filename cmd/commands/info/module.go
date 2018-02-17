package info

import (
	"errors"
	dsg "github.com/bwmarrin/discordgo"
	f "github.com/skilstak/go-discord"
	"github.com/skilstak/go-discord/dat"
	"github.com/skilstak/go-discord/flags"
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
	dat.Log.Println(errors.New("Received f1. As follows:\"" + f1 + "\"."))
	f2 := strings.SplitAfterN(f1, f.MyBot.Prefs.Prefix+"info", 2)
	dat.Log.Println(errors.New("Received f2. As follows:\"" + f2[1] + "\"."))
	f3 := strings.Split(f2[1], " ")
	dat.Log.Println(errors.New("Received f3. As follows:\"" + f3 + "\"."))

	f := flags.Parse(f3)
	for _, myflags := range f {
		if myflags.Type == flags.Dash && myflags.Name == "e" {
			cfg.embed = true
		} else if myflags.Type == flags.Dash && myflags.Name == "t" {
			cfg.embed = false
		} else if myflags.Type == flags.Dash && myflags.Name == "m" {
			cfg.useDMs = true
		} else if myflags.Type == flags.Dash && myflags.Name == "c" {
			cfg.useDMs = false
		}
	}

	if cfg.useDMs {
		if cfg.embed {
			session.ChannelMessageSendEmbed(message.Author.ID, getBotInfoAsEmbed())
		} else {
			session.ChannelMessageSend(message.Author.ID, getBotInfoAsText())
		}
	} else {
		if cfg.embed {
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
	return "```" + `Bot information:
	A list of commands can be brought up with ` + "`" + f.MyBot.Prefs.Prefix + "help`" + `

	Bot github link: https://github.com/skilstak/discord-public
	Bot version    : ` + f.MyBot.Prefs.Version + `
	Command version: ` + cfg.version + `
	` + "```"
}
