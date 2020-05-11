package gbfsspec

import "testing"

func TestFeedFreeBikeStatus_FeedKey(t *testing.T) {
	var f FeedFreeBikeStatus

	if k := f.FeedKey(); k != FeedKeyFreeBikeStatus {
		t.Errorf("expect '%s' got '%s'", FeedKeyFreeBikeStatus, k)
	}
}
