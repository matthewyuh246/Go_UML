package main

import "fmt"

type UserState interface {
	isAuthenticated() bool
	displayPage()
	nextState() UserState
}

type AuthorizedState struct{}

func (as *AuthorizedState) isAuthenticated() bool {
	return true
}

func (as *AuthorizedState) displayPage() {
	fmt.Println("TOPページ")
}

func (as *AuthorizedState) nextState() UserState {
	return &UnAuthorizedState{}
}

type UnAuthorizedState struct{}

func (us *UnAuthorizedState) isAuthenticated() bool {
	return false
}

func (us *UnAuthorizedState) displayPage() {
	fmt.Println("エラーページ: 認証されていません")
}

func (us *UnAuthorizedState) nextState() UserState {
	return &AuthorizedState{}
}

type User struct {
	state UserState
}

func NewUser() *User {
	return &User{
		state: &UnAuthorizedState{},
	}
}

func (u *User) isAuthenticated() bool {
	return u.state.isAuthenticated()
}

func (u *User) displayPage() {
	u.state.displayPage()
}

func (u *User) switchState() {
	u.state = u.state.nextState()
}

// func main() {
// 	user := NewUser()
// 	fmt.Println(user.isAuthenticated())
// 	user.displayPage()

// 	user.switchState()
// 	fmt.Println(user.isAuthenticated())
// 	user.displayPage()
// }
