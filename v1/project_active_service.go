package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type ProjectActiveService struct {
	client          *Client
	query           string
	projectTypes    []ProjectType
	projectUpgrades []ProjectUpgradeType
	sortFields      []SortFieldsType
}

func (c *Client) NewProjectActiveService() *ProjectActiveService {
	return &ProjectActiveService{client: c}
}

func (s *ProjectActiveService) Do(ctx context.Context) (*ResponseProjects, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "projects/0.1/projects/active/",
	}
	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &ResponseProjects{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
