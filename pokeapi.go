package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/superz97/go-pokedex/internal/pokecache"
)

type locationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

func fetchLocationAreas(url string, cache *pokecache.Cache) (locationAreaResponse, error) {
	start := time.Now()
	if val, ok := cache.Get(url); ok {
		fmt.Printf("[cache HIT] %s (took %v)\n", url, time.Since(start))
		var result locationAreaResponse
		if err := json.Unmarshal(val, &result); err != nil {
			return locationAreaResponse{}, err
		}
		return result, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return locationAreaResponse{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationAreaResponse{}, err
	}
	cache.Add(url, data)
	fmt.Printf("[cache MISS] %s (took %v)\n", url, time.Since(start))

	var result locationAreaResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return locationAreaResponse{}, err
	}
	return result, nil
}
