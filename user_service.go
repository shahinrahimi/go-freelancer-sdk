package freelancer

import "net/http"

type UserService struct {
	client *Client
}

func (s *UserService) GetUser(id int) (*User, *http.Response, error) {
	return nil, nil, nil
}
