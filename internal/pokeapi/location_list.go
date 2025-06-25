package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespListLocation, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, exists := c.cache.Get(url); exists {
		locationsResp := RespListLocation{}
		if err := json.Unmarshal(val, &locationsResp); err != nil {
			return RespListLocation{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespListLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespListLocation{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespListLocation{}, err
	}

	locationsResp := RespListLocation{}
	if err = json.Unmarshal(data, &locationsResp); err != nil {
		return RespListLocation{}, err
	}

	c.cache.Add(url, data)
	return locationsResp, nil
}
