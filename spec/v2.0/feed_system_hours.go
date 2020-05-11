package gbfsspec

type (
	FeedSystemHours struct {
		Metadata

		Data SystemHoursData `json:"data"`
	}

	SystemHoursData struct {
		// Array of objects as defined below. The array must contain a minimum of one object identifying hours
		// for every day of the week or a maximum of two for each day of the week objects (one for each user type).
		RentalHours []SystemHoursRentalHours `json:"rental_hours"`
	}

	SystemHoursRentalHours struct {
		// An array of member and/or nonmember value(s).
		// This indicates that this set of rental hours applies to either members or non-members only.
		UserTypes []UserType `json:"user_types"`

		// An array of abbreviations (first 3 letters) of English
		// names of the days of the week for which this object applies
		Days []Day `json:"days"`

		// Start time for the hours of operation of the system in the time zone indicated in system_information.
		StartTime Time `json:"start_time"`

		// End time for the hours of operation of the system in the time zone indicated in system_information.
		EndTime Time `json:"end_time"`
	}
)

func (_ FeedSystemHours) FeedKey() string {
	return FeedKeySystemHours
}
