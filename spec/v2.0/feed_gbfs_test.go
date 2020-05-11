package gbfsspec

import "testing"

func TestFeedGBFS_FeedKey(t *testing.T) {
	var f FeedGBFS

	if k := f.FeedKey(); k != FeedKeyAutoDiscovery {
		t.Errorf("expect '%s' got '%s'", FeedKeyAutoDiscovery, k)
	}
}
