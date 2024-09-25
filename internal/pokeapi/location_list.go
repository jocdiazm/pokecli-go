package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// Check if the location is cached
	if cached, ok := c.cache.Get(url); ok {
		var locationsResp RespShallowLocations
		err := json.Unmarshal(cached, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return RespShallowLocations{}, err
	}
	var locationsResp RespShallowLocations
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// Cache the response
	c.cache.Add(url, data)
	return locationsResp, nil
}
