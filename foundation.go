package foundation

import (
	dsg "github.com/bwmarrin/discordgo"
)

/* # Big Bot Boye
* This struct is all about who the bot is: its prefrences, authentication, etc.
*
* The fields are self explanitory so I'm not going to detail them.
*
* TODO: I think the auth section needs to be either a different json file or
*       passed in as parameters because the rest isn't nearly as important.
*
 */
type BotType struct {
	Auth struct {
		Username string `json:"username"`
		ClientID string `json:"clientID"`
		Secret   string `json:"secret"`
		Token    string `json:"token"`
	} `json:"auth"`
	Prefs struct {
		Prefix  string `json:"prefix"`
		Playing string `json:"playing"`
		Version string `json:"version"`
	} `json:"prefs"`
	Perms struct {
		WhitelistChannels   bool     `json:"whitelistChannels"`
		WhitelistedChannels []string `json:"whitelistedChannels"`
		BlacklistedChannels []string `json:"blacklistedChannels"`
	} `json:"perms"`
	Users struct {
		RoleWhitelist    bool     `json:"roleWhitelist"`
		ReportPermFails  bool     `json:"reportPermissionFailures"`
		BlacklistedRoles []string `json:"blacklistedRoles"`
		AdminRoles       []string `json:"adminRoles"`
		StaffRoles       []string `json:"staffRoles"`
		StandardRoles    []string `json:"standardUser"`
	} `json:"users"`
}

/* Defines the actual action the bot takes
* This is a key component of the Command struct, it is the actual *thing* that
* the command does when it is run. This can be as simple as printing a string
* or embed to discord or even be a wrapper for a massive series of functions
* for your discord-based RPG.
*
* Parameters:
* - session (*discordgo.Session) The bot session, in case you need to pull data
*	from about discord itself to complete your task
* - session (*discordgo.Message) The entire message that triggered the command
*	this includes the prefix and command itself. You will have to parse out
*	flags and clean the input.
*
* NOTE: THIS DOES NOT RETURN ERRORS. YOU MUST HANDLE ERRORS.
 */
type Action func(session *dsg.Session, message *dsg.MessageCreate)

/* Defines static data about commands the bot runs.
* This is a very large structure that defines all the needed bits for a bot
* command. All bot modules MUST have one of these along with a few outher key
* components so that the bot works.
 */
type Command struct {
	Name   string `json:"name"`
	Help   string `json:"help"`
	Action Action `json:"-"`
}

/* # Get the guild a message was sent in.
* What a pain in the arse.
*
* Parameters:
* - s (type *discordgo.Session) | The current running discord session,
*     (discordgo needs that always apparently)
* - message (type *discordgo.Message) | the author's id is extracted from this.
*
* Returns:
* - st (type *discordgo.Guild) | The guild the message was found in
* - err (type error)           | If an error was encountered during the process
*	This error is an SEP (someone else's problem).
 */
func GetGuild(s *dsg.Session, m *dsg.Message) (st *dsg.Guild, err error) {
	chn, err := s.Channel(m.ChannelID)
	if err != nil {
		return &dsg.Guild{}, err
	}

	gid := chn.GuildID

	return s.Guild(gid)

}

/* # Checks if user can run a command
* This is a more detailed check to see if a user has the role perms to run a
* command. It is a very complicated wrapper for what would in theory be a
* very simple task.
*
* Do not confuse this with commandHandler's "canTriggerBot()". This function
* determines if a user has a specific role, that just checks if the command is
* run in the correct channel and if they aren't blacklisted.
*
* Parameters:
* - s (type *discordgo.Session) The current running discord session,
*     (discordgo needs that always apparently)
* - message (type *discordgo.Message) | the author's id is extracted from this.
* - role (type string) | the role **ID NUMBER!!** that needs to be matched.
*
* Returns:
* - bool | if the role was found or not, if an error was found it will be false
* - err  | if an error was encountered during the process.
*     To be handled by not-the-function
*
* User's whose roles are considered "admin" by the json config file return an
* automatic true regardless if they have the role listed. Hence if you want to
* lock off a command to only "admin" users, provide an empty string
 */
func HasRole(s *dsg.Session, m *dsg.Message, role string) (bool, error) {
	guild, err := GetGuild(s, m)
	if err != nil {
		print(err)
		return false, err
	}
	member, err := s.GuildMember(guild.ID, m.Author.ID)
	if err != nil {
		print(err)
		return false, err
	}
	for _, b := range MyBot.Users.AdminRoles {
		if Contains(member.Roles, b) {
			return true, nil
		}
	}

	hasRole := Contains(member.Roles, role)
	return hasRole, nil
}

/* # Check if item is in array
* This function checks if a value is in a slice (string only)
*
* Parameters:
* - list ([]string) | the slice to be checking against
* - item (string)   | the item looked for in the slice
*
* Returns:
* - bool | If the item was found or not
*
* NOTE: If another Contains() funciton is needed for a different type, rename
* this function to ContainsSliceString() and the other function to
* ContainsSlice<T>() where <T> is the generic type.
 */
func Contains(list []string, item string) bool {
	for _, b := range list {
		if b == item {
			return true
		}
	}
	return false
}

func Panic(s *dsg.Session, m *dsg.MessageCreate, err error, fatal bool) {
	s.ChannelMessageSend(m.ChannelID, "**ERROR ENCOUNTERED. DETAILS FOLLOW:**\n```"+err.Error()+"```\nThis incident will be reported.")
	if fatal {
		s.ChannelMessageSend(m.ChannelID, "The bot is now \"gracefully\" force quitting, however it will fail to close out of its session with discord and may still apear online.\n*Have a good day!*")
		panic(-1)
	}
}

// An initialized instance of the BotType for use everywhere in this project.
var MyBot BotType

// A discordgo session global variable as .Session is needed for a lot.
// This is written to in the main.go file.
var DG *dsg.Session
