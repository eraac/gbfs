package gbfsspec

import "testing"

func TestFeedSystemInformation_FeedKey(t *testing.T) {
	var f FeedSystemInformation

	if k := f.FeedKey(); k != FeedKeySystemInformation {
		t.Errorf("expect '%s' got '%s'", FeedKeySystemInformation, k)
	}
}

func TestSystemInformationData_GetStartDate(t *testing.T) {
	f := SystemInformationData{
		StartDate: "2020-01-02",
		Timezone: "europe/paris",
	}

	s, err := f.GetStartDate()
	if err != nil {
		t.Errorf("expect 'nil' got '%s'", err)
		t.FailNow()
	}

	y, m, d := s.Date()
	if y != 2020 {
		t.Errorf("expect '2020' got '%d'", y)
	}

	if m != 1 {
		t.Errorf("expect '1' got '%d'", y)
	}

	if d != 2 {
		t.Errorf("expect '2' got '%d'", y)
	}
}
