package ping

import (
	dsg "github.com/bwmarrin/discordgo"
	f "github.com/skilstak/go-discord"
)

var Commands = make(map[string]*f.Command)

func init() {
	Commands["ping"] = &f.Command{
		Name:   "Ping The Bot",
		Help:   "Pings the bot to see if its online.",
		Action: ping,
	}
}

func ping(session *dsg.Session, message *dsg.MessageCreate) {
	session.ChannelMessageSend(message.ChannelID, "Pong!")
}
