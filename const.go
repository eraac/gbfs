package gbfs

const (
	FeedKeyAutoDiscovery      FeedKey = "gbfs"
	FeedKeySystemInformation  FeedKey = "system_information"
	FeedKeyStationInformation FeedKey = "station_information"
	FeedKeyStationStatus      FeedKey = "station_status"
	FeedKeyFreeBikeStatus     FeedKey = "free_bike_status"
	FeedKeySystemHours        FeedKey = "system_hours"
	FeedKeySystemCalendar     FeedKey = "system_calendar"
	FeedKeySystemRegions      FeedKey = "system_regions"
	FeedKeySystemPricingPlans FeedKey = "system_pricing_plans"
	FeedKeySystemAlerts       FeedKey = "system_alerts"
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

type (
	FeedKey      string
	RentalMethod string
	UserType     string
	Day          string
	AlertType    string
)
