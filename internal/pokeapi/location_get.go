package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLotication(locationName string) (RespGetLocation, error) {
	url := baseURL + "/location-area/" + locationName

	if val, exists := c.cache.Get(url); exists {
		locationsResp := RespGetLocation{}
		if err := json.Unmarshal(val, &locationsResp); err != nil {
			return RespGetLocation{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespGetLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespGetLocation{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespGetLocation{}, err
	}

	location := RespGetLocation{}
	if err = json.Unmarshal(data, &location); err != nil {
		return RespGetLocation{}, err
	}

	c.cache.Add(url, data)
	return location, nil
}
