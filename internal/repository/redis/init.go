package redis

import (
	"context"
	"github.com/Mufidzz/bareksa-test/pkg/response"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	newsClient *redis.Client

	//// Use Cmdable for Testing Purpose
	//newsCmdable *redis.Cmdable
}

func New(newsRedisOption redis.Options) (*Redis, error) {
	ctx := context.Background()

	newsRedisClient := redis.NewClient(&newsRedisOption)
	err := newsRedisClient.Ping(ctx).Err()
	if err != nil {
		return nil, response.InternalError{
			Type:         "Repo",
			Name:         "Redis",
			FunctionName: "New",
			Description:  "Failed get connection to Redis Server",
			Trace:        err,
		}.Error()
	}

	return &Redis{
		newsClient: newsRedisClient,
	}, nil
}

func NewFromObject(newsRedisClient *redis.Client) *Redis {
	return &Redis{
		newsClient: newsRedisClient,
	}
}
