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
type Flag struct {
	fType      FlagType
	flag       string
	flagValues string
}

/*
* ParseFlags parses a message for flags.
*
* Parameters:
* - args ([]string) | A message split into []string
*
* Returns:
* - ([]*Flag) | A slice of each flag type
 */
func ParseFlags(args []string) []*Flag {
	curFlag := new(Flag)
	curFlagValues := ""
	output := make([]*Flag, 0)
	for _, arg := range args {
		if arg[:2] == "--" || arg[0] == '-' {
			curFlag.flagValues = curFlagValues
			output = append(output, curFlag)
			curFlag = new(Flag)
			curFlagValues = ""
			switch {
			case arg[:2] == "--":
				curFlag.fType = DoubleDash
				curFlag.flag = arg[2:]
			case arg[0] == '-':
				curFlag.fType = Dash
				curFlag.flag = arg[1:]
			}
		} else {
			curFlagValues += " " + arg
		}
	}

	curFlag.flagValues = curFlagValues
	output = append(output, curFlag)
	return output[1:]
}
