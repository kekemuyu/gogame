package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const (
	BLANK_DIS string = "  "
	FOOD_DIS  string = " @"
	SNAKE_DIS string = " *"
	WALL_DIS  string = " #"

	SPEED = 500
)

func disClear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func display(snakeMap *Map) {
	var temps string
	site := snakeMap.Site
	disClear()
	for row := 0; row < MAXSIZE; row++ {
		for col := 0; col < MAXSIZE; col++ {
			switch site[row][col] {
			case MOD_BLANK:
				temps = BLANK_DIS
			case MOD_FOOD:
				temps = FOOD_DIS
			case MOD_SNAKE:
				temps = SNAKE_DIS
			case MOD_WALL:
				temps = WALL_DIS
			}
			fmt.Printf("%s", temps)
		}

		fmt.Printf("\n")

	}

}

func command(key string, snake *Snake) {
	switch key {
	case "w":
		if snake.Dir != SNAKE_DOWN {
			snake.Dir = SNAKE_UP
		}
	case "s":
		if snake.Dir != SNAKE_UP {
			snake.Dir = SNAKE_DOWN
		}

	case "a":
		if snake.Dir != SNAKE_RIGHT {
			snake.Dir = SNAKE_LEFT
		}

	case "d":
		if snake.Dir != SNAKE_LEFT {
			snake.Dir = SNAKE_RIGHT
		}

	}
}

func main() {
	snakeMap := NewMap()
	snake := NewSnake()
	food := NewFood()
	snakeMap.Write(snake, food)

	display(snakeMap)
	key := NewKey()

	gameover := make(chan bool)
	go func() {
		for {
			time.Sleep(time.Millisecond * SPEED)
			snake.Run()
			if snake.Eat(food) {
				food.Generate()
			}
			snakeMap.Clear()
			snakeMap.Write(snake, food)
			display(snakeMap)
			if snake.HitDetect() {
				gameover <- true
			}
		}
	}()

	go func() {
		for {
			key.Get()
			command(key.Val, snake)
		}
	}()
	<-gameover
	fmt.Println("game over")
}
