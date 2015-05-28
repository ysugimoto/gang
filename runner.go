package gang

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ysugimoto/go-cliargs"
	"github.com/ysugimoto/pecolify"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

var conf *Config = &Config{}

func init() {
	var buffer []byte

	// check configuration file
	confFile := os.Getenv("HOME") + "/.gang"
	if _, err := os.Stat(confFile); err != nil {
		// create file
		buffer = []byte("{\"commands\":[]}")
		ioutil.WriteFile(confFile, buffer, 0644)
	} else {
		buffer, _ = ioutil.ReadFile(confFile)
	}
	json.Unmarshal(buffer, conf)
}

type Runner struct {
	args *cliarg.Arguments
}

func NewRunner(args *cliarg.Arguments) *Runner {
	return &Runner{
		args: args,
	}
}

func (r *Runner) Run() {

	ret := r._run()
	if ret == 0 {
		conf.Save()
	}
}

func (r *Runner) _run() int {
	if r.args.GetCommandSize() == 0 {
		list := r.getCommandList(false)

		if len(list) == 0 {
			fmt.Println("No Commands Available")
			return 0
		}

		p := pecolify.New()
		selected, _ := p.Transform(list)
		if selected != "" {
			r.runCommand(selected)
		}
		return 0
	}

	switch cmd, _ := r.args.GetCommandAt(1); cmd {
	case "reload":
		command := r.getLatestCommand()
		name := strings.Split(command, " ")
		conf.AddCommand(name[0], command)
		return 0

	case "ammo":
		list := r.getCommandList(false)

		if len(list) == 0 {
			fmt.Println("No Commands Available")
			return 0
		}

		p := pecolify.New()
		selected, _ := p.Transform(list)
		if selected != "" {
			r.runCommand(selected)
		}
		return 0

	case "kill":
		name, ok := r.args.GetCommandAt(2)
		if !ok {
			fmt.Println("[Error] \"kill\" must be supplied command name.")
			return 1
		}
		commands := []Command{}

		for _, cmd := range conf.Commands {
			if cmd.Name != name {
				commands = append(commands, cmd)
			}
		}

		if len(commands) == len(conf.Commands) {
			fmt.Printf("[Error] \"%s\" command is not exists. Nothig to do.\n", name)
			return 1
		}

		conf.Commands = commands
		fmt.Printf("Command %s killed.\n", name)
		return 0

	case "bullet":
		var (
			name, cmd string
			ok        bool
		)

		if name, ok = r.args.GetCommandAt(2); !ok {
			fmt.Println("[Error] \"bullet\" command needs 2 parametes: [name/command].")
			return 1
		}
		if cmd, ok = r.args.GetCommandAt(3); !ok {
			fmt.Println("[Error] \"bullet\" command needs 2 parametes: [name/command].")
			return 1
		}

		for _, cmd := range conf.Commands {
			if cmd.Name == name {
				fmt.Printf("[Error] command \"%s\" already exists.\n", name)
			}
		}

		conf.Commands = append(conf.Commands, Command{
			Name:  name,
			Cmd:   cmd,
			Times: 0,
		})

		fmt.Printf("Command \"%s\" Bulleted.\n", name)
		return 0

	default:

	}

	return 0
}

func (r *Runner) getCommandList(sorted bool) (list []string) {
	if sorted {
		sort.Sort(conf.Commands)
	}

	for _, cmd := range conf.Commands {
		list = append(list, cmd.String())
	}

	return
}

func (r *Runner) runCommand(cmd string) {
	name, command := r.ParseCommand(cmd)

	conf.IncrementCount(name)
	fmt.Printf("Execute Command: %s\n", command)

	shell := NewShell(command)
	out, _ := shell.Run()

	output := bytes.Trim(out, "\r\n\"")
	if len(output) > 0 {
		fmt.Println(string(bytes.Trim(out, "\r\n\"")))
	}
}

func (r *Runner) ParseCommand(cmd string) (name, command string) {
	spl := strings.Split(cmd, ":")
	name = spl[0]
	command = strings.Join(spl[1:], ":")

	return name, strings.TrimSpace(command)
}

func (r *Runner) getLatestCommand() string {
	return ""
}
