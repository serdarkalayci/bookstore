package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

// HealthRepository represent a structure that will communicate to MongoDB to accomplish health related transactions
type HealthRepository struct {
	helper dbHelper
}

func newHealthRepository(client *redis.Client) HealthRepository {
	return HealthRepository{
		helper: redisHelper{client: client},
	}
}

// Ready checks the arangodb connection
func (hr HealthRepository) Ready() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check the connection
	status := hr.helper.Ping(ctx)
	if status == false {
		log.Error().Msg("An error occured while connecting to tha database")
		return false
	}
	log.Info().Msg("Connection to Redis checked successfuly!")
	return true
}
