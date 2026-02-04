package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenW = 960
	screenH = 540
)

type Game struct {
	players []*Player
	arena   *Arena
}

func New() *Game {
	arena := NewArena()

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
		color.RGBA{255, 0, 0, 255},
		Controls{
			Up: ebiten.KeyArrowUp, Down: ebiten.KeyArrowDown,
			Left: ebiten.KeyArrowLeft, Right: ebiten.KeyArrowRight,
		},
	)

	return &Game{
		players: []*Player{p1, p2},
		arena:   arena,
	}
}

func (g *Game) Update() error {
	for _, p := range g.players {
		if p.Alive {
			p.Update(g.arena)
		}
	}

	// colisão entre players
	ResolveCollision(g.players[0], g.players[1])

	g.checkDeaths()

	return nil
}

func (g *Game) checkDeaths() {
	for i, p := range g.players {

		// caiu abaixo da tela = morte
		if !p.Alive {
			other := g.players[1-i]
			other.Score++
			p.Reset(480, 200)
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{20, 20, 30, 255})

	g.arena.Draw(screen)

	for _, p := range g.players {
		p.Draw(screen)
	}

	scoreText := fmt.Sprintf("Azul: %d 		Vermelhor: %d",
		g.players[0].Score,
		g.players[1].Score,
	)

	ebitenutil.DebugPrint(screen, scoreText)
}

func (g *Game) Layout(w, h int) (int, int) {
	return screenW, screenH
}
