package gbfsspec

type (
	RentalURIs struct {
		// URI that can be passed to an Android app with an android.intent.action.VIEW Android intent to support
		// Android Deep Links
		Android string `json:"android,omitempty"`

		// URI that can be used on iOS to launch the rental app for this station. More information on this
		// iOS feature can be found here. Please use iOS Universal Links
		IOS string `json:"ios,omitempty"`

		// URL that can be used by a web browser to show more information about renting a vehicle at this station.
		Web string `json:"web,omitempty"`
	}
)
