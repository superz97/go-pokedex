package pokeapi

func (c *Client) GetLocationAreas(url string) (LocationAreaListResponse, error) {
	var result LocationAreaListResponse
	if err := c.get(url, &result); err != nil {
		return LocationAreaListResponse{}, err
	}
	return result, nil
}
