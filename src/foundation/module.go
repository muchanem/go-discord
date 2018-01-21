package foundation

import (
//dsg "github.com/bwmarrin/discordgo"
)

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
		BlacklistedUsers []string `json:"blacklistedUsers"`
		AlphaUsers       []string `json:"alphaUsers"`
		BravoUsers       []string `json:"bravoUsers"`
		BravoOverride    bool     `json:"bravoOverride"`
		CharlieUsers     []string `json:"charlieUsers"`
	***REMOVED*** `json:"users"`
***REMOVED***

func Contains(list []string, item string) bool ***REMOVED***
	for _, b := range list ***REMOVED***
		if b == item ***REMOVED***
			return true
		***REMOVED***
	***REMOVED***
	return false
***REMOVED***

var MyBot BotType
