package main

const (
	SNAKE_UP int = iota
	SNAKE_DOWN
	SNAKE_LEFT
	SNAKE_RIGHT
)

type Node struct {
	Row  int
	Col  int
	Next *Node
}

type Snake struct {
	Head Node   //snake header
	Body []Node //contain head
	Dir  int    //snake direction
}

func NewSnake() *Snake {
	var snake Snake
	snake.Head = Node{
		Row:  MAXSIZE / 2,
		Col:  MAXSIZE / 2,
		Next: nil,
	}
	snake.Body = make([]Node, 0)
	snake.Body = append(snake.Body, snake.Head) //create a head for body
	snake.Dir = SNAKE_RIGHT
	return &snake
}

func (c *Snake) Run() {
	body_len := len(c.Body)
	if body_len > 1 {
		for k, _ := range c.Body {
			n := body_len - k - 1
			if n == 0 {
				break
			}

			c.Body[n].Row = c.Body[n-1].Row
			c.Body[n].Col = c.Body[n-1].Col

		}
	}

	switch c.Dir {
	case SNAKE_UP:
		c.Head.Row--
	case SNAKE_DOWN:
		c.Head.Row++
	case SNAKE_LEFT:
		c.Head.Col--
	case SNAKE_RIGHT:
		c.Head.Col++
	}
	c.Body[0].Row = c.Head.Row
	c.Body[0].Col = c.Head.Col
}

//detect snake if it bite itsef or hit the wall
func (c *Snake) HitDetect() bool {
	switch c.Dir {
	case SNAKE_UP:

		if c.Head.Row <= WALL_UP_LIMIT {
			return true
		}
	case SNAKE_DOWN:
		if c.Head.Row >= WALL_DOWN_LIMIT {
			return true
		}
	case SNAKE_LEFT:

		if c.Head.Col <= WALL_LEFT_LIMIT {
			return true
		}
	case SNAKE_RIGHT:

		if c.Head.Col >= WALL_RIGHT_LIMIT {
			return true
		}
	}

	for k, v := range c.Body {

		if k > 0 && v.Row == c.Head.Row && v.Col == c.Head.Col { //snake bite itself
			return true
		}
	}
	return false
}

func (c *Snake) Eat(node *Food) bool {
	var nextNode Node
	body_len := len(c.Body)
	if c.Head.Row == node.Row && c.Head.Col == node.Col {
		nextNode = c.Body[body_len-1]
		switch c.Dir {
		case SNAKE_UP:
			nextNode.Row++
		case SNAKE_DOWN:
			nextNode.Row--
		case SNAKE_LEFT:
			nextNode.Col++
		case SNAKE_RIGHT:
			nextNode.Col--
		}
		c.Body[body_len-1].Next = &nextNode
		c.Body = append(c.Body, nextNode)
		return true
	}
	return false
}
