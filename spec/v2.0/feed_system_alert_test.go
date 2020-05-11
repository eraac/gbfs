package gbfsspec

import "testing"

func TestFeedSystemAlerts_FeedKey(t *testing.T) {
	var f FeedSystemAlerts

	if k := f.FeedKey(); k != FeedKeySystemAlerts {
		t.Errorf("expect '%s' got '%s'", FeedKeySystemAlerts, k)
	}
}
