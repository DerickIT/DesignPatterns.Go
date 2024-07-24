package main

import (
	"fmt"
)

type Target interface {
	Request() string
}

type Adaptee struct {
}

func (a *Adaptee) SpecificRequest() string {
	return "Adaptee specific request"
}

type Adapter struct {
	adaptee *Adaptee
}

func NewAdapter(adaptee *Adaptee) *Adapter {
	return &Adapter{adaptee: adaptee}
}

func (a *Adapter) Request() string {
	return fmt.Sprintf("Adapter: %s", a.adaptee.SpecificRequest())
}

func main() {
	adaptee := &Adaptee{}
	adapter := NewAdapter(adaptee)

	client := &Client{}
	result := client.UseTarget(adapter)
	fmt.Println(result)
}

type Client struct{}

func (c *Client) UseTarget(target Target) string {
	return target.Request()
}
