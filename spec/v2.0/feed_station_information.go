package gbfsspec

type (
	FeedStationInformation struct {
		Metadata

		Data StationInformationData `json:"data"`
	}

	StationInformationData struct {
		// Array that contains one object per station as defined below.
		Stations []StationInformation `json:"stations"`
	}

	StationInformation struct {
		// Identifier of a station.
		StationID string `json:"station_id"`

		// Public name of the station.
		Name string `json:"name"`

		// Short name or other type of identifier.
		ShortName string `json:"short_name,omitempty"`

		// The latitude of station.
		Latitude float64 `json:"lat"`

		// The longitude of station.
		Longitude float64 `json:"lon"`

		// Address (street number and name) where station is located. This
		Address string `json:"address,omitempty"`

		// Cross street or landmark where the station is located.
		CrossStreet string `json:"cross_street,omitempty"`

		// Identifier of the region where station is located. See SystemRegion
		RegionID string `json:"region_id,omitempty"`

		// Postal code where station is located.
		PostCode string `json:"post_code,omitempty"`

		// Payment methods accepted at this station.
		RentalMethods []RentalMethod `json:"rental_methods,omitempty"`

		// Number of total docking points installed at this station, both available and unavailable.
		Capacity int `json:"capacity,omitempty"`

		// Contains rental URIs for Android, iOS, and web in the android, ios, and web fields.
		RentalURIs RentalURIs `json:"rental_uris,omitempty"`
	}
)

func (_ FeedStationInformation) FeedKey() string {
	return FeedKeyStationInformation
}
