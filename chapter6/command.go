package main

import "fmt"

type File struct {
	name string
}

func NewFile(name string) *File {
	return &File{name: name}
}

func (f *File) open() {
	fmt.Printf("%sが開かれました。\n", f.name)
}

func (f *File) compress() {
	fmt.Printf("%sが圧縮されました。\n", f.name)
}

func (f *File) close() {
	fmt.Printf("%sが閉じられました。\n", f.name)
}

type Command interface {
	execute()
}

type OpenCommand struct {
	file *File
}

func NewOpenCommand(file *File) *OpenCommand {
	return &OpenCommand{file: file}
}

func (oc *OpenCommand) execute() {
	oc.file.open()
}

type CompressCommand struct {
	file *File
}

func NewCompressCommand(file *File) *CompressCommand {
	return &CompressCommand{file: file}
}

func (cc *CompressCommand) execute() {
	cc.file.compress()
}

type CloseCommand struct {
	file *File
}

func NewCloseCommand(file *File) *CloseCommand {
	return &CloseCommand{file: file}
}

func (cc *CloseCommand) execute() {
	cc.file.close()
}

type Queue struct {
	commands []Command
}

func NewQueue() *Queue {
	return &Queue{
		commands: make([]Command, 0),
	}
}

func (q *Queue) addCommand(command Command) {
	q.commands = append(q.commands, command)
}

func (q *Queue) executeCommand() {
	for _, command := range q.commands {
		command.execute()
	}
}

// func main() {
// 	file := NewFile("command.ts")
// 	queue := NewQueue()

// 	queue.addCommand(NewOpenCommand(file))
// 	queue.addCommand(NewCompressCommand(file))
// 	queue.addCommand(NewCloseCommand(file))

// 	queue.executeCommand()
// }
