package main

import (
	"github.com/eiannone/keyboard"
)

type Key struct {
	Val string
}

func NewKey() *Key {
	key := &Key{
		Val: "",
	}

	return key
}

func (c *Key) Get() {
	c.Val = ""
	temp, _, _ := keyboard.GetSingleKey()
	c.Val = string(temp)
}
