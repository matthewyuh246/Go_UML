package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Context struct {
	expression string
	date       time.Time
}

func NewContext(expression string, date time.Time) (*Context, error) {
	ctx := &Context{
		expression: expression,
		date:       date,
	}
	if err := ctx.validate(); err != nil {
		return nil, err
	}
	return ctx, nil
}

func (c *Context) validate() error {
	if len(c.expression) != 10 {
		return errors.New("expressionが不正です")
	}
	if !strings.Contains(c.expression, "YYYY") || !strings.Contains(c.expression, "MM") || !strings.Contains(c.expression, "DD") {
		return errors.New("expressionが不正です")
	}
	return nil
}

type expression interface {
	interpret(context *Context) *Context
}

type yearExpression struct {
	child expression
}

func (y *yearExpression) setChild(child expression) {
	y.child = child
}

func (y *yearExpression) interpret(context *Context) *Context {
	year := fmt.Sprintf("%d", context.date.Year())
	context.expression = strings.Replace(context.expression, "YYYY", year, 1)
	if y.child != nil {
		return y.child.interpret(context)
	}
	return context
}

type monthExpression struct {
	child expression
}

func (m *monthExpression) setChild(child expression) {
	m.child = child
}

func (m *monthExpression) interpret(context *Context) *Context {
	month := fmt.Sprintf("%d", context.date.Month())
	context.expression = strings.Replace(context.expression, "MM", month, 1)
	if m.child != nil {
		return m.child.interpret(context)
	}
	return context
}

type dayExpression struct {
	child expression
}

func (d *dayExpression) setChild(child expression) {
	d.child = child
}

func (d *dayExpression) interpret(context *Context) *Context {
	day := fmt.Sprintf("%d", context.date.Day())
	context.expression = strings.Replace(context.expression, "DD", day, 1)
	if d.child != nil {
		return d.child.interpret(context)
	}
	return context
}

func main() {
	now := time.Now()
	expression := "MM/DD/YYYY"

	ctx, err := NewContext(expression, now)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	yearExpression := &yearExpression{}
	monthExpression := &monthExpression{}
	dayExpression := &dayExpression{}

	monthExpression.setChild(dayExpression)
	yearExpression.setChild(monthExpression)

	result := yearExpression.interpret(ctx)

	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(result.expression)
}
