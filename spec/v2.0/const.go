package gbfsspec

const Version = "2.0"

const (
	FeedKeyAutoDiscovery      = "gbfs"
	FeedKeyGBFSVersions       = "gbfs_versions"
	FeedKeySystemInformation  = "system_information"
	FeedKeyStationInformation = "station_information"
	FeedKeyStationStatus      = "station_status"
	FeedKeyFreeBikeStatus     = "free_bike_status"
	FeedKeySystemHours        = "system_hours"
	FeedKeySystemCalendar     = "system_calendar"
	FeedKeySystemRegions      = "system_regions"
	FeedKeySystemPricingPlans = "system_pricing_plans"
	FeedKeySystemAlerts       = "system_alerts"
)

const (
	RentalMethodKey           RentalMethod = "KEY"
	RentalMethodCreditCard    RentalMethod = "CREDITCARD"
	RentalMethodPayPass       RentalMethod = "PAYPASS"
	RentalMethodApplePay      RentalMethod = "APPLEPAY"
	RentalMethodAndroidPay    RentalMethod = "ANDROIDPAY"
	RentalMethodTransitCard   RentalMethod = "TRANSITCARD"
	RentalMethodAccountNumber RentalMethod = "ACCOUNTNUMBER"
	RentalMethodPhone         RentalMethod = "PHONE"
)

const (
	UserTypeMember    UserType = "member"
	UserTypeNonMember UserType = "nonmember"
)

const (
	DayMonday    Day = "mon"
	DayTuesday   Day = "tue"
	DayWednesday Day = "wed"
	DayThursday  Day = "thu"
	DayFriday    Day = "fri"
	DaySaturday  Day = "sat"
	DaySunday    Day = "sun"
)

const (
	AlertTypeSystemClosure  AlertType = "SYSTEM_CLOSURE"
	AlertTypeStationClosure AlertType = "STATION_CLOSURE"
	AlertTypeStationMove    AlertType = "STATION_MOVE"
	AlertTypeOther          AlertType = "OTHER"
)
