package gang

import (
	"fmt"
	"os"
)

const (
	VERSION = "0.8.0"
)

func ShowHelp() {
	fmt.Println(GREENB+"====================================================", RESET)
	fmt.Println(GREENB+" Gang - A console command snippet management tool", RESET)
	fmt.Println(GREENB+"====================================================", RESET)

	help := `Usage:
    gang [subcommand|snippet-name] [paramters...] [option]

Options:
    -h, --help      : Show this help
    -d, --directory : Change current directory

Subcommands:
    mode [ls|peco]                  : Change list mode
    ammo                            : Show list sorted by call times
    kill [snippet-name]             : Remove snippet
    bullet [snippet-name] [command] : Register the snippet
    [snippet-name] [bind,...]       : Run the snippet command
`
	fmt.Println(help)
	os.Exit(0)
}
