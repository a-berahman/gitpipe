package github

import "time"

// GistsForUserRS implemented from GitHub List gists for a user schema
// https://docs.github.com/en/rest/reference/gists#list-gists-for-a-user
type GistsForUserRS struct {
	URL         string                 `json:"url"`
	ForksURL    string                 `json:"forks_url"`
	CommitsURL  string                 `json:"commits_url"`
	ID          string                 `json:"id"`
	NodeID      string                 `json:"node_id"`
	GitPullURL  string                 `json:"git_pull_url"`
	GitPushURL  string                 `json:"git_push_url"`
	HTMLURL     string                 `json:"html_url"`
	Files       map[string]interface{} `json:"files,omitempty"`
	Public      bool                   `json:"public"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
	Description string                 `json:"description"`
	Comments    int                    `json:"comments"`
	User        interface{}            `json:"user"`
	CommentsURL string                 `json:"comments_url"`
	Owner       Owner                  `json:"owner"`
	Truncated   bool                   `json:"truncated"`
}

// Owner implemented from GitHub Owner schema
// https://docs.github.com/en/rest/reference/gists#list-gists-for-a-user
type Owner struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}
