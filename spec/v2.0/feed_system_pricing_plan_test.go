package gbfsspec

import "testing"

func TestFeedSystemPricingPlans_FeedKey(t *testing.T) {
	var f FeedSystemPricingPlans

	if k := f.FeedKey(); k != FeedKeySystemPricingPlans {
		t.Errorf("expect '%s' got '%s'", FeedKeySystemPricingPlans, k)
	}
}
