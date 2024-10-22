package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type CommonCountriesService struct {
	client       *Client
	extraDetails bool
}

func (c *CommonCountriesService) Do(ctx context.Context) (*ResponseCountries, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/common/0.1/countries",
	}
	if c.extraDetails {
		r.setParam("extra_details", c.extraDetails)
	}
	data, err := c.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &ResponseCountries{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *CommonCountriesService) SetExtraDetails(extraDetails bool) *CommonCountriesService {
	c.extraDetails = extraDetails
	return c
}
