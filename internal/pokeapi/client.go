package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (c *Client) get(url string, target interface{}) error {
	start := time.Now()
	if val, ok := c.cache.Get(url); ok {
		fmt.Printf("[cache HIT] %s (took %v)\n", url, time.Since(start))
		return json.Unmarshal(val, target)
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	c.cache.Add(url, data)
	fmt.Printf("[cache MISS] %s (took %v)\n", url, time.Since(start))

	return json.Unmarshal(data, target)
}
