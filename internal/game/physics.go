package game

import "math"

func ResolveCollision(a, b *Player) {
	if !a.Alive || !b.Alive {
		return
	}

	dx := b.X - a.X
	dy := b.Y - a.Y

	dist := math.Hypot(dx, dy)
	minDist := a.Radius + b.Radius

	if dist == 0 || dist >= minDist {
		return
	}

	// normal
	nx := dx / dist
	ny := dy / dist

	// separa (evita grudar)
	overlap := minDist - dist
	a.X -= nx * overlap / 2
	a.Y -= ny * overlap / 2
	b.X += nx * overlap / 2
	b.Y += ny * overlap / 2

	// ===== FÍSICA DE MOMENTUM =====
	// projeta velocidades no eixo da colisão

	va := a.VelX*nx + a.VelY*ny
	vb := b.VelX*nx + b.VelY*ny

	relative := vb - va

	if relative > 0 {
		return
	}

	const strength = 1.8

	impulse := -relative * strength

	a.VelX -= impulse * nx
	a.VelY -= impulse * ny

	b.VelX += impulse * nx
	b.VelY += impulse * ny
}
