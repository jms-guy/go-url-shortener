package store

import "fmt"

func SaveUrlMap(shortUrl, longUrl string) {
	err := storeService.client.Set(ctx, shortUrl, longUrl, CacheDuration).Err()
	if err != nil {
		fmt.Printf("Error initializing key:value in redis instance: %v", err)
		return
	}
}

func GetInitialUrl(shortUrl string) string {
	val, err := storeService.client.Get(ctx, shortUrl).Result()
	if err != nil {
		fmt.Printf("Error retrieving Url value from redis database: %v", err)
		return ""
	}

	return val
}