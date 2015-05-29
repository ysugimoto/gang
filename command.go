package gang

type Command struct {
	Name  string `json:"name"`
	Cmd   string `json:"cmd"`
	Times int    `json:"times"`
}

func (c Command) String() string {
	return c.Name + ": " + c.Cmd
}

func (c *Command) Increment() {
	c.Times++
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

func (cl CommandList) Find(name string) *Command {
	for _, c := range cl {
		if c.Name == name {
			return &c
		}
	}

	return nil
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
