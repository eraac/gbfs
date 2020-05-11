package gbfsspec

import "testing"

func TestFeedSystemHours_FeedKey(t *testing.T) {
	var f FeedSystemHours

	if k := f.FeedKey(); k != FeedKeySystemHours {
		t.Errorf("expect '%s' got '%s'", FeedKeySystemHours, k)
	}
}
