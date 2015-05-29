package gang

import (
	"fmt"
	"unicode/utf8"
)

type Command struct {
	Name  string `json:"name"`
	Cmd   string `json:"cmd"`
	Times int    `json:"times"`
}

func (c Command) String(size int) string {
	format := fmt.Sprintf("%%-%ds", size+1)
	return fmt.Sprintf(format, c.Name) + ": " + c.Cmd
}

func (c *Command) Increment() {
	c.Times++
}

type CommandList []Command

func (cl CommandList) GetMaxNameSize() (max int) {
	for _, c := range cl {
		size := utf8.RuneCountInString(c.Name)
		if max < size {
			max = size
		}
	}

	return max
}

func (cl CommandList) Len() int {
	return len(cl)
}

func (cl CommandList) Swap(i, j int) {
	cl[i], cl[j] = cl[j], cl[i]
}

func (cl CommandList) Less(i, j int) bool {
	return cl[i].Times > cl[j].Times
}

func (cl CommandList) Find(name string) *Command {
	for _, c := range cl {
		if c.Name == name {
			return &c
		}
	}

	return nil
}

func (cl CommandList) Kill(name string) (killed bool) {
	for i, c := range cl {
		if c.Name == name {
			tmp := cl[:i]
			tmp = append(tmp, cl[i:]...)
			cl = tmp
			killed = true
		}
	}

	return
}

func (cl CommandList) FindIndex(index int) *Command {
	return &cl[index]
}

func (cl CommandList) Add(name, cmd string) {
	cl = append(cl, Command{
		Name: name,
		Cmd:  cmd,
	})
}
