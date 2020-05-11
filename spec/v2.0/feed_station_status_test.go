package gbfsspec

import "testing"

func TestFeedStationStatus_FeedKey(t *testing.T) {
	var f FeedStationStatus

	if k := f.FeedKey(); k != FeedKeyStationStatus {
		t.Errorf("expect '%s' got '%s'", FeedKeyStationStatus, k)
	}
}
