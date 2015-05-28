package gang

import (
	"bytes"
	"fmt"
	"github.com/codeskyblue/go-sh"
	"regexp"
	"strings"
)

const (
	SQ = 39
	DQ = 34
)

var BIND = regexp.MustCompile("(\\{:([0-9a-zA-Z\\.@\\-_]+?)\\})")

type Shell struct {
	command []string
}

func NewShell(cmd string) *Shell {
	cmd = parseCommand(cmd)
	return &Shell{
		command: splitString(cmd, "|"),
	}
}

func (s *Shell) Run() ([]byte, error) {
	sess := sh.NewSession()

	for _, cmd := range s.command {
		c := splitString(cmd, " ")
		name := c[0]
		sess.Command(name, c[1:])
	}

	return sess.Output()
}

func parseCommand(cmd string) string {
	var input string
	for {
		matches := BIND.FindStringSubmatch(cmd)
		if len(matches) == 0 {
			break
		}
		fmt.Print("Bind Parameter \"" + matches[2] + "\" is: ")
		fmt.Scanf("%s", &input)
		cmd = strings.Replace(cmd, matches[1], input, -1)
	}

	return cmd
}

func splitString(str, sep string) []string {
	stack := []byte{}
	parsed := []string{}
	input := []byte(str)
	quote := false
	squote := false
	dquote := false

	for _, v := range input {
		if string(v) == sep && !quote {
			parsed = append(parsed, string(bytes.TrimSpace(stack)))
			stack = []byte{}
			continue
		}
		if v == SQ {
			if squote {
				quote = false
				squote = false
			} else {
				quote = true
				squote = true
			}
		} else if v == DQ {
			if dquote {
				quote = false
				dquote = false
			} else {
				quote = true
				dquote = true
			}
		}
		stack = append(stack, v)
	}

	if len(stack) > 0 {
		parsed = append(parsed, string(bytes.TrimSpace(stack)))
	}

	return parsed
}
