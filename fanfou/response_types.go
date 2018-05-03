package fanfou

type responseError struct {
	Request string `json:"request"`
	Error   string `json:"error"`
}

type responseUserID string

type responseTag string

type responseUser struct {
	ID                        responseUserID `json:"id"`
	UniqueID                  string         `json:"unique_id"`
	Name                      string         `json:"name"`
	ScreenName                string         `json:"screen_name"`
	Location                  string         `json:"location"`
	Gender                    string         `json:"gender"`
	Birthday                  string         `json:"birthday"`
	Description               string         `json:"description"`
	ProfileImageURL           string         `json:"profile_image_url"`
	ProfileImageURLLarge      string         `json:"profile_image_url_large"`
	URL                       string         `json:"url"`
	Protected                 bool           `json:"protected"`
	FollowersCount            int            `json:"followers_count"`
	FriendsCount              int            `json:"friends_count"`
	FavouritesCount           int            `json:"favourites_count"`
	StatusesCount             int            `json:"statuses_count"`
	Following                 bool           `json:"following"`
	Notifications             bool           `json:"notifications"`
	CreatedAt                 string         `json:"created_at"`
	UtcOffset                 int            `json:"utc_offset"`
	ProfileBackgroundColor    string         `json:"profile_background_color"`
	ProfileTextColor          string         `json:"profile_text_color"`
	ProfileLinkColor          string         `json:"profile_link_color"`
	ProfileSidebarFillColor   string         `json:"profile_sidebar_fill_color"`
	ProfileSidebarBorderColor string         `json:"profile_sidebar_border_color"`
	ProfileBackgroundImageURL string         `json:"profile_background_image_url"`
	ProfileBackgroundTile     bool           `json:"profile_background_tile"`
	Status                    struct {
		CreatedAt           string `json:"created_at"`
		ID                  string `json:"id"`
		Text                string `json:"text"`
		Source              string `json:"source"`
		Truncated           bool   `json:"truncated"`
		InReplyToLastmsgID  string `json:"in_reply_to_lastmsg_id"`
		InReplyToUserID     string `json:"in_reply_to_user_id"`
		Favorited           bool   `json:"favorited"`
		InReplyToScreenName string `json:"in_reply_to_screen_name"`
	} `json:"status"`
}

type responseStatus struct {
	CreatedAt           string               `json:"created_at"`
	ID                  string               `json:"id"`
	Rawid               int                  `json:"rawid"`
	Text                string               `json:"text"`
	Source              string               `json:"source"`
	Truncated           bool                 `json:"truncated"`
	InReplyToStatusID   string               `json:"in_reply_to_status_id"`
	InReplyToUserID     string               `json:"in_reply_to_user_id"`
	InReplyToScreenName string               `json:"in_reply_to_screen_name"`
	RepostStatusID      string               `json:"repost_status_id"`
	RepostStatus        responseRepostStatus `json:"repost_status"`
	RepostUserID        string               `json:"repost_user_id"`
	RepostScreenName    string               `json:"repost_screen_name"`
	Favorited           bool                 `json:"favorited"`
	User                responseUser         `json:"user"`
	Photo               struct {
		Imageurl string `json:"imageurl"`
		Thumburl string `json:"thumburl"`
		Largeurl string `json:"largeurl"`
	} `json:"photo"`
}

type responseRepostStatus struct {
	ID                  string       `json:"id"`
	CreatedAt           string       `json:"created_at"`
	InReplyToScreenName string       `json:"in_reply_to_screen_name"`
	Rawid               int          `json:"rawid"`
	InReplyToStatusID   string       `json:"in_reply_to_status_id"`
	Location            string       `json:"location"`
	InReplyToUserID     string       `json:"in_reply_to_user_id"`
	Text                string       `json:"text"`
	Truncated           bool         `json:"truncated"`
	Favorited           bool         `json:"favorited"`
	IsSelf              bool         `json:"is_self"`
	User                responseUser `json:"user"`
}

type responseRateLimitStatus struct {
	ResetTime          string `json:"reset_time"`
	RemainingHits      int    `json:"remaining_hits"`
	HourlyLimit        int    `json:"hourly_limit"`
	ResetTimeInSeconds int    `json:"reset_time_in_seconds"`
}

