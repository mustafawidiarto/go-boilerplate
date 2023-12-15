package persist

import (
	"context"

	"github.com/mustafawidiarto/go-boilerplate/model/entity"

	"gorm.io/gorm"
)

type pgsqlRoom struct {
	db *gorm.DB
}

// NewPgsqlRoom creates and returns a new instance of the `pgsqlRoom` struct which implements the `RoomRepository` interface.
// It takes a `gorm.DB` object as an argument, which is used to connect to a PostgreSQL database.
func NewPgsqlRoom(db *gorm.DB) *pgsqlRoom {
	return &pgsqlRoom{db}
}

func (r *pgsqlRoom) Save(ctx context.Context, room *entity.Room) (err error) {
	err = r.db.WithContext(ctx).Save(room).Error
	return
}

func (r *pgsqlRoom) Fetch(ctx context.Context) (rooms entity.Rooms, err error) {
	err = r.db.WithContext(ctx).Find(&rooms).Error
	return
}

func (r *pgsqlRoom) GetByID(ctx context.Context, id int64) (room entity.Room, err error) {
	err = r.db.WithContext(ctx).First(&room, id).Error
	return
}

func (r *pgsqlRoom) DeleteBy(ctx context.Context, query map[string]interface{}) (err error) {
	err = r.db.WithContext(ctx).Delete(&entity.Room{}, query).Error
	return
}
