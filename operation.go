package gang

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/ysugimoto/pecolify"
	"os"
	"strings"
)

type Operation struct {
	config *Config
}

func NewOperation(conf *Config) *Operation {
	return &Operation{
		config: conf,
	}
}

func (o *Operation) RunListMode(mode string) int {
	o.config.ListMode = mode
	QPrintf("List mode changed to \"%s\".\n", mode)
	return 0
}

func (o *Operation) RunShellMode(shell string) int {
	o.config.Shell = shell
	QPrintf("Shell changed to \"%s\".\n", shell)
	return 0
}

func (o *Operation) RunShell(command string) int {
	shell := NewShell(command)
	out, _ := shell.Run()

	output := bytes.Trim(out, "\r\n\"")
	if len(output) > 0 {
		fmt.Println(string(bytes.Trim(out, "\r\n\"")))
		return 0
	}
	return 1
}

func (o *Operation) RunCommand(cmd string) int {
	name, command := o.parseCommand(cmd)

	if c := conf.Commands.Find(name); c != nil {
		c.Increment()
	} else {
		EPrintf("Command %s is not available.\n", name)
		return 1
	}

	o.RunShell(command)
	return 0
}

func (o *Operation) RunList(list []string) int {
	if len(list) == 0 {
		EPrintln("No Commands Available")
		return 0
	}

	switch conf.ListMode {
	case MODE_PECO:
		p := pecolify.New()
		selected, _ := p.Transform(list)
		if selected != "" {
			o.RunCommand(selected)
		}
	default:
		fmt.Println(MAGENTA + "Available Commands =========================" + RESET)
		for i, c := range list {
			fmt.Printf("[%d] %s\n", i+1, c)
		}
		selected := GetInputOfSelectIndex(list)
		if selected == 0 {
			Println("aborted.")
			return 1
		}
		if c := o.config.Commands.FindIndex(selected - 1); c != nil {
			c.Increment()
			o.RunShell(c.Cmd)
		}
	}
	return 0
}

func (o *Operation) RunKill(name string) int {
	commands := CommandList{}
	for _, c := range o.config.Commands {
		if c.Name != name {
			commands = append(commands, c)
		}
	}

	if len(conf.Commands) == len(commands) {
		EPrintf("[Error] \"%s\" command is not exists. Nothig to do.\n", name)
		return 1
	}

	o.config.Commands = commands

	QPrintf("Command %s killed.\n", name)
	return 0
}

func (o *Operation) RunBullet(name string, nameOk bool, cmd string, cmdOk bool) int {
	var overwrite bool

	if !nameOk {
		EPrintln("[Error] \"bullet\" command needs at least 1 parametes: [name].")
		return 1
	}

	if c := o.config.Commands.Find(name); c != nil {
		msg := QSprintf("Command name \"%s\" already exists. Override it? [y/n]: ", name)
		if !GetInputOfYesNo(msg) {
			Println("aborted.")
			return 1
		}
		overwrite = true
	}

	if !cmdOk {
		Printf("[bullet:%s] input runnable command you like\n", name)
		Print("command> ")
		reader := bufio.NewReader(os.Stdin)
		cmd, _ = reader.ReadString('\n')
		cmd = strings.TrimRight(cmd, "\n")
		if cmd == "" {
			EPrintln("[Error] Comamnd must not empty!")
			return 1
		}
	}

	if !overwrite {
		o.config.Commands = append(o.config.Commands, Command{
			Name: name,
			Cmd:  cmd,
		})
		QPrintf("Command \"%s\" Bulleted.\n", name)
	} else {
		for i, c := range o.config.Commands {
			if c.Name == name {
				conf.Commands[i] = Command{
					Name: name,
					Cmd:  cmd,
				}
				break
			}
		}
		QPrintf("Command \"%s\" Reloaded.\n", name)
	}
	return 0
}

func (o *Operation) RunDefault(cmd string) int {
	if c := o.config.Commands.Find(cmd); c != nil {
		o.RunShell(c.Cmd)
		return 0
	} else {
		EPrintf("Command %s is not available.\n", cmd)
		return 1
	}
}

func (o *Operation) parseCommand(cmd string) (name, command string) {
	spl := strings.Split(cmd, ":")
	name = spl[0]
	command = strings.Join(spl[1:], ":")

	return strings.TrimSpace(name), strings.TrimSpace(command)
}
