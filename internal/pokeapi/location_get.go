package pokeapi

func (c *Client) GetLocationAreaDetail(url string) (LocationAreaDetailResponse, error) {
	var result LocationAreaDetailResponse
	if err := c.get(url, &result); err != nil {
		return LocationAreaDetailResponse{}, err
	}
	return result, nil
}
