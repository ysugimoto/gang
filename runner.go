package gang

import (
	"encoding/json"
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
	op := NewOperation(conf)

	if help, _ := r.args.GetOptionAsBool("help"); help {
		ShowHelp()
	}

	if r.args.GetCommandSize() == 0 {
		list := r.getCommandList(conf.Commands, false)
		return op.RunList(list)
	}

	switch cmd, _ := r.args.GetCommandAt(1); cmd {
	case "mode":
		mode, ok := r.args.GetCommandAt(2)
		if ok {
			conf.ListMode = mode
		}
		return op.RunListMode(mode)

	case "reload":
		//r.getLatestCommand()
		//name := strings.Split(command, " ")
		//r.commands.Add(name[0], command)
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
