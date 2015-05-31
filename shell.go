package gang

import (
	"github.com/codeskyblue/go-sh"
	"github.com/ysugimoto/splitter"
	"regexp"
	"strings"
)

const (
	SQ = 39
	DQ = 34
)

var BIND = regexp.MustCompile("(\\{:([0-9a-zA-Z\\.@\\-_]+?)\\})")
var DOUBLE_SPACE = regexp.MustCompile("\\s{2,}")

type Shell struct {
	command []string
}

func NewShell(cmd string) *Shell {
	cmd = parseCommand(cmd)
	return &Shell{
		command: splitter.SplitString(cmd, "|"),
	}
}

func (s *Shell) Run() ([]byte, error) {
	sess := sh.NewSession()

	for _, cmd := range s.command {
		cmd = DOUBLE_SPACE.ReplaceAllString(strings.TrimSpace(cmd), " ")
		c := splitter.SplitString(cmd, " ")
		name := c[0]
		sess.Command(name, c[1:])
	}

	return sess.Output()
}

func parseCommand(cmd string) string {
	var input string
	var shown bool

	for {
		matches := BIND.FindStringSubmatch(cmd)
		if len(matches) == 0 {
			break
		}
		if !shown {
			Printf("Execute command needs parameter: %s\n", cmd)
			shown = true
		}

		QPrint("Bind Parameter \"" + matches[2] + "\" is: ")
		Scanf("%s", &input)
		cmd = strings.Replace(cmd, matches[1], input, -1)
	}

	return cmd
}
