package gbfsspec

type (
	FeedFreeBikeStatus struct {
		Metadata

		Data FreeBikeStatusData `json:"data"`
	}

	FreeBikeStatusData struct {
		// Array that contains one object per bike that is currently stopped as defined below.
		Bikes []FreeBikeStatus `json:"bikes"`
	}

	FreeBikeStatus struct {
		// Identifier of a bike, rotated to a random string, at minimum, after each trip to protect privacy.
		BikeID string `json:"bike_id"`

		// Latitude of the bike.
		Latitude float64 `json:"lat"`

		// Longitude of the bike.
		Longitude float64 `json:"lon"`

		// Is the bike currently reserved?
		IsReserved Boolean `json:"is_reserved"`

		// Is the bike currently disabled (broken)?
		IsDisabled Boolean `json:"is_disabled"`

		// Object that contains rental URIs for Android, iOS, and web in the android, ios, and web fields
		RentalURIs RentalURIs `json:"rental_uris,omitempty"`
	}
)

func (_ FeedFreeBikeStatus) FeedKey() string {
	return FeedKeyFreeBikeStatus
}
