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
	Result            string   `json:"result,omitempty"`
	Reason            string   `json:"reason,omitempty"`
	FirstTime         bool     `json:"firstTime,omitempty"`
	ProductConfigured bool     `json:"productConfigured,omitempty"`
	UserInfo          userInfo `json:"userInfo,omitempty"`
	CanTry            canTry   `json:"canTry,omitempty"`
}

type userInfo struct {
	Id                 int      `json:"id,omitempty"`
	Name               string   `json:"name,omitempty"`
	IsMasterAdmin      bool     `json:"isMasterAdmin,omitempty"`
	IsAdmin            bool     `json:"isAdmin,omitempty"`
	Permissions        []string `json:"permissions,omitempty"`
	Tid                string   `json:"tid,omitempty"`
	Email              string   `json:"email,omitempty"`
	FirstLoginRelative int      `json:"firstloginRelative,omitempty"`
}

type canTry struct {
	IsPossible     bool `json:"isPossible,omitempty"`
	WaitTimeLeft   int  `json:"waitTimeLeft,omitempty"`
	FailedAttempts int  `json:"failedAttempts,omitempty"`
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
