package pokeapi

import (
	"encoding/json"
	"io"
	"log"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaResp, error) {
	endpoint := "/location"

	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}
	locations := LocationAreaResp{}
	data, hit := c.cache.Get(fullURL)

	if hit {
		err := json.Unmarshal(data, &locations)

		if err != nil {
			return LocationAreaResp{}, err
		}

		return locations, nil
	}

	res, err := c.httpClient.Get(fullURL)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return LocationAreaResp{}, err
	}
	if err != nil {
		log.Fatal(err)
		return LocationAreaResp{}, err
	}

	err = json.Unmarshal(body, &locations)

	if err != nil {
		return LocationAreaResp{}, err
	}

	c.cache.Add(fullURL, body)

	return locations, nil
}
