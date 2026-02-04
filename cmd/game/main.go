package main

import (
	"log"

	"inercia/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(960, 540)
	ebiten.SetWindowTitle("Inércia")

	g := game.New()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
