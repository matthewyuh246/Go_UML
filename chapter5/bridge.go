package main

import "fmt"

type MessageApp interface {
	send()
}

type LINE struct{}

func (l *LINE) send() {
	fmt.Println("LINEでメッセージ送信")
}

type Twitter struct{}

func (t *Twitter) send() {
	fmt.Println("Twitterでメッセージ送信")
}

type Facebook struct{}

func (f *Facebook) send() {
	fmt.Println("Facebookでメッセージ送信")
}

type OS struct {
	app MessageApp
}

func (o *OS) setApp(app MessageApp) {
	o.app = app
}

type IOS struct {
	OS
}

func (i *IOS) sendMessage() {
	fmt.Println("iOSでメッセージ送信")

	if i.app != nil {
		i.app.send()
	} else {
		panic("アプリが指定されていません")
	}
}

type Android struct {
	OS
}

func (a *Android) sendMessage() {
	fmt.Println("Androidでメッセージ送信")

	if a.app != nil {
		a.app.send()
	} else {
		panic("アプリが指定されていません")
	}
}

func main() {
	line := &LINE{}
	facebook := &Facebook{}

	ios := &IOS{}
	android := &Android{}

	ios.setApp(line)
	ios.sendMessage()
	android.setApp(facebook)
	android.sendMessage()
}
