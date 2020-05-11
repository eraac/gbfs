package gbfsspec

import "testing"

func TestFeedGBFSVersions_FeedKey(t *testing.T) {
	var f FeedGBFSVersions

	if k := f.FeedKey(); k != FeedKeyGBFSVersions {
		t.Errorf("expect '%s' got '%s'", FeedKeyGBFSVersions, k)
	}
}
