package gbfsspec

import "testing"

func TestFeedSystemRegions_FeedKey(t *testing.T) {
	var f FeedSystemRegions

	if k := f.FeedKey(); k != FeedKeySystemRegions {
		t.Errorf("expect '%s' got '%s'", FeedKeySystemRegions, k)
	}
}
