package store

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/go-redis/redis/v8"
)

type StorageService struct {
	client *redis.Client
}

var (
	storeService = &StorageService{}
	ctx = context.Background()
)
var CacheDuration = 6 * time.Hour


func InitializeStore() *StorageService {
	dbPass := os.Getenv("DB_PASS")
	dbURL := os.Getenv("DB_URL")
	dbUser := os.Getenv("DB_USERNAME")

	rdb := redis.NewClient(&redis.Options{
		Addr:     dbURL,
		Username: dbUser,
		Password: dbPass,
		DB:       0,
	})

	ping, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Error init Redis: %v", err)
	}

	fmt.Printf("\nRedis started successfully: %s", ping)
	storeService.client = rdb
	return storeService
}