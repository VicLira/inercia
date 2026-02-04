package game

import "math"

func ResolveCollision(a, b *Player) {
	dx := b.X - a.X
	dy := b.Y - a.Y

	dist := math.Hypot(dx, dy)
	minDist := a.Radius + b.Radius

	if dist >= minDist || dist == 0 {
		return
	}

	overlap := minDist - dist
	nx := dx / dist
	ny := dy / dist

	a.X -= nx * overlap / 2
	a.Y -= ny * overlap / 2
	b.X += nx * overlap / 2
	b.Y -= ny * overlap / 2

	// troca velocidades (efeito de impacto)
	a.VelX, b.VelX = b.VelX, a.VelX
	a.VelY, b.VelY = b.VelY, a.VelY
}
