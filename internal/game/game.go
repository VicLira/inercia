package game

import (
	"fmt"
	"image/color"
	"inercia/internal/netcode"
	"os"

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

	server *netcode.Server
	client *netcode.NetClient
	isHost bool

	world netcode.WorldState
}

func New() *Game {
	arena := NewArena()

	g := &Game{
		arena: arena,
	}

	if len(os.Args) > 1 && os.Args[1] == "--host" {

		srv, _ := netcode.NewServer("3000")
		g.server = srv
		g.isHost = true

	} else {
		addr := "localhost:3000"

		if len(os.Args) > 1 {
			addr = os.Args[1]
		}

		cl, err := netcode.NewClient(addr)
		if err != nil {
			panic(err)
		}

		g.client = cl
	}

	return g
}

func (g *Game) Update() error {

	if g.isHost {
		g.checkDeaths()

		// cria players conforme clientes conectam
		for len(g.players) < len(g.server.Clients) {
			id := len(g.players) + 1
			name := fmt.Sprintf("New Player%d", id)

			g.players = append(g.players,
				NewPlayer(name, 400, 200, color.White, Controls{}))
		}

		// aplica input
		for i, c := range g.server.Clients {
			g.players[i].ApplyInput(c.Input)
		}

		for _, p := range g.players {
			p.Update(g.arena)
		}

		// colisão entre todos players
		for i := 0; i < len(g.players); i++ {
			for j := i + 1; j < len(g.players); j++ {
				ResolveCollision(g.players[i], g.players[j])
			}
		}

		// monta estado
		var state netcode.WorldState

		for _, p := range g.players {
			state.Players = append(state.Players, netcode.PlayerState{
				X: p.X, Y: p.Y,
				Score: p.Score,
				Alive: p.Alive,
				Name:  p.Name,
			})
		}

		g.server.Broadcast(state)

	} else {

		in := netcode.Input{
			Up:    ebiten.IsKeyPressed(ebiten.KeyW),
			Down:  ebiten.IsKeyPressed(ebiten.KeyS),
			Left:  ebiten.IsKeyPressed(ebiten.KeyA),
			Right: ebiten.IsKeyPressed(ebiten.KeyD),
		}

		g.client.SendInput(in)

		g.client.Receive(&g.world)

		// cria players só se faltar
		for len(g.players) < len(g.world.Players) {
			ps := g.world.Players[len(g.players)]

			p := NewPlayer(ps.Name, ps.X, ps.Y, color.White, Controls{})
			g.players = append(g.players, p)
		}

		// só atualiza estado
		for i, ps := range g.world.Players {
			p := g.players[i]

			p.X = ps.X
			p.Y = ps.Y
			p.Score = ps.Score
			p.Alive = ps.Alive
			p.Name = ps.Name
		}
	}

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

	scoreText := ""

	for _, p := range g.players {
		scoreText += fmt.Sprintf("%s: %d   ", p.Name, p.Score)
	}

	if len(g.players) == 0 {
		scoreText = "Waiting for players..."
	}

	ebitenutil.DebugPrint(screen, scoreText)
}

func (g *Game) Layout(w, h int) (int, int) {
	return screenW, screenH
}
