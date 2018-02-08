package commands

import (
	dsg "github.com/bwmarrin/discordgo"
	f "github.com/skilstak/discord-public/lib"
)

var Command = make(map[string]*f.Command)

func init() {
	Command = &f.Command{
		Name:   "Ping",
		Help:   "Pings the system to see if its online.",
		Action: Ping,
	}
}

func Ping(s *dsg.Session, m *dsg.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Pong!")
}
