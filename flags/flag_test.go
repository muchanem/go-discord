package flags

import (
	"strings"
	"testing"
)

func TestNoCommand(t *testing.T) ***REMOVED***
	argstr := `--name gabe miller -foo bar`
	args := strings.Split(argstr, " ")
	myflags := Parse(args)
	for _, f := range myflags ***REMOVED***
		t.Log(f)
	***REMOVED***
***REMOVED***
