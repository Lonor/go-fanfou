package fanfou

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// UsersService handles communication with the users related
// methods of the Fanfou API.
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/API-Endpoints#users
type UsersService struct {
	client *Client
}

// User specifies Fanfou's users data structure
type User struct {
	ID                        string  `json:"id,omitempty"`
	Name                      string  `json:"name,omitempty"`
	ScreenName                string  `json:"screen_name,omitempty"`
	Location                  string  `json:"location,omitempty"`
	Gender                    string  `json:"gender,omitempty"`
	Birthday                  string  `json:"birthday,omitempty"`
	Description               string  `json:"description,omitempty"`
	ProfileImageURL           string  `json:"profile_image_url,omitempty"`
	ProfileImageURLLarge      string  `json:"profile_image_url_large,omitempty"`
	URL                       string  `json:"url,omitempty"`
	Protected                 bool    `json:"protected,omitempty"`
	FollowersCount            int64   `json:"followers_count,omitempty"`
	FriendsCount              int64   `json:"friends_count,omitempty"`
	FavouritesCount           int64   `json:"favourites_count,omitempty"`
	StatusesCount             int64   `json:"statuses_count,omitempty"`
	Following                 bool    `json:"following,omitempty"`
	Notifications             bool    `json:"notifications,omitempty"`
	CreatedAt                 string  `json:"created_at,omitempty"`
	UtcOffset                 int64   `json:"utc_offset,omitempty"`
	ProfileBackgroundColor    string  `json:"profile_background_color,omitempty"`
	ProfileTextColor          string  `json:"profile_text_color,omitempty"`
	ProfileLinkColor          string  `json:"profile_link_color,omitempty"`
	ProfileSidebarFillColor   string  `json:"profile_sidebar_fill_color,omitempty"`
	ProfileSidebarBorderColor string  `json:"profile_sidebar_border_color,omitempty"`
	ProfileBackgroundImageURL string  `json:"profile_background_image_url,omitempty"`
	ProfileBackgroundTile     bool    `json:"profile_background_tile,omitempty"`
	Status                    *Status `json:"status,omitempty"`
}

// Tag specifies Fanfou's tags data structure
type Tag string

// UsersOptParams specifies the optional params for statuses API
type UsersOptParams struct {
	ID     string
	Page   int64
	Count  int64
	Mode   string
	Format string
}

// Tagged shall get users by tag
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/users.tagged
func (s *UsersService) Tagged(Tag string, opt *UsersOptParams) ([]User, error) {
	u := fmt.Sprintf("users/tagged.json")
	params := url.Values{}
	params.Add("tag", Tag)

	if opt != nil {
		if opt.Count != 0 {
			params.Add("count", strconv.FormatInt(opt.Count, 10))
		}
		if opt.Page != 0 {
			params.Add("page", strconv.FormatInt(opt.Page, 10))
		}
		if opt.Mode != "" {
			params.Add("mode", opt.Mode)
		}
		if opt.Format != "" {
			params.Add("format", opt.Format)
		}
	}

	u += "?" + params.Encode()

	req, err := s.client.NewRequest(http.MethodGet, u, "")
	if err != nil {
		return nil, err
	}

	newUsers := new([]User)
	_, err = s.client.Do(req, newUsers)
	if err != nil {
		return nil, err
	}

	return *newUsers, nil
}

