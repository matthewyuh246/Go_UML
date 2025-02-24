package main

import "fmt"

type Button interface {
	Press()
}

type Checkbox interface {
	Switch()
}

type GUIFactory interface {
	CreateButton() Button
	CreateCheckBox() Checkbox
}

type WindowsButton struct{}

func (w WindowsButton) Press() {
	fmt.Println("Windows用のボタンが押されました")
}

type WindowsCheckbox struct{}

func (w WindowsCheckbox) Switch() {
	fmt.Println("Windowsのチェックボックスが切り替えられました")
}

type WindowsGUIFactory struct{}

func (wg WindowsGUIFactory) CreateButton() Button {
	return WindowsButton{}
}

func (wg WindowsGUIFactory) CreateCheckBox() Checkbox {
	return WindowsCheckbox{}
}

type MacButton struct{}

func (w MacButton) Press() {
	fmt.Println("Mac用のボタンが押されました")
}

type MacCheckbox struct{}

func (w MacCheckbox) Switch() {
	fmt.Println("Macのチェックボックスが切り替えられました")
}

type MacGUIFactory struct{}

func (wg MacGUIFactory) CreateButton() Button {
	return MacButton{}
}

func (wg MacGUIFactory) CreateCheckBox() Checkbox {
	return MacCheckbox{}
}

func run(factory GUIFactory) {
	button := factory.CreateButton()
	checkbox := factory.CreateCheckBox()
	button.Press()
	checkbox.Switch()
}

func main() {
	run(WindowsGUIFactory{})
	run(MacGUIFactory{})
}
