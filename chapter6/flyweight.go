package main

import "fmt"

type IStamp interface {
	print()
}

type Stamp struct {
	char string
}

func (s *Stamp) print() {
	fmt.Println(s.char)
}

type StampFactory struct {
	pool map[string]*Stamp
}

func NewStampFactory() *StampFactory {
	return &StampFactory{
		pool: make(map[string]*Stamp),
	}
}

func (sf *StampFactory) getStamp(char string) *Stamp {
	if stamp, ok := sf.pool[char]; ok {
		return stamp
	}

	newStamp := &Stamp{char: char}
	sf.pool[char] = newStamp
	return newStamp
}

func (sf *StampFactory) GetPool() map[string]*Stamp {
	return sf.pool
}

func main() {
	factory := NewStampFactory()
	stamp1 := factory.getStamp("し")
	stamp2 := factory.getStamp("ん")
	stamp3 := factory.getStamp("ぶ")
	stamp4 := factory.getStamp("ん")
	stamp5 := factory.getStamp("し")

	stamp1.print()
	stamp2.print()
	stamp3.print()
	stamp4.print()
	stamp5.print()

	fmt.Println(factory.GetPool())
}