// Show shall get a user by ID, or the current user if not specified
// ID represents user ID
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/users.show
func (s *UsersService) Show(opt *UsersOptParams) (*User, error) {
	u := fmt.Sprintf("users/show.json")
	params := url.Values{}

	if opt != nil {
		if opt.ID != "" {
			params.Add("id", opt.ID)
		}
		if opt.Mode != "" {
			params.Add("mode", opt.Mode)
		}
		if opt.Format != "" {
			params.Add("format", opt.Format)
		}
	}

	u += "?" + params.Encode()

	req, err := s.client.NewRequest(http.MethodGet, u, "")
	if err != nil {
		return nil, err
	}

	newUser := new(User)
	_, err = s.client.Do(req, newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

// TagList shall get tags of a specified user or of the current user if not specified
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/users.tag-list
func (s *UsersService) TagList(opt *UsersOptParams) ([]Tag, error) {
	u := fmt.Sprintf("users/tag_list.json")
	params := url.Values{}

	if opt != nil {
		if opt.ID != "" {
			params.Add("id", opt.ID)
		}
	}

	u += "?" + params.Encode()

	req, err := s.client.NewRequest(http.MethodGet, u, "")
	if err != nil {
		return nil, err
	}

	newTags := new([]Tag)
	_, err = s.client.Do(req, newTags)
	if err != nil {
		return nil, err
	}

	return *newTags, nil
}

// Followers shall get recently logged in followers
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/users.followers
func (s *UsersService) Followers(opt *UsersOptParams) ([]User, error) {
	u := fmt.Sprintf("users/followers.json")
	params := url.Values{}

	if opt != nil {
		if opt.ID != "" {
			params.Add("id", opt.ID)
		}
		if opt.Count != 0 {
			params.Add("count", strconv.FormatInt(opt.Count, 10))
		}
		if opt.Page != 0 {
			params.Add("page", strconv.FormatInt(opt.Page, 10))
		}
		if opt.Mode != "" {
			params.Add("mode", opt.Mode)
		}
		if opt.Format != "" {
			params.Add("format", opt.Format)
		}
	}

	u += "?" + params.Encode()

	req, err := s.client.NewRequest(http.MethodGet, u, "")
	if err != nil {
		return nil, err
	}

	newUsers := new([]User)
	_, err = s.client.Do(req, newUsers)
	if err != nil {
		return nil, err
	}

	return *newUsers, nil
}

// Friends shall get recently logged in friends
// ID represents user ID
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/users.friends
func (s *UsersService) Friends(opt *UsersOptParams) ([]User, error) {
	u := fmt.Sprintf("users/friends.json")
	params := url.Values{}

	if opt != nil {
		if opt.ID != "" {
			params.Add("id", opt.ID)
		}
		if opt.Count != 0 {
			params.Add("count", strconv.FormatInt(opt.Count, 10))
		}
		if opt.Page != 0 {
			params.Add("page", strconv.FormatInt(opt.Page, 10))
		}
		if opt.Mode != "" {
			params.Add("mode", opt.Mode)
		}
		if opt.Format != "" {
			params.Add("format", opt.Format)
		}
	}

	u += "?" + params.Encode()

	req, err := s.client.NewRequest(http.MethodGet, u, "")
	if err != nil {
		return nil, err
	}

	newUsers := new([]User)
	_, err = s.client.Do(req, newUsers)
	if err != nil {
		return nil, err
	}

	return *newUsers, nil
}

// Recommendation shall get users recommended by Fanfou
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/users.recommendation
func (s *UsersService) Recommendation(opt *UsersOptParams) ([]User, error) {
	u := fmt.Sprintf("2/users/recommendation.json")
	params := url.Values{}

	if opt != nil {
		if opt.Count != 0 {
			params.Add("count", strconv.FormatInt(opt.Count, 10))
		}
		if opt.Page != 0 {
			params.Add("page", strconv.FormatInt(opt.Page, 10))
		}
		if opt.Mode != "" {
			params.Add("mode", opt.Mode)
		}
		if opt.Format != "" {
			params.Add("format", opt.Format)
		}
	}

	u += "?" + params.Encode()

	req, err := s.client.NewRequest(http.MethodGet, u, "")
	if err != nil {
		return nil, err
	}

	newUsers := new([]User)
	_, err = s.client.Do(req, newUsers)
	if err != nil {
		return nil, err
	}

	return *newUsers, nil
}

// CancelRecommendation shall dismiss user recommended by Fanfou by ID
// ID represents the user ID
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/users.cancel_recommendation
func (s *UsersService) CancelRecommendation(ID string, opt *UsersOptParams) (*User, error) {
	u := fmt.Sprintf("2/users/cancel_recommendation.json")
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

	newUser := new(User)
	_, err = s.client.Do(req, newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}