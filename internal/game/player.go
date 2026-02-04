package game

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Controls struct {
	Up, Down, Left, Right ebiten.Key
}

type Player struct {
	X, Y   float64
	VelX   float64
	VelY   float64
	Radius float64

	Accel    float64
	Friction float64
	Bounce   float64

	Color    color.Color
	Controls Controls

	Alive bool
	Score int
}

func NewPlayer(x, y float64, c color.Color, controls Controls) *Player {
	p := &Player{
		X:        x,
		Y:        y,
		Radius:   20,
		Accel:    0.6,
		Friction: 0.96,
		Bounce:   0.9,
		Color:    c,
		Controls: controls,
	}

	p.Reset(x, y)
	return p
}

// respawn
func (p *Player) Reset(x, y float64) {
	p.X = x
	p.Y = y
	p.VelX = 0
	p.VelY = 0
	p.Alive = true
}

func (p *Player) Update(arena *Arena) {
	if !p.Alive {
		return
	}

	var ax, ay float64

	// INPUT -> ACELERAÇÃO
	if ebiten.IsKeyPressed(p.Controls.Left) {
		ax--
	}
	if ebiten.IsKeyPressed(p.Controls.Right) {
		ax++
	}
	if ebiten.IsKeyPressed(p.Controls.Up) {
		ay--
	}
	if ebiten.IsKeyPressed(p.Controls.Down) {
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

	// morte, caso cair fora
	const margin = 150

	if p.X < -margin ||
		p.X > screenW+margin ||
		p.Y < -margin ||
		p.Y > screenH+margin {

		p.Alive = false
		return
	}

	// colisão paredes da arena
	for _, plat := range arena.Platforms {

		left := plat.X
		right := plat.X + plat.W
		top := plat.Y
		bottom := plat.Y + plat.H

		playerLeft := p.X - p.Radius
		playerRight := p.X + p.Radius
		playerTop := p.Y - p.Radius
		playerBottom := p.Y + p.Radius

		// colisão AABB simples
		if playerRight > left &&
			playerLeft < right &&
			playerBottom > top &&
			playerTop < bottom {

			// calcula penetração
			penX := math.Min(playerRight-left, right-playerLeft)
			penY := math.Min(playerBottom-top, bottom-playerTop)

			// resolve no menor eixo
			if penX < penY {
				// colisão lateral (parede)

				if p.X < plat.X {
					p.X -= penX
				} else {
					p.X += penX
				}

				p.VelX *= -p.Bounce

			} else {
				// colisão vertical (chão/teto)

				if p.Y < plat.Y {
					p.Y -= penY
				} else {
					p.Y += penY
				}

				p.VelY *= -p.Bounce
			}
		}

	}

}

func (p *Player) Draw(screen *ebiten.Image) {
	if !p.Alive {
		return
	}

	vector.FillCircle(
		screen,
		float32(p.X),
		float32(p.Y),
		float32(p.Radius),
		p.Color,
		false,
	)
}
