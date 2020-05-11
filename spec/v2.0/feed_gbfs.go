package gbfsspec

type (
	FeedGBFS struct {
		Metadata

		Data GBFSData `json:"data"`
	}

	GBFSData struct {
		// The language that will be used throughout the rest of the files.
		// It must match the value in the system_information.json file.
		Languages map[string]GBFSLanguage `json:"languages"`
	}

	GBFSLanguage struct {
		// An array of all of the feeds that are published by this auto-discovery file.
		// Each element in the array is an object with the keys below.
		Feeds []GBFSFeed `json:"feeds"`
	}

	GBFSFeed struct {
		// Key identifying the type of feed this is. The key must be the base file name defined in the spec for
		// the corresponding feed type (system_information for system_information.json file,
		// station_information for station_information.json file).
		Name string `json:"name"`

		// URL for the feed. Note that the actual feed endpoints (urls) may not be defined in the file_name.json format.
		// For example, a valid feed endpoint could end with station_info instead of station_information.json.
		URL string `json:"url"`
	}
)

func (_ FeedGBFS) FeedKey() string {
	return FeedKeyAutoDiscovery
}
