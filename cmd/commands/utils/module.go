package commands

import (
	dsg "github.com/bwmarrin/discordgo"
	f "github.com/skilstak/discord-public/lib"
)

var Commands = make(map[string]*f.Command)

func init() {
	Commands["getroles"] = &f.Command{
		Name:   "Get Server Roles",
		Help:   "Goes through all of the server's roles and posts them and their IDs.",
		Action: getRoles,
	}
}

/* # Get server roles
* A g-d impossibility.
*
* Parameters/return values:
* This function complies with the foundation's action function protocol.
* For documentation on that, please see https://github.com/skilstak/discord-public
*
* TODO: Make a godoc for our nonsence.
*
* NOTE: If you print this into a discord chat, it WILL mention @everyone
 */
func getRoles(session *dsg.Session, message *dsg.Message) {
	s := session
	m := message

	guild, err := f.GetGuild(f.DG, m)
	if err != nil {
		f.Panic(s, m, err, false)
		return
	}
	roles, err := f.DG.GuildRoles(guild.ID)
	if err != nil {
		f.Panic(s, m, err, false)
		return
	}
	role := "Server role list:\n```\n"
	for _, r := range roles {
		role += "Name: " + r.Name + "; ID: " + r.ID + ";\n"
	}
	role += "```"
	s.ChannelMessageSend(m.ChannelID, role)
}
