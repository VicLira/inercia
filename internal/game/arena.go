package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Platform struct {
	X, Y float64
	W, H float64
}

type Arena struct {
	Platforms []Platform
}

func NewArena() *Arena {
	return &Arena{
		Platforms: []Platform{

			// chão principal
			{120, 420, 360, 40},

			// paredes laterais
			{120, 220, 20, 200},
			{820, 60, 20, 200},

			// teto
			{120, 60, 360, 20},
		},
	}
}

func (a *Arena) Draw(screen *ebiten.Image) {
	clr := color.RGBA{60, 60, 70, 255}

	for _, p := range a.Platforms {
		img := ebiten.NewImage(int(p.W), int(p.H))
		img.Fill(clr)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(p.X, p.Y)

		screen.DrawImage(img, op)
	}
}
