package gbfsspec

import "time"

type (
	FeedSystemInformation struct {
		Metadata

		Data SystemInformationData `json:"data"`
	}

	SystemInformationData struct {
		// Identifier for this bike share system. This should be globally unique (even between different systems) -
		// for example, bcycle_austin or biketown_pdx. It is up to the publisher of the feed to guarantee uniqueness.
		// This value is intended to remain the same over the life of the system.
		SystemID string `json:"system_id"`

		// The language that will be used throughout the rest of the files.
		// It must match the value in the gbfs.json file.
		Language string `json:"language"`

		// Name of the system to be displayed to customers.
		Name string `json:"name"`

		// Optional abbreviation for a system.
		ShortName string `json:"short_name,omitempty"`

		// Name of the operator.
		Operator string `json:"operator,omitempty"`

		// The URL of the bike share system.
		URL string `json:"url,omitempty"`

		// URL where a customer can purchase a membership.
		PurchaseURL string `json:"purchase_url,omitempty"`

		// Date that the system began operations.
		StartDate Date `json:"start_date,omitempty"`

		// A single voice telephone number for the specified system that presents the telephone number as typical for
		// the system's service area. It can and should contain punctuation marks to group the digits of the number.
		// Dialable text (for example, Capital Bikeshare’s "877-430-BIKE") is permitted,
		// but the field must not contain any other descriptive text.
		PhoneNumber string `json:"phone_number,omitempty"`

		// Email address actively monitored by the operator’s customer service department.
		// This email address should be a direct contact point where riders can reach a customer service representative.
		Email string `json:"email,omitempty"`

		// A single contact email address for consumers of this feed to report technical issues.
		FeedContactEmail string `json:"feed_contact_email,omitempty"`

		// The time zone where the system is located.
		Timezone string `json:"timezone"`

		// A fully qualified URL of a page that defines the license terms for the GBFS data for this system,
		// as well as any other license terms the system would like to define
		// (including the use of corporate trademarks, etc)
		LicenseURL string `json:"license_url,omitempty"`

		// Contains rental app information in the android and ios JSON objects.
		RentalApps SystemInformationRentalApp `json:"rental_apps,omitempty"`
	}

	SystemInformationRentalApp struct {
		// Contains rental app download and app discovery information for the Android platform
		// in the store_uri and discovery_uri fields
		Android SystemInformationRentalAppAndroid `json:"android,omitempty"`

		// Contains rental information for the iOS platform in the store_uri and discovery_uri fields
		IOS SystemInformationRentalAppIOS `json:"ios,omitempty"`
	}

	SystemInformationRentalAppAndroid struct {
		// URI where the rental Android app can be downloaded from. Typically this will be a URI to an app store
		// such as Google Play. If the URI points to an app store such as Google Play, the URI should follow Android
		// best practices so the viewing app can directly open the URI to the native app store app instead of a website.
		StoreURI string `json:"store_uri,omitempty"`

		// URI that can be used to discover if the rental Android app is installed on the device
		// (e.g., using PackageManager.queryIntentActivities()). This intent is used by viewing apps prioritize rental
		// apps for a particular user based on whether they already have a particular rental app installed.
		DiscoveryURI string `json:"discovery_uri,omitempty"`
	}

	SystemInformationRentalAppIOS struct {
		// URI where the rental iOS app can be downloaded from. Typically this will be a URI to an app store
		// such as the Apple App Store. If the URI points to an app store such as the Apple App Store,
		// the URI should follow iOS best practices  so the viewing app can directly open the URI to
		// the native app store app instead of a website.
		StoreURI string `json:"store_uri,omitempty"`

		// URI that can be used to discover if the rental iOS app is installed on the device
		// (e.g., using UIApplication canOpenURL:).
		// This intent is used by viewing apps prioritize rental apps for a particular user based on whether
		// they already have a particular rental app installed.
		DiscoveryURI string `json:"discovery_uri,omitempty"`
	}
)

func (_ FeedSystemInformation) FeedKey() string {
	return FeedKeySystemInformation
}

func (s SystemInformationData) GetStartDate() (time.Time, error) {
	return s.StartDate.ToTime(s.Timezone)
}
