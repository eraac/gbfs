package gbfsspec

import "testing"

func TestFeedSystemCalendars_FeedKey(t *testing.T) {
	var f FeedSystemCalendars

	if k := f.FeedKey(); k != FeedKeySystemCalendar {
		t.Errorf("expect '%s' got '%s'", FeedKeySystemCalendar, k)
	}
}
