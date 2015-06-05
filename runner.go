package gang

import (
	"encoding/json"
	"fmt"
	"github.com/ysugimoto/go-cliargs"
	"io/ioutil"
	"os"
	"sort"
)

var conf *Config = &Config{}

func init() {
	var buffer []byte

	// check configuration file
	confFile := os.Getenv("HOME") + "/.gang"
	if _, err := os.Stat(confFile); err != nil {
		// create file
		buffer = []byte("{\"listmode\":\"ls\",\"shell\":\"bash\",\"commands\":[]}")
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
	op := NewOperation(conf)

	// switch current directory
	if dir, ok := r.args.GetOptionAsString("directory"); ok {
		if err := os.Chdir(dir); err != nil {
			EPrintf("%v", err)
			os.Exit(1)
		}
	}

	if help, _ := r.args.GetOptionAsBool("help"); help {
		ShowHelp()
	}
	if version, _ := r.args.GetOptionAsBool("version"); version {
		fmt.Println(VERSION)
		os.Exit(1)
	}

	if r.args.GetCommandSize() == 0 {
		list := r.getCommandList(conf.Commands, false)
		return op.RunList(list)
	}

	switch cmd, _ := r.args.GetCommandAt(1); cmd {
	case "dump":
		c, _ := json.MarshalIndent(conf, "  ", "  ")
		fmt.Println(string(c))
		return 0

	case "shell":
		shell, ok := r.args.GetCommandAt(2)
		if !ok {
			EPrintln("shell subcommand needs second parameter ( bash or zsh )")
			return 1
		}
		return op.RunShellMode(shell)

	case "mode":
		mode, ok := r.args.GetCommandAt(2)
		if !ok {
			EPrintln("mode subcommand needs second parameter ( ls or peco )")
			return 1
		}
		return op.RunListMode(mode)

	case "reload":
		/* This implements is experimental
		if last, err := r.getLastCommand(); err != nil {
			EPrintln("Cannot find last command on your environment.")
			return 1
		} else {
			Printf("Last command is %s", string(last))
			return 0
		}
		*/
		return 0

	case "ammo":
		list := r.getCommandList(conf.Commands, true)
		return op.RunList(list)

	case "kill":
		name, ok := r.args.GetCommandAt(2)
		if !ok {
			EPrintln("[Error] \"kill\" must be supplied command name.")
			return 1
		}
		return op.RunKill(name)

	case "bullet":
		var (
			name, cmd     string
			nameOk, cmdOk bool
		)

		name, nameOk = r.args.GetCommandAt(2)
		cmd, cmdOk = r.args.GetCommandAt(3)

		return op.RunBullet(name, nameOk, cmd, cmdOk)

	default:
		return op.RunDefault(cmd)
	}

	return 0
}

func (r *Runner) getCommandList(commands CommandList, sorted bool) (list []string) {
	if sorted {
		sort.Sort(commands)
	}

	max := commands.GetMaxNameSize()
	for _, cmd := range commands {
		list = append(list, cmd.String(max))
	}

	return
}

/*
func (r *Runner) getLastCommand() ([]byte, error) {
	var historyFile string

	switch conf.Shell {
	case SHELL_BASH:
		historyFile = os.Getenv("HOME") + "/.bash_history"
	case SHELL_ZSH:
		historyFile = os.Getenv("HOME") + "/.zsh_history"
	default:
		return nil, errors.New("")
	}

	if _, err := os.Stat(historyFile); err != nil {
		return nil, nil
	}

	return exec.Command("tail", "-n", "1", historyFile).Output()
}
*/
