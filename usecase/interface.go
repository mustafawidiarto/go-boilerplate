package usecase

import (
	"context"

	"github.com/mustafawidiarto/go-boilerplate/model/entity"
)

// RoomRepository defines an interface for saving Room data to persistent storage.
type RoomRepository interface {
	Save(ctx context.Context, room *entity.Room) (err error)
	Fetch(ctx context.Context) (rooms entity.Rooms, err error)
	GetByID(ctx context.Context, id int64) (room entity.Room, err error)
	DeleteBy(ctx context.Context, query map[string]interface{}) (err error)
}

// RoomCacheRepository represents a repository that provides cache related operations.
type RoomCacheRepository interface {
	Save(ctx context.Context, room entity.Room) (err error)
	GetByID(ctx context.Context, id int64) (room entity.Room, err error)
	DeletetByID(ctx context.Context, id int64) (err error)
}

// OmnichannelRepository defines an interface for interacting with an omnichannel platform.
type OmnichannelRepository interface {
	CreateRoomTag(ctx context.Context, roomID string, tag string) (err error)
	ResolvedRoom(ctx context.Context, roomID string) (err error)
}
