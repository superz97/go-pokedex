package pokeapi

import "github.com/superz97/go-pokedex/internal/pokecache"

type Client struct {
	cache *pokecache.Cache
}

func NewClient(cache *pokecache.Cache) Client {
	return Client{cache: cache}
}