type responseAccountNotification struct {
	Mentions       int `json:"mentions"`
	DirectMessages int `json:"direct_messages"`
	FriendRequests int `json:"friend_requests"`
}

type responseNotifyNum struct {
	Result    string `json:"result"`
	NotifyNum int    `json:"notify_num"`
}

type responseSavedSearch struct {
	ID        int    `json:"id"`
	Query     string `json:"query"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

type responseTrend struct {
	Name  string `json:"name"`
	Query string `json:"query"`
	URL   string `json:"url"`
}

type responseTrends struct {
	AsOf   string `json:"as_of"`
	Trends []responseTrend
}

type responseFriendship struct {
	Relationship struct {
		Source struct {
			ID                   string `json:"id"`
			ScreenName           string `json:"screen_name"`
			Following            string `json:"following"`
			FollowedBy           string `json:"followed_by"`
			NotificationsEnabled string `json:"notifications_enabled"`
			Blocking             string `json:"blocking"`
		} `json:"source"`
		Target struct {
			ID         string `json:"id"`
			ScreenName string `json:"screen_name"`
			Following  string `json:"following"`
			FollowedBy string `json:"followed_by"`
		} `json:"target"`
	} `json:"relationship"`
}

type responseDirectMessage struct {
	ID                  string `json:"id"`
	Text                string `json:"text"`
	SenderID            string `json:"sender_id"`
	RecipientID         string `json:"recipient_id"`
	CreatedAt           string `json:"created_at"`
	SenderScreenName    string `json:"sender_screen_name"`
	RecipientScreenName string `json:"recipient_screen_name"`
	Sender              struct {
		ID                        string `json:"id"`
		Name                      string `json:"name"`
		ScreenName                string `json:"screen_name"`
		Location                  string `json:"location"`
		Gender                    string `json:"gender"`
		Birthday                  string `json:"birthday"`
		Description               string `json:"description"`
		ProfileImageURL           string `json:"profile_image_url"`
		ProfileImageURLLarge      string `json:"profile_image_url_large"`
		URL                       string `json:"url"`
		Protected                 bool   `json:"protected"`
		FollowersCount            int    `json:"followers_count"`
		FriendsCount              int    `json:"friends_count"`
		FavouritesCount           int    `json:"favourites_count"`
		StatusesCount             int    `json:"statuses_count"`
		Following                 bool   `json:"following"`
		Notifications             bool   `json:"notifications"`
		CreatedAt                 string `json:"created_at"`
		UtcOffset                 int    `json:"utc_offset"`
		ProfileBackgroundColor    string `json:"profile_background_color"`
		ProfileTextColor          string `json:"profile_text_color"`
		ProfileLinkColor          string `json:"profile_link_color"`
		ProfileSidebarFillColor   string `json:"profile_sidebar_fill_color"`
		ProfileSidebarBorderColor string `json:"profile_sidebar_border_color"`
		ProfileBackgroundImageURL string `json:"profile_background_image_url"`
		ProfileBackgroundTile     bool   `json:"profile_background_tile"`
	} `json:"sender"`
	Recipient struct {
		ID                        string `json:"id"`
		Name                      string `json:"name"`
		ScreenName                string `json:"screen_name"`
		Location                  string `json:"location"`
		Gender                    string `json:"gender"`
		Birthday                  string `json:"birthday"`
		Description               string `json:"description"`
		ProfileImageURL           string `json:"profile_image_url"`
		ProfileImageURLLarge      string `json:"profile_image_url_large"`
		URL                       string `json:"url"`
		Protected                 bool   `json:"protected"`
		FollowersCount            int    `json:"followers_count"`
		FriendsCount              int    `json:"friends_count"`
		FavouritesCount           int    `json:"favourites_count"`
		StatusesCount             int    `json:"statuses_count"`
		Following                 bool   `json:"following"`
		Notifications             bool   `json:"notifications"`
		CreatedAt                 string `json:"created_at"`
		UtcOffset                 int    `json:"utc_offset"`
		ProfileBackgroundColor    string `json:"profile_background_color"`
		ProfileTextColor          string `json:"profile_text_color"`
		ProfileLinkColor          string `json:"profile_link_color"`
		ProfileSidebarFillColor   string `json:"profile_sidebar_fill_color"`
		ProfileSidebarBorderColor string `json:"profile_sidebar_border_color"`
		ProfileBackgroundImageURL string `json:"profile_background_image_url"`
		ProfileBackgroundTile     bool   `json:"profile_background_tile"`
	} `json:"recipient"`
	InReplyTo struct {
		ID                  string `json:"id"`
		Text                string `json:"text"`
		SenderID            string `json:"sender_id"`
		RecipientID         string `json:"recipient_id"`
		CreatedAt           string `json:"created_at"`
		RecipientScreenName string `json:"recipient_screen_name"`
		SenderScreenName    string `json:"sender_screen_name"`
		Recipient           struct {
			ID                        string `json:"id"`
			Name                      string `json:"name"`
			ScreenName                string `json:"screen_name"`
			Location                  string `json:"location"`
			Gender                    string `json:"gender"`
			Birthday                  string `json:"birthday"`
			Description               string `json:"description"`
			ProfileImageURL           string `json:"profile_image_url"`
			ProfileImageURLLarge      string `json:"profile_image_url_large"`
			URL                       string `json:"url"`
			Protected                 bool   `json:"protected"`
			FollowersCount            int    `json:"followers_count"`
			FriendsCount              int    `json:"friends_count"`
			FavouritesCount           int    `json:"favourites_count"`
			StatusesCount             int    `json:"statuses_count"`
			Following                 bool   `json:"following"`
			Notifications             bool   `json:"notifications"`
			CreatedAt                 string `json:"created_at"`
			UtcOffset                 int    `json:"utc_offset"`
			ProfileBackgroundColor    string `json:"profile_background_color"`
			ProfileTextColor          string `json:"profile_text_color"`
			ProfileLinkColor          string `json:"profile_link_color"`
			ProfileSidebarFillColor   string `json:"profile_sidebar_fill_color"`
			ProfileSidebarBorderColor string `json:"profile_sidebar_border_color"`
			ProfileBackgroundImageURL string `json:"profile_background_image_url"`
			ProfileBackgroundTile     bool   `json:"profile_background_tile"`
		} `json:"recipient"`
		Sender struct {
			ID                        string `json:"id"`
			Name                      string `json:"name"`
			ScreenName                string `json:"screen_name"`
			Location                  string `json:"location"`
			Gender                    string `json:"gender"`
			Birthday                  string `json:"birthday"`
			Description               string `json:"description"`
			ProfileImageURL           string `json:"profile_image_url"`
			ProfileImageURLLarge      string `json:"profile_image_url_large"`
			URL                       string `json:"url"`
			Protected                 bool   `json:"protected"`
			FollowersCount            int    `json:"followers_count"`
			FriendsCount              int    `json:"friends_count"`
			FavouritesCount           int    `json:"favourites_count"`
			StatusesCount             int    `json:"statuses_count"`
			Following                 bool   `json:"following"`
			Notifications             bool   `json:"notifications"`
			CreatedAt                 string `json:"created_at"`
			UtcOffset                 int    `json:"utc_offset"`
			ProfileBackgroundColor    string `json:"profile_background_color"`
			ProfileTextColor          string `json:"profile_text_color"`
			ProfileLinkColor          string `json:"profile_link_color"`
			ProfileSidebarFillColor   string `json:"profile_sidebar_fill_color"`
			ProfileSidebarBorderColor string `json:"profile_sidebar_border_color"`
			ProfileBackgroundImageURL string `json:"profile_background_image_url"`
			ProfileBackgroundTile     bool   `json:"profile_background_tile"`
		} `json:"sender"`
	} `json:"in_reply_to"`
}

type responseDirectMessageConversationItem struct {
	DM      responseDirectMessage `json:"dm"`
	OtherID responseUserID        `json:"otherid"`
	MsgNum  int                   `json:"msg_num"`
	NewConv bool                  `json:"new_conv"`
}
