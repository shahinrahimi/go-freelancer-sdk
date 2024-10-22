package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type ProjectsCurrenciesService struct {
	client                    *Client
	currencyCodes             []string
	currencyIDs               []int
	includeExternalCurrencies bool
}

func (s *ProjectsCurrenciesService) Do(ctx context.Context) (*ResponseCurrencies, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "projects/0.1/currencies/",
	}
	if s.currencyCodes != nil {
		r.addParam("currency_codes", s.currencyCodes)
	}
	if s.currencyIDs != nil {
		r.addParam("currency_ids", s.currencyIDs)
	}
	if s.includeExternalCurrencies {
		r.addParam("include_external_currencies", s.includeExternalCurrencies)
	}
	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &ResponseCurrencies{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ProjectsCurrenciesService) SetCurrencyCodes(codes []string) *ProjectsCurrenciesService {
	s.currencyCodes = codes
	return s
}

func (s *ProjectsCurrenciesService) SetCurrencyIDs(ids []int) *ProjectsCurrenciesService {
	s.currencyIDs = ids
	return s
}

func (s *ProjectsCurrenciesService) SetIncludeExternalCurrencies(include bool) *ProjectsCurrenciesService {
	s.includeExternalCurrencies = include
	return s
}
