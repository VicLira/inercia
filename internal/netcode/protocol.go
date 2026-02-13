package netcode

type Input struct {
	Up, Down, Left, Right bool
}

type PlayerState struct {
	X, Y  float64
	Score int
	Alive bool
	Name  string
}

type WorldState struct {
	Players []PlayerState
}
