package commands

import (
	dsg "github.com/bwmarrin/discordgo"
	//	"strconv"
	// muchanem: only used within the "flags variabe (line 51)" and the commented help variable
	//"time"
)

func ping(s *dsg.Session, m *dsg.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Pong!")
}
