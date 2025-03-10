package main

import (
	"fmt"
	"time"
)

type Memento struct {
	date string
	memo string
}

func newMemento(memo string) *Memento {
	return &Memento{
		date: time.Now().Format("2006-01-02 15:04:05"),
		memo: memo,
	}
}

func (m *Memento) GetMemo() string {
	return m.memo
}

func (m *Memento) GetInfo() string {
	return fmt.Sprintf("%s / (%s)", m.date, m.memo)
}

type Notepad struct {
	memo string
}

func NewNotepad(memo string) *Notepad {
	return &Notepad{memo: memo}
}

func (n *Notepad) AddMemo(memo string) {
	n.memo = memo
}

func (n *Notepad) GetMemo() string {
	return n.memo
}

func (n *Notepad) Save() *Memento {
	fmt.Println("メモを保存しました")
	return newMemento(n.GetMemo())
}

func (n *Notepad) Restore(m *Memento) {
	n.AddMemo(m.GetMemo())
}

type Caretaker struct {
	notepad  *Notepad
	mementos []*Memento
}

func NewCaretaker(notepad *Notepad) *Caretaker {
	return &Caretaker{
		notepad:  notepad,
		mementos: make([]*Memento, 0),
	}
}

func (c *Caretaker) Backup() {
	c.mementos = append(c.mementos, c.notepad.Save())
}

func (c *Caretaker) Undo() {
	if len(c.mementos) == 0 {
		return
	}
	idx := len(c.mementos) - 1
	memento := c.mementos[idx]
	c.mementos = c.mementos[:idx]
	c.notepad.Restore(memento)
}

func (c *Caretaker) ShowHistory() {
	for _, m := range c.mementos {
		fmt.Println(m.GetInfo())
	}
}

// func main() {
// 	notepad := NewNotepad("first memo")
// 	caretaker := NewCaretaker(notepad)

// 	caretaker.Backup()

// 	notepad.AddMemo("second memo")
// 	caretaker.Backup()
// 	notepad.AddMemo("third memo")
// 	caretaker.Backup()

// 	// 現在の状態を表示
// 	fmt.Println(notepad.GetMemo())
// 	caretaker.ShowHistory()

// 	fmt.Println("")
// 	// Undo 操作で状態を復元
// 	caretaker.Undo()
// 	fmt.Println(notepad.GetMemo())
// 	caretaker.Undo()
// 	fmt.Println(notepad.GetMemo())
// 	caretaker.Undo()
// 	fmt.Println(notepad.GetMemo())
// 	// さらに Undo を試みる（履歴が空の場合は状態は変わらない）
// 	caretaker.Undo()
// 	fmt.Println(notepad.GetMemo())

// 	caretaker.ShowHistory()
// }
