package gbfs

import (
	"encoding/json"
	"time"
)

type (
	// JSON common to all feeds
	JSON struct {
		LastUpdated int64           `json:"last_updated"`
		TTL         int             `json:"ttl"`
		Data        json.RawMessage `json:"data"`
	}

	// GBFS auto-discovery file that links to all of the other files published by the system
	// Optional
	GBFS struct {
		// JSON contains JSON about the JSON file (last_update, ttl and raw value)
		JSON *JSON `json:"-"`

		// Languages that all of the contained files will be published in.
		// This language must match the value in the system_information file
		// Required
		Languages map[string]Feeds `json:"data"`
	}

	// Feeds is an array contains all feeds for a system and language
	Feeds struct {
		// Feeds list for one language and one system
		Feeds []Feed `json:"feeds"`
	}

	// Feed that are published by this auto-discovery file
	Feed struct {
		// Name identifying the type of feed this is
		// Required
		Name FeedKey `json:"name"`

		// URL for the feed
		// Required
		URL string `json:"url"`
	}

	// SystemInformation including System operator, System location, year implemented, URLs, contact info, time zone
	// Required
	SystemInformation struct {
		// JSON contains JSON about the JSON file (last_update, ttl and raw value)
		JSON *JSON `json:"-"`

		// SystemID is an identifier for this bike share system.
		// This should be globally unique (even between different systems) and it is currently up to the publisher of the feed to guarantee uniqueness.
		// In addition, this value is intended to remain the same over the life of the system
		// Required
		SystemID string `json:"system_id"`

		// An IETF language tag indicating the language that will be used throughout the rest of the files.
		// This is a string that defines a single language tag only
		// Required
		Language string `json:"language"`

		// Name of the system to be displayed to customers
		// Required
		Name string `json:"name"`

		// ShortName is an abbreviation for a system
		// Optional
		ShortName string `json:"short_name"`

		// URL of the bike share system.
		// The value must be a fully qualified URL that includes http:// or https://, and any special characters in the URL must be correctly escaped
		// Optional
		URL string `json:"url"`

		// PurchaseURL where a customer can purchase a membership or learn more about purchasing memberships
		// Optional
		PurchaseURL string `json:"purchase_url"`

		// StartDate (format YYYY-MM-DD) representing the date that the system began operations
		// Optional
		StartDate string `json:"start_date"`

		// PhoneNumber is a single voice telephone number for the specified system.
		// This field is a string value that presents the telephone number as typical for the system's service area.
		// It can and should contain punctuation marks to group the digits of the number.
		// Dialable text (for example, Capital Bikeshare’s "877-430-BIKE") is permitted, but the field must not contain any other descriptive text
		// Optional
		PhoneNumber string `json:"phone_number"`

		// Email address for customers to address questions about the system
		// Optional
		Email string `json:"email"`

		// Timezone where the system is located. Time zone names never contain the space character but may contain an underscore.
		// Please refer to the "TZ" value in https://en.wikipedia.org/wiki/List_of_tz_database_time_zones for a list of valid values
		// Required
		Timezone string `json:"timezone"`

		// URL of a page that defines the license terms for the GBFS data for this system,
		// as well as any other license terms the system would like to define (including the use of corporate trademarks, etc)
		// Optional
		LicenseURL string `json:"license_url"`
	}

	// StationsInformation like their capacities and locations
	// Required of systems utilizing docks
	StationsInformation struct {
		// JSON contains JSON about the JSON file (last_update, ttl and raw value)
		JSON *JSON `json:"-"`

		// Stations is an array that contains one object per station in the system as defined below
		// Required
		Stations []StationInformation `json:"stations"`
	}

	// StationInformation in the system
	StationInformation struct {
		// StationID is a unique identifier of a station.
		// Required
		StationID string `json:"station_id"`

		// Name of the station
		// Required
		Name string `json:"name"`

		// ShortName or other type of identifier, as used by the data publisher
		// Optional
		ShortName string `json:"short_name"`

		// Latitude of station
		// Required
		Latitude float64 `json:"lat"`

		// Longitude of station
		// Required
		Longitude float64 `json:"lon"`

		// Address where station is located
		// Optional
		Address string `json:"address"`

		// CrossStreet of where the station is located.
		// This field is intended to be a descriptive field for human consumption.
		// In cities, this would be a cross street, but could also be a description of a location in a park, etc.
		// Optional
		CrossStreet string `json:"cross_street"`

		// TODO specification say "string", provider return an int :shrug: - maybe find another lib for json, can fallback to string for an int
		// RegionID where station is located (see system_regions.json)
		// Optional
		RegionID int `json:"region_id"`

		// PostalCode where station is located
		// Optional
		PostalCode string `json:"post_code"`

		// RentalMethods containing the payment methods accepted at this station
		// Optional
		RentalMethods []RentalMethod `json:"rental_methods"`

		// Capacity is number of total docking points installed at this station, both available and unavailable
		// Optional
		Capacity int `json:"capacity"`
	}

	// StationsStatus like number of available bikes and docks at each station and station availability.
	// Required of systems utilizing docks
	StationsStatus struct {
		// JSON contains JSON about the JSON file (last_update, ttl and raw value)
		JSON *JSON `json:"-"`

		// Stations array that contains one object per station in the system
		// Required
		Stations []StationStatus `json:"stations"`
	}

	// StationStatus in the system
	StationStatus struct {
		// StationID is a unique identifier of a station
		// Required
		StationID string `json:"station_id"`

		// NumBikesAvailable for rental
		// Required
		NumBikesAvailable int `json:"num_bikes_available"`

		// NumBikesDisabled at the station. Vendors who do not want to publicize the number of disabled bikes or docks
		// in their system can opt to omit station capacity (in station_information), num_bikes_disabled and num_docks_disabled.
		// If station capacity is published then broken docks/bikes can be inferred
		// (though not specifically whether the decreased capacity is a broken bike or dock)
		// Optional
		NumBikesDisabled int `json:"num_bikes_disabled"`

		// NumDocksAvailable accepting bike returns
		// Required
		NumDocksAvailable int `json:"num_docks_available"`

		// NumDocksDisabled empty but disabled dock points at the station
		// Optional
		NumDocksDisabled int `json:"num_docks_disabled"`

		// IsInstalled station currently on the street
		// Required
		IsInstalled bool `json:"is_installed"`

		// IsRenting bikes currently bikes
		// Required
		IsRenting bool `json:"is_renting"`

		// LastReported POSIX timestamp indicating the last time this station reported its status to the backend
		// Required
		LastReported int `json:"last_reported"`
	}

	// FreeBikeStatus describes bikes that are available for rent.
	// Required of systems that don't utilize docks or offer bikes for rent outside of stations
	FreeBikeStatus struct {
		// JSON contains JSON about the JSON file (last_update, ttl and raw value)
		JSON *JSON `json:"-"`

		// Bikes array that contains one object per bike that is currently docked/stopped outside of the system
		Bikes []BikeStatus `json:"bikes"`
	}

	// BikeStatus that is currently docked/stopped outside of the system
	BikeStatus struct {
		// BikeID is a unique identifier of a bike
		// Required
		BikeID string `json:"bike_id"`

		// Latitude of the bike
		// Required
		Latitude float64 `json:"lat"`

		// Longitude of the bike
		// Required
		Longitude float64 `json:"lon"`

		// IsReserved bike for someone else
		// Required
		IsReserved bool `json:"is_reserved"`

		// IsDisabled bike (broken)
		// Required
		IsDisabled bool `json:"is_disabled"`
	}

	// SystemHours describes the hours of operation for the system
	// Optional
	SystemHours struct {
		// JSON contains JSON about the JSON file (last_update, ttl and raw value)
		JSON *JSON `json:"-"`

		// RentalHours is an array of hour objects as defined below.
		// Can contain a minimum of one object identifying hours for all days of the week
		// or a maximum of fourteen hour objects are allowed (one for each day of the week for each "member" or "nonmember" user type)
		// Required
		RentalHours []RentalHours `json:"rental_hours"`
	}

	// RentalHours information
	RentalHours struct {
		// An array of "member" and "nonmember" values.
		// This indicates that this set of rental hours applies to either members or non-members only.
		// Required
		UserTypes []UserType `json:"user_types"`

		// An array of abbreviations (first 3 letters) of English names of the days of the week that this hour object applies to (i.e. ["mon", "tue"]).
		// Each day can only appear once within all of the hours objects in this feed.
		// Required
		Days []Day `json:"days"`

		// StartTime for the hours of operation of the system in the time zone indicated in system_information.json (00:00:00 - 23:59:59)
		// Required
		StartTime string `json:"start_time"`

		// End time for the hours of operation of the system in the time zone indicated in system_information.json (00:00:00 - 47:59:59).
		// Time can stretch up to one additional day in the future to accommodate situations where,
		// for example, a system was open from 11:30pm - 11pm the next day (i.e. 23:30-47:00)
		// Required
		EndTime string `json:"end_time"`
	}

	// SystemCalendar describes the days of operation for the system
	// Optional
	SystemCalendar struct {
		// JSON contains JSON about the JSON file (last_update, ttl and raw value)
		JSON *JSON `json:"-"`

		// Calendars is an array of year objects describing the system operational calendar.
		// A minimum of one calendar object is required, which could indicate a general calendar,
		// or multiple calendars could be present indicating arbitrary start and end dates
		// Required
		Calendars []Calendar `json:"calendars"`
	}

	// Calendar describing the system operational
	Calendar struct {
		// StartDay for the system operations (1-31)
		// Required
		StartDay int `json:"start_day"`

		// StartMonth for the system operations (1-12)
		// Required
		StartMonth int `json:"start_month"`

		// StartYear for the system operations
		// Optional
		StartYear int `json:"start_year"`

		// EndDay for the system operations (1-12)
		// Required
		EndDay int `json:"end_day"`

		// EndMonth for the system operations (1-12)
		// Required
		EndMonth int `json:"end_month"`

		// EndYear for the system operations
		// Optional
		EndYear int `json:"end_year"`
	}

	// SystemRegions describes list of regions the system is broken up by geographic or political region
	// Optional
	SystemRegions struct {
		// JSON contains JSON about the JSON file (last_update, ttl and raw value)
		JSON *JSON `json:"-"`

		// Regions is an array of region objects
		Regions []Region `json:"regions"`
	}

	// Region describes by geographic or political region
	Region struct {
		// RegionID is a unique identifier for this region
		// Required
		RegionID string `json:"region_id"`

		// Name for this region
		// Required
		Name string `json:"name"`
	}

	// SystemPricingPlans describes the system pricing
	// Optional
	SystemPricingPlans struct {
		// JSON contains JSON about the JSON file (last_update, ttl and raw value)
		JSON *JSON `json:"-"`

		// Plans is an array of any number of plan
		Plans []Plan `json:"plans"`
	}

	// Plan available for a system
	Plan struct {
		// PlanID is a unique identifier for this plan in the system
		// Required
		PlanID string `json:"plan_id"`

		// URL where the customer can learn more about this particular scheme
		// Optional
		URL string `json:"url"`

		// Name of this pricing scheme
		// Required
		Name string `json:"name"`

		// Currency this pricing is in (ISO 4217)
		// Required
		Currency string `json:"currency"`

		// Price for this plan
		// Required
		Price float64 `json:"price"`

		// IsTaxable false indicates that no additional tax will be added (either because tax is not charged, or because it is included)
		// Required
		IsTaxable bool `json:"is_taxable"`

		// Description this particular pricing plan in human readable terms
		// Required
		Description string `json:"description"`
	}

	// SystemAlerts describes current system alerts
	// This feed is intended to inform customers about changes to the system that do not fall
	// within the normal system operations. For example, system closures due to weather would
	// be listed here, but a system that only operated for part of the year would have that
	// schedule listed in the system_calendar.json feed.
	// Optional
	SystemAlerts struct {
		// JSON contains JSON about the JSON file (last_update, ttl and raw value)
		JSON *JSON `json:"-"`

		// Alerts array contains objects each indicating a separate system alert
		// Required
		Alerts []Alert `json:"alerts"`
	}

	// Alert is to inform customers about changes to the system that do not fall within the normal system operations
	Alert struct {
		// AlertID unique identifier for this alert
		// Required
		AlertID string `json:"alert_id"`

		// Type of the alert
		// Required
		Type string `json:"type"`

		// TODO check the type
		// Times is array of hashes with the keys "start" and "end" indicating when the alert is in effect
		// Optional
		Times map[string]AlertTime `json:"times"`

		// StationIDs is a list of stations affected by this alert
		// Optional
		StationIDs []string `json:"station_ids"`

		// RegionIDs is a list of region affected by this alert
		// Optional
		RegionIDs []string `json:"region_ids"`

		// URL where the customer can learn more information about this alert
		// Optional
		URL string `json:"url"`

		// Summary of this alert to be displayed to the customer
		// Optional
		Summary string `json:"summary"`

		// Description give more detail about the alert
		// Optional
		Description string `json:"description"`

		// LastUpdated POSIX timestamp indicating the last time the info for the particular alert was updated
		// Optional
		LastUpdated int `json:"last_updated"`
	}

	// AlertTime indicating when the alert is in effect
	AlertTime struct {
		// Start POSIX timestamp
		// Required
		Start int `json:"start"`

		// End POSIX timestamp
		// Optional
		End int `json:"end"`
	}
)

func (j *JSON) IsOutdated() bool {
	n := time.Now().Unix()

	return j.LastUpdated + int64(j.TTL) < n
}
