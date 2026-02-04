package game

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Player struct {
	X, Y   float64
	Radius float64
	Speed  float64
	Color  color.Color
}

func NewPlayer(x, y float64) *Player {
	return &Player{
		X:      x,
		Y:      y,
		Radius: 20,
		Speed:  4,
		Color:  color.RGBA{80, 200, 255, 255},
	}
}
func (p *Player) Update() {
	var dx, dy float64

	// WASD ou setas
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		dx--
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		dx++
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		dy--
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		dy++
	}

	// normaliza diagonal
	if dx != 0 || dy != 0 {
		length := math.Hypot(dx, dy)
		dx /= length
		dy /= length
	}

	p.X += dx * p.Speed
	p.Y += dy * p.Speed
}

func (p *Player) Draw(screen *ebiten.Image) {
	vector.FillCircle(
		screen,
		float32(p.X),
		float32(p.Y),
		float32(p.Radius),
		p.Color,
		false,
	)
}
