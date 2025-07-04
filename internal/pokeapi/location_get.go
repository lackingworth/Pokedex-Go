package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(locationName); ok {
		locationResp := Location{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return Location{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return Location{}, errors.New(res.Status)
	}

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	locationResp := Location{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, dat)

	return locationResp, nil
}
