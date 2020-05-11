package gbfs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	// HTTPClient use HTTP protocol to fetch feeds
	HTTPClient struct {
		client http.Client

		baseURL  string
		urls     map[string]string
		language string
	}
)

// NewHTTPClient return a gbfs.Client will use http to fetch feeds
func NewHTTPClient(opts ...HTTPOption) (Client, error) {
	c := &HTTPClient{urls: make(map[string]string)}

	for _, opt := range opts {
		opt(c)
	}

	if c.baseURL == "" {
		return nil, ErrBaseURLMissing
	}

	return c, nil
}

// ForceURLs set the full URL for each feed. Where map key is the feed name and value is URL
// Set replace to true when you want to clean previous build URL
// Useful for provider that doesn't respect the standard for the URL
func (c *HTTPClient) ForceURLs(urls map[string]string, replace bool) {
	if replace {
		c.urls = urls
		return
	}

	for k, u := range urls {
		c.urls[k] = u
	}
}

// Get one feed and try to decode the response in 'out' structure
func (c *HTTPClient) Get(key string, out Feed) error {
	if out.FeedKey() != key {
		return ErrInvalidFeed
	}

	req, err := http.NewRequest(http.MethodGet, c.url(key), nil)
	if err != nil {
		return fmt.Errorf("http.NewRequest: %w", err)
	}

	res, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("http.Do: %w", err)
	}
	defer func() { _ = res.Body.Close() }()

	if res.StatusCode != http.StatusOK {
		switch res.StatusCode {
		case http.StatusNotFound:
			return ErrFeedNotExist
		default:
			return fmt.Errorf("invalid status code (%d)", res.StatusCode)
		}
	}

	return json.NewDecoder(res.Body).Decode(out)
}

// Refresh the feed when is expired or forced (via 'forceRefresh')
func (c *HTTPClient) Refresh(f Feed, forceRefresh bool) error {
	// not forced and feed not expired
	if !forceRefresh && !f.IsExpired() {
		return nil
	}

	return c.Get(f.FeedKey(), f)
}

func (c *HTTPClient) url(key string) string {
	if u, ok := c.urls[key]; ok {
		return u
	}

	l := c.language

	if l != "" {
		l = fmt.Sprintf("%s/", l)
	}

	c.urls[key] = fmt.Sprintf("%s/%s%s.json", c.baseURL, l, key)

	return c.urls[key]
}

// ==========
//  OPTIONS
// ==========

// HTTPOptionClient specify and http.Client for gbfs.HTTPClient
func HTTPOptionClient(h http.Client) HTTPOption {
	return func(c *HTTPClient) {
		c.client = h
	}
}

// HTTPOptionBaseURL specify the base URL for the feed
func HTTPOptionBaseURL(url string) HTTPOption {
	return func(c *HTTPClient) {
		c.baseURL = url
	}
}

// HTTPOptionLanguage specify the language of the feed
// Used to determined the path of the feed
func HTTPOptionLanguage(lang string) HTTPOption {
	return func(c *HTTPClient) {
		c.language = lang
	}
}

// HTTPOptionForceURL specify an URL to use for the feed
func HTTPOptionForceURL(key, url string) HTTPOption {
	return func(c *HTTPClient) {
		c.urls[key] = url
	}
}
