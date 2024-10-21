package utils

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"log"
	"time"
)

func GetCacheConnection() CacheRepository {

	bigCache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))
	if err != nil {
		log.Fatalf("error when connect cache %v", err.Error())
	}
	return bigCache
}
