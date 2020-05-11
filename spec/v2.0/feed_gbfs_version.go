package gbfsspec

type (
	FeedGBFSVersions struct {
		Metadata

		Data GBFSVersionData `json:"data"`
	}

	GBFSVersionData struct {
		// Contains one object, as defined below, for each of the available versions of a feed.
		// The array must be sorted by increasing MAJOR and MINOR version number.
		Versions []GBFSVersion `json:"versions"`
	}

	GBFSVersion struct {
		// The semantic version of the feed in the form X.Y.
		Version string `json:"version"`

		// URL of the corresponding gbfs.json endpoint.
		URL string `json:"url"`
	}
)

func (_ FeedGBFSVersions) FeedKey() string {
	return FeedKeyGBFSVersions
}
