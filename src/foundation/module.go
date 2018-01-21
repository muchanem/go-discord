package foundation

import (
//dsg "github.com/bwmarrin/discordgo"
)

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
		BlacklistedUsers []string `json:"blacklistedUsers"`
		AlphaUsers       []string `json:"alphaUsers"`
		BravoUsers       []string `json:"bravoUsers"`
		BravoOverride    bool     `json:"bravoOverride"`
		CharlieUsers     []string `json:"charlieUsers"`
	} `json:"users"`
}

func Contains(list []string, item string) bool {
	for _, b := range list {
		if b == item {
			return true
		}
	}
	return false
}

var MyBot BotType
