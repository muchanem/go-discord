package flagParser

type FlagType string

const (
	Dash       FlagType = "-"
	DoubleDash FlagType = "--"
)

/*
* Flag struct is used to store a single flags data.
*
* Fields:
*  - fType: Type of the flag. "-" or "--"
*  - flag: The name of the flag. Ex: --name gabe miller --> flag = name
*  - flagValues: The name of the values after the flag. Ex: --name gabe miller --> flagValues = gabe miller
 */
type Flag struct ***REMOVED***
	fType      FlagType
	flag       string
	flagValues string
***REMOVED***

/*
* ParseFlags parses a message for flags.
*
* Parameters:
* - args ([]string) | A message split into []string
*
* Returns:
* - ([]*Flag) | A slice of each flag type
 */
func ParseFlags(args []string) []*Flag ***REMOVED***
	curFlag := new(Flag)
	curFlagValues := ""
	output := make([]*Flag, 0)
	for _, arg := range args ***REMOVED***
		if arg[:2] == "--" || arg[0] == '-' ***REMOVED***
			curFlag.flagValues = curFlagValues
			output = append(output, curFlag)
			curFlag = new(Flag)
			curFlagValues = ""
			switch ***REMOVED***
			case arg[:2] == "--":
				curFlag.fType = DoubleDash
				curFlag.flag = arg[2:]
			case arg[0] == '-':
				curFlag.fType = Dash
				curFlag.flag = arg[1:]
			***REMOVED***
		***REMOVED*** else ***REMOVED***
			curFlagValues += " " + arg
		***REMOVED***
	***REMOVED***

	curFlag.flagValues = curFlagValues
	output = append(output, curFlag)
	return output[1:]
***REMOVED***
