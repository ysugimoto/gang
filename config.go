package gang

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Commands CommandList `json:"commands"`
}

func (c *Config) AddCommand(name, cmd string) {
	c.Commands = append(c.Commands, Command{
		Name:  name,
		Cmd:   cmd,
		Times: 0,
	})
}

func (c *Config) IncrementCount(name string) {
	for index, cmd := range c.Commands {
		if cmd.Name != name {
			continue
		}
		c.Commands[index] = Command{
			Name:  cmd.Name,
			Cmd:   cmd.Cmd,
			Times: cmd.Times + 1,
		}
		break
	}
}

func (c *Config) Save() {
	setting, _ := json.Marshal(c)
	confFile := os.Getenv("HOME") + "/.gang"

	ioutil.WriteFile(confFile, setting, 0644)
}

type CommandList []Command

func (cl CommandList) Len() int {
	return len(cl)
}

func (cl CommandList) Swap(i, j int) {
	cl[i], cl[j] = cl[j], cl[i]
}

func (cl CommandList) Less(i, j int) bool {
	return cl[i].Times < cl[j].Times
}

type Command struct {
	Name  string `json:"name"`
	Cmd   string `json:"cmd"`
	Times int    `json:"times"`
}

func (c Command) String() string {
	return c.Name + ": " + c.Cmd
}
