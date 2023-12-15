package handler

import (
	"context"

	"github.com/mustafawidiarto/go-boilerplate/model/entity"
)

// RoomUsecase main application business logic hold room usecases
type RoomUsecase interface {
	CreateRoom(ctx context.Context, room *entity.Room) (err error)
	GetRoomByID(ctx context.Context, id int64) (room entity.Room, err error)
}
