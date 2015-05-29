package gang

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	MODE_LS   = "ls"
	MODE_PECO = "peco"
)

type Config struct {
	ListMode string      `json:"listmode"`
	Commands CommandList `json:"commands"`
}

func (c *Config) Save() {
	setting, _ := json.Marshal(c)
	confFile := os.Getenv("HOME") + "/.gang"

	ioutil.WriteFile(confFile, setting, 0644)
}
