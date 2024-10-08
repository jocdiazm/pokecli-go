package pokeapi

import (
	"net/http"
	"time"

	"github.com/jocdiazmu/pokedexcli/internal/pokecache"
)

type Client struct {
	cache         pokecache.Cache
	httpClient    http.Client
	caughtPokemon map[string]Pokemon
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		cache:         pokecache.NewCache(cacheInterval),
		caughtPokemon: map[string]Pokemon{},
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
