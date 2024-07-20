package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(url *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullUrl := baseURL + endpoint
	if url != nil {
		fullUrl = *url
	}

	data, ok := c.cache.Get(fullUrl)
	if ok {
		locationAreasResp := LocationAreasResp{}
		err := json.Unmarshal(data, &locationAreasResp)
		if err != nil {
			return LocationAreasResp{}, err
		}

		return locationAreasResp, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return LocationAreasResp{}, fmt.Errorf("received status code %d", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	c.cache.Add(fullUrl, data)

	return locationAreasResp, nil
}

func (c *Client) GetLocationAreaDetails(url *string) (LocationAreaDetailedResp, error) {
	fullUrl := baseURL + "/location-area/" + *url

	data, ok := c.cache.Get(fullUrl)
	if ok {
		locationAreaDetailsResp := LocationAreaDetailedResp{}
		err := json.Unmarshal(data, &locationAreaDetailsResp)
		if err != nil {
			return LocationAreaDetailedResp{}, err
		}

		return locationAreaDetailsResp, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreaDetailedResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaDetailedResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return LocationAreaDetailedResp{}, fmt.Errorf("location area not found")
	}

	if resp.StatusCode > 299 {
		return LocationAreaDetailedResp{}, fmt.Errorf("received status code %d", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaDetailedResp{}, err
	}

	locationAreaDetailedResp := LocationAreaDetailedResp{}
	err = json.Unmarshal(data, &locationAreaDetailedResp)
	if err != nil {
		return LocationAreaDetailedResp{}, err
	}

	c.cache.Add(fullUrl, data)

	return locationAreaDetailedResp, nil
}
