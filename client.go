package freelancer

import (
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
}

func NewClient(apiToken string) *Client {
	return &Client{
		logger:     log.Default(),
		HTTPClient: &http.Client{},
		apiToken:   apiToken,
		baseURL:    getBaseURL(),
	}
}

func (c *Client) NewProjectService() *ProjectService {
	return &ProjectService{client: c}
}

func (c *Client) NewUserService() *UserService {
	return &UserService{client: c}
}
