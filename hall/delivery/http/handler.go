package http

import (
	"github.com/jasonsoft/forest-party/hall"
	"github.com/jasonsoft/napnap"
)

func NewRouter(handler *HallHttpHandler) *napnap.Router {
	router := napnap.NewRouter()

	router.Post("/v1/members/token", handler.memberTokenCreateEndpoint)
	router.Post("/v1/rooms/:room_id/join", handler.roomJoinEndpoint)
	router.Get("/v1/rooms", handler.roomListEndpoint)
	router.Get("/v1/games", handler.gameListEndpoint)

	return router
}

type HallHttpHandler struct {
	gameSvc hall.GameServicer
	roomSvc hall.RoomServicer
}
