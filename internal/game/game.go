package game

import "fmt"

type Game struct {
    Message string
}

func NewGame() *Game {
    return &Game{
        Message: "Hello from the game!",
    }
}

func (g *Game) SendHello() string {
    return fmt.Sprintf("Game says: %s", g.Message)
}