package describe

import (
	"context"

	"github.com/eko/gocache/v3/cache"
	"github.com/eko/gocache/v3/store"
	gocache "github.com/patrickmn/go-cache"
)

type CachedWikimediaClient struct {
	wrapped      WikimediaClient
	cacheManager cache.Cache[string]
}

// QueryText implements WikimediaClient
func (client *CachedWikimediaClient) QueryText(ctx context.Context, name string) (string, error) {
	result, _ := client.cacheManager.Get(ctx, name)
	if result != "" {
		return result, nil
	}
	result, err := client.wrapped.QueryText(ctx, name)
	if err == nil {
		client.cacheManager.Set(ctx, name, result)
	}
	return result, err
}

func MakeCachedWikimediaClient(cfg *Config, wrapped WikimediaClient) WikimediaClient {
	gocacheClient := gocache.New(cfg.CacheTTL, 2*cfg.CacheTTL)
	gocacheStore := store.NewGoCache(gocacheClient)

	cacheManager := cache.New[string](gocacheStore)
	return &CachedWikimediaClient{wrapped: wrapped, cacheManager: *cacheManager}
}
