package forestparty

import (
	"context"
	"time"
)

type BettingTarget struct {
	ID             string
	TotalBetAmount int64
}

type Round struct {
	ID             string
	BettingTargets []BettingTarget
	StartTime      time.Time
	EndTime        time.Time
}



type ForestPartyServicer interface {
	StartGames(ctx context.Context) error
	Round(ctx context.Context, id string) (*Round, error)
}
