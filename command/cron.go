package command

import (
	"context"
	"os"
	"time"

	"github.com/mustafawidiarto/go-boilerplate/common"
	"github.com/mustafawidiarto/go-boilerplate/repository/api"
	"github.com/mustafawidiarto/go-boilerplate/repository/cache"
	"github.com/mustafawidiarto/go-boilerplate/repository/persist"
	"github.com/mustafawidiarto/go-boilerplate/usecase"

	"github.com/go-co-op/gocron"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

// RoomUsecase main application business logic hold room usecases
type RoomUsecase interface {
	ExecuteResolvedRoom(ctx context.Context) (err error)
}

type Cron struct {
	roomUC RoomUsecase
}

// NewCron creates a new instance of Cron struct.
func NewCron() *Cron {
	dbConn := common.NewDatabase()
	cacheConn := common.NewCache(os.Getenv("REDIS_URL"))

	roomRepo := persist.NewPgsqlRoom(dbConn)
	roomCacheRepo := cache.NewRedisRoom(cacheConn, 10*time.Minute)
	omniRepo := api.NewApiQismo(os.Getenv("QISCUS_APP_ID"), os.Getenv("QISCUS_SECRET_KEY"))

	roomUC := usecase.NewRoom(roomRepo, omniRepo, roomCacheRepo)

	cron := &Cron{
		roomUC: roomUC,
	}

	return cron
}

// Run starts the cron job and schedules it to execute every minute.
func (c *Cron) Run() {
	log.Info().Msg("cron is started")

	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Minute().Do(func() {
		reqID := uuid.New().String()
		ctx := log.With().Str("request_id", reqID).Logger().WithContext(context.Background())

		c.roomUC.ExecuteResolvedRoom(ctx)
	})

	s.StartBlocking()
}
