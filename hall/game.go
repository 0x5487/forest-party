package hall

type Game struct {
	ID          string
	Name        string
	State       int // 1: open 2: disable 3: maintenance 4: coming soon
	PlayerCount int
}

type GameRepository interface {
	Game(id string) (*Game, error)
}

type GameServicer interface {
	Games() ([]Game, error)
}
