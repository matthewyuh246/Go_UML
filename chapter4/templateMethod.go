package main

import "fmt"

type Testable interface {
	Setup()
	Execute()
	Teardown()
}

type BaseTest struct{}

func (b *BaseTest) Teardown() {
	fmt.Println("teardown")
}

// func Test(t Testable) {// 	t.Setup()
// 	t.Execute()
// 	t.Teardown()
// }

type ItemServiceTest struct {
	BaseTest
}

func (i *ItemServiceTest) Setup() {
	fmt.Println("setup: ItemServiceTest")
}

func (i *ItemServiceTest) Execute() {
	fmt.Println("execute: ItemServiceTest")
}

type UserServiceTest struct {
	BaseTest
}

func (u *UserServiceTest) Setup() {
	fmt.Println("setup: UserServiceTest")
}

func (u *UserServiceTest) Execute() {
	fmt.Println("execute: UserServiceTest")
}

// func main() {
// 	itemServiceTest := &ItemServiceTest{}
// 	userServiceTest := &UserServiceTest{}

// 	Test(itemServiceTest)
// 	fmt.Println("")
// 	Test(userServiceTest)
// }
