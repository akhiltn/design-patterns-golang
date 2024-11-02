package behavioral

type Command interface {
  Execute()
}

type Light struct {
  state bool
}

type LightOnCommand struct {
  light *Light
}

func (c *LightOnCommand) Execute() {
  c.light.state = true
}

type LightOffCommand struct {
  light *Light
}

func (c *LightOffCommand) Execute() {
  c.light.state = false
}

type RemoteControl struct {
  command Command
}

func (r *RemoteControl) setCommand(c Command) {
  r.command = c
}

func (r *RemoteControl) Press() {
  r.command.Execute()
}
