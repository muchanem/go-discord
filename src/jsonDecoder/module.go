package nil

import (
	f "../foundation"
	"encoding/json"
	"io/ioutil"
)

func GetBotInfo(path string) (f.BotType, error) ***REMOVED***
	raw, err0 := ioutil.ReadFile(path)
	var b f.BotType

	if err0 != nil ***REMOVED***
		return b, err0
	***REMOVED***

	err1 := json.Unmarshal(raw, &b)

	if err1 != nil ***REMOVED***
		return b, err1
	***REMOVED***

	return b, nil
***REMOVED***
