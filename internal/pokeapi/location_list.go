package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreasResp, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationResp := LocationAreasResp{}
		err := json.Unmarshal([]byte(val), &locationResp)
		if err != nil {
			return LocationAreasResp{}, err
		}

		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	var locationAreas LocationAreasResp
	err = json.Unmarshal(dat, &locationAreas)
	if err != nil {
		return LocationAreasResp{}, err
	}

	c.cache.Add(url, dat)
	return locationAreas, nil
}
