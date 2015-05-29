package gang

import (
	"bufio"
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
		buffer = []byte("{\"listmode\":\"ls\",\"commands\":[]}")
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

		r.runList(list)

		return 0
	}

	switch cmd, _ := r.args.GetCommandAt(1); cmd {
	case "mode":
		mode, ok := r.args.GetCommandAt(1)
		if ok {
			conf.ListMode = mode
		}
		return 0

	case "reload":
		//r.getLatestCommand()
		//name := strings.Split(command, " ")
		//r.commands.Add(name[0], command)
		return 0

	case "ammo":
		list := r.getCommandList(true)

		if len(list) == 0 {
			fmt.Println("No Commands Available")
			return 0
		}

		r.runList(list)
		return 0

	case "kill":
		name, ok := r.args.GetCommandAt(2)
		if !ok {
			fmt.Println("[Error] \"kill\" must be supplied command name.")
			return 1
		}

		commands := CommandList{}
		for _, c := range conf.Commands {
			if c.Name != name {
				commands = append(commands, c)
			}
		}

		if len(conf.Commands) == len(commands) {
			fmt.Printf("[Error] \"%s\" command is not exists. Nothig to do.\n", name)
			return 1
		}

		conf.Commands = commands

		fmt.Printf("Command %s killed.\n", name)
		return 0

	case "bullet":
		var (
			name, cmd     string
			ok, overwrite bool
		)

		if name, ok = r.args.GetCommandAt(2); !ok {
			fmt.Println("[Error] \"bullet\" command needs at least 1 parametes: [name].")
			return 1
		}

		if c := conf.Commands.Find(name); c != nil {
			msg := fmt.Sprintf("Command name \"%s\" already exists. Override it? [y/n]: ", name)
			if !GetInputOfYesNo(msg) {
				fmt.Println("aborted.")
				return 1
			}
			overwrite = true
		}

		if cmd, ok = r.args.GetCommandAt(3); !ok {
			fmt.Printf("[bullet:%s] input runnable command you like\n", name)
			fmt.Print("command> ")
			reader := bufio.NewReader(os.Stdin)
			cmd, _ = reader.ReadString('\n')
			cmd = strings.TrimRight(cmd, "\n")
			if cmd == "" {
				fmt.Println("Comamnd must not empty!")
				return 1
			}
		}

		if !overwrite {
			conf.Commands = append(conf.Commands, Command{
				Name: name,
				Cmd:  cmd,
			})
			fmt.Printf("Command \"%s\" Bulleted.\n", name)
		} else {
			for i, c := range conf.Commands {
				if c.Name == name {
					conf.Commands[i] = Command{
						Name: name,
						Cmd:  cmd,
					}
					break
				}
			}
			fmt.Printf("Command \"%s\" Reloaded.\n", name)
		}
		return 0

	default:
		if c := conf.Commands.Find(cmd); c != nil {
			r.shell(c.Cmd)
		} else {
			fmt.Printf("Command %s is not available.\n", cmd)
		}
	}

	return 0
}

func (r *Runner) runList(list []string) {
	switch conf.ListMode {
	case MODE_PECO:
		p := pecolify.New()
		selected, _ := p.Transform(list)
		if selected != "" {
			r.runCommand(selected)
		}
	default:
		fmt.Println("Available Commands =========================")
		for i, c := range list {
			fmt.Printf("[%d] %s\n", i+1, c)
		}
		selected := GetInputOfSelectIndex(list)
		if c := conf.Commands.FindIndex(selected - 1); c != nil {
			c.Increment()
			r.shell(c.Cmd)
		}
	}
}

func (r *Runner) getCommandList(sorted bool) (list []string) {
	if sorted {
		sort.Sort(conf.Commands)
	}

	max := conf.Commands.GetMaxNameSize()
	for _, cmd := range conf.Commands {
		list = append(list, cmd.String(max))
	}

	return
}

func (r *Runner) runCommand(cmd string) {
	name, command := r.ParseCommand(cmd)

	if c := conf.Commands.Find(name); c != nil {
		c.Increment()
	}
	//fmt.Printf("Execute Command: %s\n", command)

	r.shell(command)
}

func (r *Runner) shell(command string) {
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
