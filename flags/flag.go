package flags

type Type string

const (
	Dash       Type = "-"
	DoubleDash Type = "--"
)

// A Flag is used to store a single flags data.
//
// Fields:
//  - Type: "-" or "--"
//  - Name: Name of the flag.
//      Ex: --name gabe miller --> Name = name
//  - Values: Single string of values after flag.
//      Ex: --name gabe miller --> Values = gabe miller
//
type Flag struct ***REMOVED***
	Type  Type
	Name  string
	Value []string
***REMOVED***

// Parse parses a message for flags.
//
// Parameters:
// - args ([]string) | A message split into []string
//
// Returns:
// - ([]*Flag) | A slice of each flag type
//
func Parse(args []string) []*Flag ***REMOVED***
	flags := []*Flag***REMOVED******REMOVED***
	var cur *Flag
	for _, arg := range args ***REMOVED***
		switch ***REMOVED***
		case len(arg) > 1 && arg[:2] == "--":
			cur = &Flag***REMOVED***
				Type: DoubleDash,
				Name: arg[2:],
			***REMOVED***
			flags = append(flags, cur)
		case arg[0] == '-':
			cur = &Flag***REMOVED***
				Type: Dash,
				Name: arg[1:],
			***REMOVED***
			flags = append(flags, cur)
		default:
			cur.Value = append(cur.Value, arg)
		***REMOVED***
	***REMOVED***
	return flags
***REMOVED***
