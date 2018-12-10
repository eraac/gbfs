package gbfs

type (
	Client interface {
		Get(key FeedKey) (*JSON, error)
		AutoDiscovery() (*GBFS, error)
		SystemInformation() (*SystemInformation, error)
		StationsInformation() (*StationsInformation, error)
		StationsStatus() (*StationsStatus, error)
		FreeBikeStatus() (*FreeBikeStatus, error)
		SystemHours() (*SystemHours, error)
		SystemCalendar() (*SystemCalendar, error)
		SystemRegions() (*SystemRegions, error)
		SystemPricingPlans() (*SystemPricingPlans, error)
		SystemAlerts() (*SystemAlerts, error)
	}

	HTTPOption func(*HTTPClient)
)
