package main

const (
	MAXSIZE          int = 22
	WALL_UP_LIMIT    int = 0
	WALL_DOWN_LIMIT  int = MAXSIZE - 1
	WALL_LEFT_LIMIT  int = 0
	WALL_RIGHT_LIMIT int = MAXSIZE - 1

	MOD_SNAKE int = 0
	MOD_FOOD  int = 1
	MOD_WALL  int = 2
	MOD_BLANK int = 3
)

type Map struct {
	Site [MAXSIZE][MAXSIZE]int
}

func NewMap() *Map {
	var site [MAXSIZE][MAXSIZE]int
	for i := 0; i < MAXSIZE; i++ {
		site[0][i] = MOD_WALL
		site[MAXSIZE-1][i] = MOD_WALL
		site[i][0] = MOD_WALL
		site[i][MAXSIZE-1] = MOD_WALL
	}

	for row := 1; row < MAXSIZE-1; row++ {
		for col := 1; col < MAXSIZE-1; col++ {
			site[row][col] = MOD_BLANK
		}
	}
	return &Map{
		Site: site,
	}
}

//write snake and food to map for display on terminal
func (c *Map) Write(snake *Snake, food *Food) {
	for _, v := range snake.Body {
		c.Site[v.Row][v.Col] = MOD_SNAKE
	}

	c.Site[food.Row][food.Col] = MOD_FOOD
}

func (c *Map) Clear() {

	for row := 1; row < MAXSIZE-1; row++ {
		for col := 1; col < MAXSIZE-1; col++ {
			c.Site[row][col] = MOD_BLANK
		}
	}
}
