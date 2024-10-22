package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type CurrencyService struct {
	client                    *Client
	currencyCodes             []string
	currencyIDs               []int
	includeExternalCurrencies bool
}

func (s *CurrencyService) Do(ctx context.Context) (*ResponseCurrencies, error) {
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

func (c *Client) NewCurrencyService() *CurrencyService {
	return &CurrencyService{client: c}
}
func (s *CurrencyService) SetCurrencyCodes(codes []string) *CurrencyService {
	s.currencyCodes = codes
	return s
}

func (s *CurrencyService) SetCurrencyIDs(ids []int) *CurrencyService {
	s.currencyIDs = ids
	return s
}

func (s *CurrencyService) SetIncludeExternalCurrencies(include bool) *CurrencyService {
	s.includeExternalCurrencies = include
	return s
}
