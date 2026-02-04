package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenW = 960
	screenH = 540
)

type Game struct {
	players []*Player
}

func New() *Game {
	p1 := NewPlayer(
		300, 270,
		color.RGBA{80, 200, 255, 255},
		Controls{
			Up: ebiten.KeyW, Down: ebiten.KeyS,
			Left: ebiten.KeyA, Right: ebiten.KeyD,
		},
	)

	p2 := NewPlayer(
		660, 270,
		color.RGBA{255, 120, 120, 255},
		Controls{
			Up: ebiten.KeyArrowUp, Down: ebiten.KeyArrowDown,
			Left: ebiten.KeyArrowLeft, Right: ebiten.KeyArrowRight,
		},
	)

	return &Game{
		players: []*Player{p1, p2},
	}
}

func (g *Game) Update() error {
	for _, p := range g.players {
		p.Update(screenW, screenH)
	}

	// colisão entre players
	ResolveCollision(g.players[0], g.players[1])

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, p := range g.players {
		p.Draw(screen)
	}
}

func (g *Game) Layout(w, h int) (int, int) {
	return screenW, screenH
}
