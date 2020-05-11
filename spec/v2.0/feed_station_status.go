package gbfsspec

type (
	FeedStationStatus struct {
		Metadata

		Data StationStatusData `json:"data"`
	}

	StationStatusData struct {
		// Array that contains one object per station in the system
		Stations []StationStatus `json:"stations"`
	}

	StationStatus struct {
		// Identifier of a station see station_information.
		StationID string `json:"station_id"`

		// Number of bikes available for rental. Number of functional bikes physically at the station.
		// To know if the bikes are available for rental.
		NumBikesAvailable int `json:"num_bikes_available"`

		// Number of disabled bikes at the station.
		NumBikesDisabled int `json:"num_bikes_disabled,omitempty"`

		// Required except for stations that have unlimited docking capacity (e.g. virtual stations).
		// Number of functional docks physically at the station.
		NumDocksAvailable int `json:"num_docks_available,omitempty"`

		// Number of empty but disabled dock points at the station.
		NumDocksDisabled int `json:"num_docks_disabled,omitempty"`

		// Is the station currently on the street?
		IsInstalled Boolean `json:"is_installed"`

		// Is the station currently renting bikes?
		IsRenting Boolean `json:"is_renting"`

		// Is the station accepting bike returns?
		IsReturning Boolean `json:"is_returning"`

		// The last time this station reported its status.
		LastReported Timestamp `json:"last_reported"`
	}
)

func (_ FeedStationStatus) FeedKey() string {
	return FeedKeyStationStatus
}
