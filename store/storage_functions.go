package store

import "fmt"

func SaveUrlMap(shortUrl, initialUrl string) error {
	err := storeService.client.Set(ctx, shortUrl, initialUrl, CacheDuration).Err()
	if err != nil {
		fmt.Printf("Error initializing key:value in redis instance: %v", err)
		return err
	}
	return nil
}

func GetInitialUrl(shortUrl string) (string, error) {
	url, err := storeService.client.Get(ctx, shortUrl).Result()
	if err != nil {
		fmt.Printf("Error retrieving Url value from redis database: %v", err)
		return "", err
	}

	return url, nil
}