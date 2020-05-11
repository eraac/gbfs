package gbfsspec

type (
	FeedSystemCalendars struct {
		Metadata

		Data SystemCalendarsData `json:"data"`
	}

	SystemCalendarsData struct {
		// Array of objects describing the system operational calendar. A minimum of one calendar object is required.
		// If start and end dates are the same every year, then start_year and end_year should be omitted.
		Calendars SystemCalendar `json:"calendars"`
	}

	SystemCalendar struct {
		// Starting date for the system operations (1-31).
		StartDay int `json:"start_day"`

		// Starting month for the system operations (1-12).
		StartMonth int `json:"start_month"`

		// Starting year for the system operations.
		StartYear int `json:"start_year,omitempty"`

		// Ending date for the system operations (1-31).
		EndDay int `json:"end_day"`

		// Ending month for the system operations (1-12).
		EndMonth int `json:"end_month"`

		// Ending year for the system operations.
		EndYear int `json:"end_year,omitempty"`
	}
)

func (_ FeedSystemCalendars) FeedKey() string {
	return FeedKeySystemCalendar
}
