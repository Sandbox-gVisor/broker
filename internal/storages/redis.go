package storages

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	RedisClient redis.Client
	Ctx         context.Context
}

func (store *RedisStorage) Init() {
	store.RedisClient = *redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	store.Ctx = context.Background()

	err := store.RedisClient.Set(store.Ctx, "length", 0, 0).Err()
	if err != nil {
		log.Println(err)
	}
	err = store.RedisClient.Set(store.Ctx, "pulled", 1, 0).Err()
	if err != nil {
		log.Println(err)
	}
}

func (store *RedisStorage) FlushStorage() {
	err := store.RedisClient.Set(store.Ctx, "length", 0, 0).Err()
	if err != nil {
		log.Println(err)
	}
	err = store.RedisClient.Set(store.Ctx, "pulled", 1, 0).Err()
	if err != nil {
		log.Println(err)
	}
}

func (store *RedisStorage) Close() {
	if err := store.RedisClient.Close(); err != nil {
		log.Fatal(err)
	}
}

func (store *RedisStorage) emitUpdate() {
	err := store.RedisClient.Publish(store.Ctx, "update", "payload").Err()
	if err != nil {
		log.Println(err)
	}
}

func (store *RedisStorage) getLength() (int, error) {
	val, err := store.RedisClient.Get(store.Ctx, "length").Result()
	if err != nil {
		return 0, err
	}
	result, err := strconv.Atoi(val)
	return result, err
}

func (store *RedisStorage) SaveMessage(str string) {
	index, err := store.getLength()
	if err != nil {
		log.Println(err)
		return
	}
	err = store.RedisClient.Set(store.Ctx, strconv.Itoa(index), str, 0).Err()
	if err != nil {
		log.Println(err)
		return
	}
	err = store.RedisClient.Set(store.Ctx, "length", index+1, 0).Err()
	if err != nil {
		log.Println(err)
		return
	}

	store.emitUpdate()
}
