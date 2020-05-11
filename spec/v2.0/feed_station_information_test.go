package gbfsspec

import "testing"

func TestFeedStationInformation_FeedKey(t *testing.T) {
	var f FeedStationInformation

	if k := f.FeedKey(); k != FeedKeyStationInformation {
		t.Errorf("expect '%s' got '%s'", FeedKeyStationInformation, k)
	}
}
