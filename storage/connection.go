package storage

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type Storage struct {
	RedisClient redis.Client
	ctx         context.Context
}

func (store *Storage) Init() {
	store.RedisClient = *redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
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

func (store *Storage) Close() {
	if err := store.RedisClient.Close(); err != nil {
		log.Fatal(err)
	}
}

func (store *Storage) SetPulled(status bool) {
	var value int
	if status {
		value = 1
	} else {
		value = 0
	}
	err := store.RedisClient.Set(store.ctx, "pulled", value, 0).Err()
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
		fmt.Println(err)
		return
	}
	err = store.RedisClient.Set(store.ctx, strconv.Itoa(index), str, 0).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = store.RedisClient.Set(store.ctx, "length", index+1, 0).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	store.SetPulled(false)
}
