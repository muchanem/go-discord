package foundation

import (
	dsg "github.com/bwmarrin/discordgo"
)

/* # Big Bot Boye
* This struct is all about the bot, its prefrences, authentication, so on.
*
* The fields are self explanitory so I'm not going to detail them.
*
* TODO: I think the auth section needs to be either a different json file or
*       passed in as parameters because the rest isn't nearly as important.
*
 */
type BotType struct ***REMOVED***
	Auth struct ***REMOVED***
		Username string `json:"username"`
		ClientID string `json:"clientID"`
		Secret   string `json:"secret"`
		Token    string `json:"token"`
	***REMOVED*** `json:"auth"`
	Prefs struct ***REMOVED***
		Prefix  string `json:"prefix"`
		Playing string `json:"playing"`
		Version string `json:"version"`
	***REMOVED*** `json:"prefs"`
	Perms struct ***REMOVED***
		WhitelistChannels   bool     `json:"whitelistChannels"`
		WhitelistedChannels []string `json:"whitelistedChannels"`
		BlacklistedChannels []string `json:"blacklistedChannels"`
	***REMOVED*** `json:"perms"`
	Users struct ***REMOVED***
		RoleWhitelist    bool     `json:"roleWhitelist"`
		ReportPermFails  bool     `json:"reportPermissionFailures"`
		BlacklistedRoles []string `json:"blacklistedRoles"`
		AdminRoles       []string `json:"adminRoles"`
		StaffRoles       []string `json:"staffRoles"`
		StandardRoles    []string `json:"standardUser"`
	***REMOVED*** `json:"users"`
***REMOVED***

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
func GetGuild(s *dsg.Session, m *dsg.Message) (st *dsg.Guild, err error) ***REMOVED***
	chn, err := s.Channel(m.ChannelID)
	if err != nil ***REMOVED***
		return &dsg.Guild***REMOVED******REMOVED***, err
	***REMOVED***

	gid := chn.GuildID

	return s.Guild(gid)

***REMOVED***

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
func HasRole(s *dsg.Session, m *dsg.Message, role string) (bool, error) ***REMOVED***
	guild, err := GetGuild(s, m)
	if err != nil ***REMOVED***
		return false, err
	***REMOVED***
	member, err := s.GuildMember(guild.ID, m.Author.ID)
	if err != nil ***REMOVED***
		return false, err
	***REMOVED***
	for _, b := range MyBot.Users.AdminRoles ***REMOVED***
		if Contains(member.Roles, b) ***REMOVED***
			return true, nil
		***REMOVED***
	***REMOVED***

	hasRole := Contains(member.Roles, role)
	return hasRole, nil
***REMOVED***

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
func Contains(list []string, item string) bool ***REMOVED***
	for _, b := range list ***REMOVED***
		if b == item ***REMOVED***
			return true
		***REMOVED***
	***REMOVED***
	return false
***REMOVED***

// An initialized instance of the BotType for use everywhere in this project.
var MyBot BotType

// A discordgo session global variable as .Session is needed for a lot.
// This is written to in the main.go file.
var DG *dsg.Session
