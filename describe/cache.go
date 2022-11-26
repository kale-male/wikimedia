package describe

import (
	"context"
	"time"

	"github.com/eko/gocache/v3/cache"
	"github.com/eko/gocache/v3/store"
	gocache "github.com/patrickmn/go-cache"
)

type CachedWikimediaClient struct {
	wrapped      WikimediaClient
	cacheManager cache.Cache[string]
}

// QueryText implements WikimediaClient
func (client *CachedWikimediaClient) QueryText(name string) (string, error) {
	ctx := context.TODO()
	result, _ := client.cacheManager.Get(ctx, name)
	if result != "" {
		return result, nil
	}
	result, err := client.wrapped.QueryText(name)
	if err == nil {
		client.cacheManager.Set(ctx, name, result)
	}
	return result, err
}

func MakeCachedWikimediaClient(wrapped WikimediaClient) WikimediaClient {
	gocacheClient := gocache.New(5*time.Minute, 10*time.Minute)
	gocacheStore := store.NewGoCache(gocacheClient)

	cacheManager := cache.New[string](gocacheStore)
	return &CachedWikimediaClient{wrapped: wrapped, cacheManager: *cacheManager}
}
