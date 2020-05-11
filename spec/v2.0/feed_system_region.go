package gbfsspec

type (
	FeedSystemRegions struct {
		Metadata

		Data SystemRegionsData `json:"data"`
	}

	SystemRegionsData struct {
		Regions []SystemRegion `json:"regions"`
	}

	SystemRegion struct {
		// Identifier for the region.
		RegionID string `json:"region_id"`

		// Public name for this region.
		Name string `json:"name"`
	}
)

func (_ FeedSystemRegions) FeedKey() string {
	return FeedKeySystemRegions
}
