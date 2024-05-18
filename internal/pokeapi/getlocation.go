package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.Cache.Get(url); ok {
		locationResp := Location{}

		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return Location{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)

	rep, err := c.httpClient.Do(req)

	if err != nil {
		return Location{}, err
	}

	defer rep.Body.Close()

	dat, err := io.ReadAll(rep.Body)

	if err != nil {
		return Location{}, err
	}

	locationResp := Location{}

	err = json.Unmarshal(dat, &locationResp)

	if err != nil {
		return Location{}, err
	}

	c.Cache.Add(url, dat)
	return locationResp, nil

}
