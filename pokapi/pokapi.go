package pokapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Srikumar6529/pokedexcli/pokecache"
)

// LocationArea is a single location entry in a location-area listing.
type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// LocationAreas is the response returned by the location-area endpoint.
type LocationAreas struct {
	Count   int            `json:"count"`
	Next    string         `json:"next"`
	Prev    string         `json:"previous"`
	Results []LocationArea `json:"results"`
}

// Client fetches data from the PokeAPI, caching responses so repeated
// requests for the same URL don't hit the network again.
type Client struct {
	cache      *pokecache.Cache
	httpClient *http.Client
}

// NewClient returns a Client whose cache entries are reaped once they are
// older than the given interval.
func NewClient(interval time.Duration) *Client {
	return &Client{
		cache:      pokecache.NewCache(interval),
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

// GetLocationAreas returns the location areas found at the given URL. The
// response is served from the cache when possible, otherwise it is fetched
// from the network and then cached.
func (c *Client) GetLocationAreas(url string) (LocationAreas, error) {
	if body, ok := c.cache.Get(url); ok {
		var res LocationAreas
		if err := json.Unmarshal(body, &res); err != nil {
			return LocationAreas{}, fmt.Errorf("unmarshalling cached response: %w", err)
		}
		return res, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return LocationAreas{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreas{}, err
	}
	if res.StatusCode != http.StatusOK {
		return LocationAreas{}, fmt.Errorf("bad status: %s", res.Status)
	}

	var parsed LocationAreas
	if err := json.Unmarshal(body, &parsed); err != nil {
		return LocationAreas{}, err
	}

	c.cache.Add(url, body)
	return parsed, nil
}
