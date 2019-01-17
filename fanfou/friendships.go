package fanfou

import (
	"fmt"
	"net/http"
	"net/url"
)

// FriendshipsService handles communication with the friendships related
// methods of the Fanfou API.
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/API-Endpoints#friendships
type FriendshipsService struct {
	client *Client
}

// FriendsOptParams specifies the optional params for friendships API
type FriendshipsOptParams struct {
	ID     string
	Page   int64
	Count  int64
	Mode   string
	Format string
}

// Create shall add the specified user as a friend (follow)
// ID represents the user ID
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/friendships.create
func (s *FriendshipsService) Create(ID string, opt *FriendshipsOptParams) (*UserResult, error) {
	u := fmt.Sprintf("friendships/create.json")
	params := url.Values{
		"id": []string{ID},
	}

	if opt != nil {
		if opt.Mode != "" {
			params.Add("mode", opt.Mode)
		}
	}

	u += "?" + params.Encode()

	req, err := s.client.NewRequest(http.MethodPost, u, "")
	if err != nil {
		return nil, err
	}

	newUser := new(UserResult)
	_, err = s.client.Do(req, newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

// Destroy shall add the specified user as a friend (follow)
// ID represents the user ID
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/friendships.destroy
func (s *FriendshipsService) Destroy(ID string, opt *FriendshipsOptParams) (*UserResult, error) {
	u := fmt.Sprintf("friendships/destroy.json")
	params := url.Values{
		"id": []string{ID},
	}

	if opt != nil {
		if opt.Mode != "" {
			params.Add("mode", opt.Mode)
		}
		if opt.Format != "" {
			params.Add("format", opt.Format)
		}
	}

	u += "?" + params.Encode()

	req, err := s.client.NewRequest(http.MethodPost, u, "")
	if err != nil {
		return nil, err
	}

	newUser := new(UserResult)
	_, err = s.client.Do(req, newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
