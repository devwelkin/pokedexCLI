package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/twomotive/pokedex/internal/pokecache"
)

// Client handles API interactions
type Client struct {
	httpClient *http.Client
	baseURL    string
	cache      *pokecache.Cache
}

// NewClient creates a new API client
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
		baseURL: "https://pokeapi.co/api/v2",
		cache:   pokecache.NewCache(5 * time.Minute),
	}
}

// GetLocationAreas retrieves location areas with pagination
func (c *Client) GetLocationAreas(url string) (*LocationAreasResponse, error) {
	if url == "" {
		url = c.baseURL + "/location-area"
	}

	// Check if result is in cache
	if cachedData, found := c.cache.Get(url); found {
		var result LocationAreasResponse
		err := json.Unmarshal(cachedData, &result)
		if err == nil {
			return &result, nil
		}
		// If unmarshal fails, continue with the regular request
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result LocationAreasResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("JSON decode error: %v", err)
	}

	// Cache the successful response
	jsonData, err := json.Marshal(result)
	if err == nil {
		c.cache.Add(url, jsonData)
	}

	return &result, nil
}

func (c *Client) GetPokemons(url string, area string) (*PokemonLocationArea, error) {
	if url == "" {
		url = c.baseURL + "/location-area/" + area
	}

	// Check if result is in cache
	if cachedData, found := c.cache.Get(url); found {
		var result PokemonLocationArea
		err := json.Unmarshal(cachedData, &result)
		if err == nil {
			return &result, nil
		}
		// If unmarshal fails, continue with the regular request
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result PokemonLocationArea
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("JSON decode error: %v", err)
	}

	// Cache the successful response
	jsonData, err := json.Marshal(result)
	if err == nil {
		c.cache.Add(url, jsonData)
	}

	return &result, nil
}

func (c *Client) GetPokemonStats(url, name string) (*Pokemon, error) {
	if url == "" {
		url = c.baseURL + "/pokemon/" + name
	}

	// Check if result is in cache
	if cachedData, found := c.cache.Get(url); found {
		var result Pokemon
		err := json.Unmarshal(cachedData, &result)
		if err == nil {
			return &result, nil
		}
		// If unmarshal fails, continue with the regular request
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result Pokemon
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("JSON decode error: %v", err)
	}

	// Cache the successful response
	jsonData, err := json.Marshal(result)
	if err == nil {
		c.cache.Add(url, jsonData)
	}

	return &result, nil

}
