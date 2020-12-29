package nakivo

import (
	"context"
	"net/http"
)

const (
	AuthenticationAction = "AuthenticationManagement"
)

type AuthenticationService service

type LoginInfo struct {
	Result            string   `json:"result"`
	Reason            string   `json:"reason"`
	FirstTime         bool     `json:"firstTime"`
	ProductConfigured bool     `json:"productConfigured"`
	UserInfo          UserInfo `json:"userInfo"`
	CanTry            CanTry   `json:"canTry"`
}

type UserInfo struct {
	Id                 int      `json:"id"`
	Name               string   `json:"name"`
	IsMasterAdmin      bool     `json:"isMasterAdmin"`
	IsAdmin            bool     `json:"isAdmin"`
	Permissions        []string `json:"permissions"`
	Tid                string   `json:"tid"`
	Email              string   `json:"email"`
	FirstLoginRelative int      `json:"firstloginRelative"`
}

type CanTry struct {
	IsPossible     bool `json:"isPossible"`
	WaitTimeLeft   int  `json:"waitTimeLeft"`
	FailedAttempts int  `json:"failedAttempts"`
}

func (s *AuthenticationService) Login(ctx context.Context, username, password string, remember bool) (*Response, *http.Response, error) {
	request := Request{
		Action: AuthenticationAction,
		Method: "login",
		Data:   []interface{}{username, password, remember},
		Type:   "rpc",
		Tid:    1,
	}

	req, err := s.client.NewRequest(&request)
	if err != nil {
		return nil, nil, err
	}
	r := Response{Data: &LoginInfo{}}
	resp, err := s.client.Do(ctx, req, &r)
	if err != nil {
		return nil, resp, err
	}
	return &r, resp, nil
}

func (s *AuthenticationService) IsLoggedIn(ctx context.Context) (*Response, *http.Response, error) {
	request := Request{
		Action: AuthenticationAction,
		Method: "isLogged",
		Type:   "rpc",
		Tid:    1,
	}

	req, err := s.client.NewRequest(&request)
	if err != nil {
		return nil, nil, err
	}
	r := Response{}
	resp, err := s.client.Do(ctx, req, &r)
	if err != nil {
		return nil, resp, err
	}
	return &r, resp, nil
}

func (s *AuthenticationService) Logout(ctx context.Context) (*Response, *http.Response, error) {
	request := Request{
		Action: AuthenticationAction,
		Method: "logoutCurrentUser",
		Type:   "rpc",
		Tid:    1,
	}

	req, err := s.client.NewRequest(&request)
	if err != nil {
		return nil, nil, err
	}
	r := Response{}
	resp, err := s.client.Do(ctx, req, &r)
	if err != nil {
		return nil, resp, err
	}
	return &r, resp, nil
}
