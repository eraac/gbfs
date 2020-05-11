package gbfsspec

type (
	FeedSystemAlerts struct {
		Metadata

		Data SystemAlertsData `json:"data"`
	}

	SystemAlertsData struct {
		Alerts []SystemAlert `json:"alerts"`
	}

	SystemAlert struct {
		// Identifier for this alert.
		AlertID string `json:"alert_id"`

		// Valid values are: SYSTEM_CLOSURE, STATION_CLOSURE, STATION_MOVE, OTHER
		Type AlertType `json:"type"`

		// Array of objects with the fields start and end indicating when the alert is in effect
		// (e.g. when the system or station is actually closed, or when it is scheduled to be moved).
		Times []SystemAlertTime `json:"times,omitempty"`

		// If this is an alert that affects one or more stations, include their ID(s).
		// Otherwise omit this field. If both station_id and region_id are omitted,
		// this alert affects the entire system.
		StationIDs []string `json:"station_ids,omitempty"`

		// If this system has regions, and if this alert only affects certain regions,
		// include their ID(s). Otherwise, omit this field. If both station_ids and region_ids are omitted,
		// this alert affects the entire system.
		RegionIDs []string `json:"region_ids,omitempty"`

		// URL where the customer can learn more information about this alert.
		URL string `json:"url,omitempty"`

		// A short summary of this alert to be displayed to the customer.
		Summary string `json:"summary"`

		// Detailed description of the alert.
		Description string `json:"description,omitempty"`

		// Indicates the last time the info for the alert was updated.
		LastUpdated Timestamp `json:"last_updated,omitempty"`
	}

	SystemAlertTime struct {
		// Start time of the alert.
		Start Timestamp `json:"start"`

		// End time of the alert. If there is currently no end time planned for the alert, this can be omitted.
		End Timestamp `json:"end"`
	}
)

func (_ FeedSystemAlerts) FeedKey() string {
	return FeedKeySystemAlerts
}
