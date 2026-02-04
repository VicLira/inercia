package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenW = 960
	screenH = 540
)

type Game struct {
	player *Player
}

func New() *Game {
	return &Game{
		player: NewPlayer(screenW/2, screenH/2),
	}
}

func (g *Game) Update() error {
	g.player.Update(screenW, screenH)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *Game) Layout(w, h int) (int, int) {
	return screenW, screenH
}
