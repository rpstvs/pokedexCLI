package pokeapi

import (
	"encoding/json"
	"fmt"
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

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return Location{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)

	if err != nil {
		return Location{}, err
	}

	locationResp := Location{}

	err = json.Unmarshal(dat, &locationResp)

	if err != nil {
		return Location{}, err
	}

	c.Cache.Add(url, dat)
	fmt.Println(locationResp)
	return locationResp, nil

}
