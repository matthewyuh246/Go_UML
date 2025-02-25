package main

import (
	"fmt"
	"regexp"
)

type ValidationHandler interface {
	SetNext(handler ValidationHandler)
	Validate(input string) bool
}

type BaseValidationHandler struct {
	nextHandler ValidationHandler
}

func (b *BaseValidationHandler) SetNext(handler ValidationHandler) {
	b.nextHandler = handler
}

type NotNullValidationHandler struct {
	BaseValidationHandler
}

func (n *NotNullValidationHandler) execValidation(input string) bool {
	result := input != ""
	fmt.Printf("NotNullValidationの結果: %v\n", result)
	return result
}

func (n *NotNullValidationHandler) getErrorMessage() {
	fmt.Println("何も入力されていません")
}

func (n *NotNullValidationHandler) Validate(input string) bool {
	result := n.execValidation(input)
	if !result {
		n.getErrorMessage()
		return false
	} else if n.nextHandler != nil {
		return n.nextHandler.Validate(input)
	} else {
		return true
	}
}

type AlphabetValidationHandler struct {
	BaseValidationHandler
}

func (a *AlphabetValidationHandler) execValidation(input string) bool {
	reg := regexp.MustCompile("^[a-zA-Z]+$")
	result := reg.MatchString(input)
	fmt.Printf("AlphabetValidationの結果: %v\n", result)
	return result
}

func (a *AlphabetValidationHandler) getErrorMessage() {
	fmt.Println("半角英字で入力してください")
}

func (a *AlphabetValidationHandler) Validate(input string) bool {
	result := a.execValidation(input)
	if !result {
		a.getErrorMessage()
		return false
	} else if a.nextHandler != nil {
		return a.nextHandler.Validate(input)
	} else {
		return true
	}
}

type MinLengthValidationHandler struct {
	BaseValidationHandler
}

func (m *MinLengthValidationHandler) execValidation(input string) bool {
	result := len(input) >= 8
	fmt.Printf("MinLengthValidationの結果: %v\n", result)
	return result
}

func (m *MinLengthValidationHandler) getErrorMessage() {
	fmt.Println("8文字以上で入力してください")
}

func (m *MinLengthValidationHandler) Validate(input string) bool {
	result := m.execValidation(input)
	if !result {
		m.getErrorMessage()
		return false
	} else if m.nextHandler != nil {
		return m.nextHandler.Validate(input)
	} else {
		return true
	}
}

func main() {
	notNullHandler := &NotNullValidationHandler{}
	alphabetHandler := &AlphabetValidationHandler{}
	minLengthHandler := &MinLengthValidationHandler{}

	notNullHandler.SetNext(alphabetHandler)
	alphabetHandler.SetNext(minLengthHandler)

	result := notNullHandler.Validate("hellowworld")
	if result {
		fmt.Println("すべてのバリデーションに成功")
	}
}
