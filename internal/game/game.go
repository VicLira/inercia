package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player *Player
}

func New() *Game {
	return &Game{
		player: NewPlayer(480, 270), // centro da tela
	}
}

func (g *Game) Update() error {
	g.player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *Game) Layout(w, h int) (int, int) {
	return 960, 540
}
