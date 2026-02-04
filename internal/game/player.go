package game

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Player struct {
	X, Y   float64
	VelX   float64
	VelY   float64
	Radius float64

	Accel    float64
	Friction float64
	Bounce   float64

	Color color.Color
}

func NewPlayer(x, y float64) *Player {
	return &Player{
		X:        x,
		Y:        y,
		Radius:   20,
		Accel:    0.6,
		Friction: 0.96,
		Bounce:   0.9,
		Color:    color.RGBA{80, 200, 255, 255},
	}
}

func (p *Player) Update(screenW, screenH float64) {
	var ax, ay float64

	// WASD ou setas
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		ax--
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		ax++
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		ay--
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		ay++
	}

	// normaliza diagonal
	if ax != 0 || ay != 0 {
		length := math.Hypot(ax, ay)
		ax /= length
		ay /= length
	}

	// física
	p.VelX += ax * p.Accel
	p.VelY += ay * p.Accel

	p.VelX *= p.Friction
	p.VelY *= p.Friction

	p.X += p.VelX
	p.Y += p.VelY

	// colisão com paredes (bounce)
	if p.X-p.Radius < 0 {
		p.X = p.Radius
		p.VelX *= -p.Bounce
	}
	if p.X+p.Radius > screenW {
		p.X = screenW - p.Radius
		p.VelX *= -p.Bounce
	}
	if p.Y-p.Radius < 0 {
		p.Y = p.Radius
		p.VelY *= -p.Bounce
	}
	if p.Y+p.Radius > screenH {
		p.Y = screenH - p.Radius
		p.VelY *= -p.Bounce
	}
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
