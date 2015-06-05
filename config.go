package gang

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	MODE_LS   = "ls"
	MODE_PECO = "peco"

	SHELL_BASH = "bash"
	SHELL_ZSH  = "zsh"
)

type Config struct {
	ListMode string      `json:"listmode"`
	Shell    string      `json:"shell"`
	Commands CommandList `json:"commands"`
}

func (c *Config) Save() {
	setting, _ := json.Marshal(c)
	confFile := os.Getenv("HOME") + "/.gang"

	ioutil.WriteFile(confFile, setting, 0644)
}
