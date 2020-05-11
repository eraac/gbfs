package gbfs

type (
	// Client interface to interact with GBFS provider
	Client interface {
		ForceURLs(map[string]string, bool)
		Get(string, Feed) error
		Refresh(Feed, bool) error
	}

	// HTTPOption for HTTPClient
	HTTPOption func(*HTTPClient)

	// Feed implementation, allow to refresh a feed without the feed name
	Feed interface {
		FeedKey() string
		IsExpired() bool
	}
)
