package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type ProjectsBudgetsService struct {
	client          *Client
	currencyCodes   []string
	currencyIDs     []int
	projectType     ProjectType
	lang            string
	currencyDetails bool
}

func (s *ProjectsBudgetsService) Do(ctx context.Context) (*ResponseBudgets, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "projects/0.1/budgets/",
	}
	if s.currencyCodes != nil {
		r.addParam("currency_codes", s.currencyCodes)
	}
	if s.currencyIDs != nil {
		r.addParam("currency_ids", s.currencyIDs)
	}
	if s.projectType != "" {
		r.addParam("project_type", s.projectType)
	}
	if s.lang != "" {
		r.addParam("lang", s.lang)
	}
	if s.currencyDetails {
		r.addParam("currency_details", s.currencyDetails)
	}
	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &ResponseBudgets{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ProjectsBudgetsService) SetCurrencyCodes(currencyCodes []string) *ProjectsBudgetsService {
	s.currencyCodes = currencyCodes
	return s
}

func (s *ProjectsBudgetsService) SetCurrencyIDs(currencyIDs []int) *ProjectsBudgetsService {
	s.currencyIDs = currencyIDs
	return s
}

func (s *ProjectsBudgetsService) SetProjectType(projectType ProjectType) *ProjectsBudgetsService {
	s.projectType = projectType
	return s
}

func (s *ProjectsBudgetsService) SetLang(lang string) *ProjectsBudgetsService {
	s.lang = lang
	return s
}

func (s *ProjectsBudgetsService) SetCurrencyDetails(currencyDetails bool) *ProjectsBudgetsService {
	s.currencyDetails = currencyDetails
	return s
}
