package gang

import (
	"github.com/codeskyblue/go-sh"
	"github.com/ysugimoto/splitter"
	"os"
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
	command string
}

func NewShell(cmd string) *Shell {
	return &Shell{
		command: parseCommand(cmd),
	}
}

func (s *Shell) Run() ([]byte, error) {
	sess := sh.NewSession()
	command := splitCommand(s.command)

	for _, cmd := range command {
		c := splitter.SplitString(cmd, " ")
		// trick: "cd path" is simply Chdir.
		if len(c) == 2 && c[0] == "cd" {
			if err := os.Chdir(c[1]); err != nil {
				return nil, err
			} else {
				continue
			}
		}

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

	return DOUBLE_SPACE.ReplaceAllString(strings.TrimSpace(cmd), " ")
}

func splitCommand(cmd string) []string {
	command := []string{}
	piped := splitter.SplitString(cmd, "|")

	for _, c := range piped {
		if c == "" {
			continue
		}
		amp := splitter.SplitString(c, "&&")
		for _, a := range amp {
			if a == "" {
				continue
			}
			sep := splitter.SplitString(a, ";")
			for _, s := range sep {
				if s == "" {
					continue
				}
				command = append(command, strings.TrimSpace(s))
			}
		}
	}

	return command
}
