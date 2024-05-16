package pokeapi

import (
	"net/http"
	"time"

	"github.com/rpstvs/pokedexCLI/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	Cache      pokecache.Cache
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		Cache: pokecache.NewCache(cacheInterval),
	}
}
