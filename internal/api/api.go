package api

import (
	"net/http"
	"pokedex/internal/cache"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache      *cache.Cache
	httpClient http.Client
}

func NewClient(interval time.Duration) Client {
	return Client{
		cache: cache.NewCache(interval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
