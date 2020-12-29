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
	UserInfo          userInfo `json:"userInfo"`
	CanTry            canTry   `json:"canTry"`
}

type userInfo struct {
	Id                 int      `json:"id"`
	Name               string   `json:"name"`
	IsMasterAdmin      bool     `json:"isMasterAdmin"`
	IsAdmin            bool     `json:"isAdmin"`
	Permissions        []string `json:"permissions"`
	Tid                string   `json:"tid"`
	Email              string   `json:"email"`
	FirstLoginRelative int      `json:"firstloginRelative"`
}

type canTry struct {
	IsPossible     bool `json:"isPossible"`
	WaitTimeLeft   int  `json:"waitTimeLeft"`
	FailedAttempts int  `json:"failedAttempts"`
}

func (s *AuthenticationService) Login(ctx context.Context, username, password string, remember bool) (*response, *http.Response, error) {
	request := request{
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
	r := response{Data: &LoginInfo{}}
	resp, err := s.client.Do(ctx, req, &r)
	if err != nil {
		return nil, resp, err
	}
	return &r, resp, nil
}

func (s *AuthenticationService) IsLoggedIn(ctx context.Context) (*response, *http.Response, error) {
	request := request{
		Action: AuthenticationAction,
		Method: "isLogged",
		Type:   "rpc",
		Tid:    1,
	}

	req, err := s.client.NewRequest(&request)
	if err != nil {
		return nil, nil, err
	}
	r := response{}
	resp, err := s.client.Do(ctx, req, &r)
	if err != nil {
		return nil, resp, err
	}
	return &r, resp, nil
}

func (s *AuthenticationService) Logout(ctx context.Context) (*response, *http.Response, error) {
	request := request{
		Action: AuthenticationAction,
		Method: "logoutCurrentUser",
		Type:   "rpc",
		Tid:    1,
	}

	req, err := s.client.NewRequest(&request)
	if err != nil {
		return nil, nil, err
	}
	r := response{}
	resp, err := s.client.Do(ctx, req, &r)
	if err != nil {
		return nil, resp, err
	}
	return &r, resp, nil
}
