package main

import (
	"math/rand"
	"time"
)

type Food struct {
	Row int
	Col int
}

func NewFood() *Food {
	var food Food
	rand.Seed(time.Now().Unix())
	food.Row = 1 + rand.Intn(MAXSIZE-2) //[1,MAXSIZE)
	food.Col = 1 + rand.Intn(MAXSIZE-2)
	return &food
}

func (c *Food) Generate() {
	c.Row = 1 + rand.Intn(MAXSIZE-2) //[1,MAXSIZE)
	c.Col = 1 + rand.Intn(MAXSIZE-2)
}
