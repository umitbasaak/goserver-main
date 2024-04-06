package db

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
)

type Database struct {
	Client *redis.Client
}

var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx    = context.TODO()
)

func NewDatabase(adress string) (*Database, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     adress,
		Password: "135980Aa@",
		DB:       0,
	})

	if err := client.Ping(Ctx).Err(); err != nil {
		return nil, err
	}
	return &Database{
		Client: client,
	}, nil
}
