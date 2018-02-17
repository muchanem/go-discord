package ping

import (
	dsg "github.com/bwmarrin/discordgo"
	f "github.com/skilstak/go-discord/lib"
)

var Commands = make(map[string]*f.Command)

func init() {
	Commands["ping"] = &f.Command{
		Name:   "Ping",
		Help:   "Pings the system to see if its online.",
		Action: ping,
	}
}

func ping(session *dsg.Session, message *dsg.MessageCreate) {
	session.ChannelMessageSend(message.ChannelID, "Pong!")
}
