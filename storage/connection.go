package storage

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type Storage struct {
	RedisClient redis.Client
	ctx         context.Context
}

func (store *Storage) Init() {
	store.RedisClient = *redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	store.ctx = context.Background()

	err := store.RedisClient.Set(store.ctx, "length", 0, 0).Err()
	if err != nil {
		log.Println(err)
	}
	err = store.RedisClient.Set(store.ctx, "pulled", 1, 0).Err()
	if err != nil {
		log.Println(err)
	}
}

func (store *Storage) FlushRedis() {
	err := store.RedisClient.Set(store.ctx, "length", 0, 0).Err()
	if err != nil {
		log.Println(err)
	}
	err = store.RedisClient.Set(store.ctx, "pulled", 1, 0).Err()
	if err != nil {
		log.Println(err)
	}
}

func (store *Storage) Close() {
	if err := store.RedisClient.Close(); err != nil {
		log.Fatal(err)
	}
}

func (store *Storage) EmitUpdate() {
	err := store.RedisClient.Publish(store.ctx, "update", "payload").Err()
	if err != nil {
		log.Println(err)
	}
}

func (store *Storage) GetLength() (int, error) {
	val, err := store.RedisClient.Get(store.ctx, "length").Result()
	if err != nil {
		return 0, err
	}
	result, err := strconv.Atoi(val)
	return result, err
}

func (store *Storage) AddString(str string) {
	index, err := store.GetLength()
	if err != nil {
		log.Println(err)
		return
	}
	err = store.RedisClient.Set(store.ctx, strconv.Itoa(index), str, 0).Err()
	if err != nil {
		log.Println(err)
		return
	}
	err = store.RedisClient.Set(store.ctx, "length", index+1, 0).Err()
	if err != nil {
		log.Println(err)
		return
	}
	store.EmitUpdate()
}
