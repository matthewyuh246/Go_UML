package main

import "fmt"

type Mediator interface {
	registerUser(user User)
	sendMessage(msg string, sendUser User)
}

type ChatRoom struct {
	members []User
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		members: []User{},
	}
}

func (c *ChatRoom) registerUser(user User) {
	c.members = append(c.members, user)
}

func (c *ChatRoom) sendMessage(msg string, sendUser User) {
	for _, member := range c.members {
		if member != sendUser {
			member.receive(msg)
		}
	}
}

type User interface {
	send(msg string)
	receive(msg string)
}

type ChatUser struct {
	mediator Mediator
	name     string
}

func NewChatUser(mediator Mediator, name string) *ChatUser {
	return &ChatUser{
		mediator: mediator,
		name:     name,
	}
}

func (cu *ChatUser) send(msg string) {
	fmt.Printf("%s -> メッセージ送信\n", cu.name)
	cu.mediator.sendMessage(msg+" from "+cu.name, cu)
}

func (cu *ChatUser) receive(msg string) {
	fmt.Printf("%s -> メッセージ受信: %s\n", cu.name, msg)
}

// func main() {
// 	chatRoom := NewChatRoom()

// 	yamada := NewChatUser(chatRoom, "yamada")
// 	suzuki := NewChatUser(chatRoom, "suzuki")
// 	tanaka := NewChatUser(chatRoom, "tanaka")

// 	chatRoom.registerUser(yamada)
// 	chatRoom.registerUser(suzuki)
// 	chatRoom.registerUser(tanaka)

// 	yamada.send("こんにちは")
// 	tanaka.send("ヤッホー")
// }
