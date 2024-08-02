package component

import (
	"context"
	"log"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/titosunu/wallet-go/core"
)

func GetCacheConnection() core.CacheRepository {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))
	if err != nil {
		log.Fatalf("error when connect cache %s", err.Error())
	}
	return cache
}