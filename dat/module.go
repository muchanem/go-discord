package nil

import (
	"encoding/json"
	f "github.com/skilstak/discord-public/lib"
	"io/ioutil"
)

func GetBotInfo(path string) (f.BotType, error) {
	raw, err0 := ioutil.ReadFile(path)
	var b f.BotType

	if err0 != nil {
		return b, err0
	}

	err1 := json.Unmarshal(raw, &b)

	if err1 != nil {
		return b, err1
	}

	return b, nil
}
