package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type ProjectsCategoriesService struct {
	client                    *Client
	categories                []int
	jobDetails                bool
	lang                      string
	activeProjectCountDetails bool
	seoDetails                bool
}

func (s *ProjectsCategoriesService) Do(ctx context.Context) (*ResponseCategories, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "projects/0.1/categories/",
	}
	if s.categories != nil {
		r.addParam("categories", s.categories)
	}
	if s.jobDetails {
		r.addParam("job_details", s.jobDetails)
	}
	if s.lang != "" {
		r.addParam("lang", s.lang)
	}
	if s.activeProjectCountDetails {
		r.addParam("active_project_count_details", s.activeProjectCountDetails)
	}
	if s.seoDetails {
		r.addParam("seo_details", s.seoDetails)
	}
	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &ResponseCategories{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ProjectsCategoriesService) SetCategories(categories []int) *ProjectsCategoriesService {
	s.categories = categories
	return s
}

func (s *ProjectsCategoriesService) SetJobDetails(jobDetails bool) *ProjectsCategoriesService {
	s.jobDetails = jobDetails
	return s
}

func (s *ProjectsCategoriesService) SetLang(lang string) *ProjectsCategoriesService {
	s.lang = lang
	return s
}

func (s *ProjectsCategoriesService) SetActiveProjectCountDetails(activeProjectCountDetails bool) *ProjectsCategoriesService {
	s.activeProjectCountDetails = activeProjectCountDetails
	return s
}

func (s *ProjectsCategoriesService) SetSeoDetails(seoDetails bool) *ProjectsCategoriesService {
	s.seoDetails = seoDetails
	return s
}
