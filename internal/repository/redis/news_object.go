package redis

import (
	"context"
	"encoding/json"
	"github.com/Mufidzz/bareksa-test/pkg/response"
)

func (redis *Redis) GetObject(key string, dest interface{}) error {
	objectJson := redis.newsClient.Get(context.Background(), key)

	err := json.Unmarshal([]byte(objectJson.Val()), &dest)
	if err != nil {
		return response.InternalError{
			Type:         "Repo",
			Name:         "Redis",
			FunctionName: "GetObject",
			Description:  "Failed Unmarshal JSON",
			Trace:        err,
		}.Error()
	}

	return nil
}

func (redis *Redis) SaveObject(key string, value interface{}) error {
	objectJson, err := json.Marshal(value)
	if err != nil {
		return response.InternalError{
			Type:         "Repo",
			Name:         "Redis",
			FunctionName: "SaveObject",
			Description:  "Failed Marshall JSON",
			Trace:        err,
		}.Error()
	}

	res := redis.newsClient.Set(context.Background(), key, objectJson, REDIS_TIMEOUT_NEWS)
	if res.Err() != nil {
		return response.InternalError{
			Type:         "Repo",
			Name:         "Redis",
			FunctionName: "SaveObject",
			Description:  "Failed to Set into database",
			Trace:        res.Err(),
		}.Error()
	}

	return nil
}

func (redis *Redis) FlushAll() error {
	res := redis.newsClient.FlushAll(context.Background())
	if res.Err() != nil {
		return response.InternalError{
			Type:         "Repo",
			Name:         "Redis",
			FunctionName: "FlushAll",
			Description:  "Failed to Flush database",
			Trace:        res.Err(),
		}.Error()
	}

	return nil
}
