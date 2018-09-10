package hall

import (
	"github.com/jasonsoft/forest-party/hub"
)

type Room struct {
	ID          string
	GameID      string
	PlayerCount int
}

func (rm *Room) Join(c *hub.Client) error {
	return nil
}

type FindRoomOptions struct {
	GameID string
}

type RoomServicer interface {
	Room(id string) (*Room, error)
	Rooms(opts FindRoomOptions) ([]*Room, error)
	CreatRoom(room *Room) error
}
