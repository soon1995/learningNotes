package commands

type Command interface {
	Execute()
	Undo()
}
