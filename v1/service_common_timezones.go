package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type CommonTimezonesService struct {
	client        *Client
	timezones     []int
	timezoneNames []string
}

func (c *CommonTimezonesService) Do(ctx context.Context) (*ResponseTimezones, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/common/0.1/timezones",
	}
	if len(c.timezones) > 0 {
		r.setParam("timezones", c.timezones)
	}
	if len(c.timezoneNames) > 0 {
		r.setParam("timezone_names", c.timezoneNames)
	}
	data, err := c.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &ResponseTimezones{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CommonTimezonesService) Timezones(timezones []int) *CommonTimezonesService {
	s.timezones = timezones
	return s
}

func (s *CommonTimezonesService) TimezoneNames(timezoneNames []string) *CommonTimezonesService {
	s.timezoneNames = timezoneNames
	return s
}
