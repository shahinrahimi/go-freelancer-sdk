package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ProjectFrontendStatus string

const (
	ProjectFrontendStatusOpen           ProjectFrontendStatus = "open"
	ProjectFrontendStatusComplete       ProjectFrontendStatus = "complete"
	ProjectFrontendStatusPending        ProjectFrontendStatus = "pending"
	ProjectFrontendStatusDraft          ProjectFrontendStatus = "draft"
	ProjectFrontendStatusWorkInProgress ProjectFrontendStatus = "work_in_progress"
)

var (
	BaseAPIMainURL    = "https://www.freelancer.com/api/"
	BaseAPISandBoxURL = "https://api-sandbox.freelancer.com/"
)

var UseSandBox = false

func getBaseURL() string {
	if UseSandBox {
		return BaseAPISandBoxURL
	}
	return BaseAPIMainURL
}

type Client struct {
	logger     *log.Logger
	HTTPClient *http.Client
	apiToken   string
	baseURL    string
	Debug      bool
}

func NewClient(apiToken string) *Client {
	return &Client{
		logger:     log.Default(),
		HTTPClient: &http.Client{},
		apiToken:   apiToken,
		baseURL:    getBaseURL(),
	}
}

func (c *Client) debug(format string, v ...interface{}) {
	if c.Debug {
		c.logger.Printf(format, v...)
	}
}

func (c *Client) parseRequest(r *request) (err error) {

	err = r.validate()
	if err != nil {
		return err
	}

	fullURL := fmt.Sprintf("%s%s", c.baseURL, r.endpoint)
	queryString := r.query.Encode()
	header := http.Header{}
	header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiToken))
	header.Add("Content-Type", "application/json")
	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}
	c.debug("full url: %s\n", fullURL)

	r.fullURL = fullURL
	r.header = header
	return nil
}

func (c *Client) callAPI(ctx context.Context, r *request) (data []byte, err error) {
	err = c.parseRequest(r)
	if err != nil {
		return []byte{}, err
	}
	req, err := http.NewRequest(r.method, r.fullURL, r.body)
	if err != nil {
		return []byte{}, err
	}
	req = req.WithContext(ctx)
	req.Header = r.header
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	defer func() {
		cerr := res.Body.Close()
		// Only overwrite the returned error if the original error was nil and an
		// error occurred while closing the body.
		if err == nil && cerr != nil {
			err = cerr
		}
	}()

	if res.StatusCode >= http.StatusBadRequest {
		apiErr := new(APIError2)
		e := json.Unmarshal(data, apiErr)
		if e != nil {
			c.debug("failed to unmarshal json: %s\n", e)
		}
		if !apiErr.IsValid() {
			apiErr.Response = data
		}
		return nil, apiErr
	}

	return data, nil

}

func (c *Client) NewProjectService() *ProjectService {
	return &ProjectService{client: c}
}

func (c *Client) NewUserService() *UserService {
	return &UserService{client: c}
}
