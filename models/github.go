package models

//Gists represents Gist data model
type Gists struct {
	URL        string `json:"url"`
	ID         string `json:"id"`
	Owner      Owner  `json:"owner"`
	ExteraInfo string `json:"extera_info"`
}

//File represents File data model
type File struct {
	Filename string      `json:"filename"`
	Type     string      `json:"type"`
	Language interface{} `json:"language"`
	RawURL   string      `json:"raw_url"`
	Size     int         `json:"size"`
}

//Owner represents Owner data model
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
