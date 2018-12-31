package gbfs

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

var (
	baseURL = "https://fake.gbfs.org"
	lang    = "en"
)

func getHTTPClient(t *testing.T, opts ...HTTPOption) *HTTPClient {
	c, err := NewHTTPClient(opts...)

	if err != nil {
		t.Errorf("expected nil got %s", err)
		return nil
	}

	hc, ok := c.(*HTTPClient)

	if !ok {
		t.Error("expected type HTTPClient")
		return nil
	}

	return hc
}

func TestHTTPOptionBaseURL(t *testing.T) {
	c := getHTTPClient(t, HTTPOptionBaseURL(baseURL))
	if c == nil {
		t.FailNow()
	}

	if c.BaseURL != baseURL {
		t.Errorf("expected %s got %s", baseURL, c.BaseURL)
	}
}

func TestHTTPOptionLanguage(t *testing.T) {
	c := getHTTPClient(t, HTTPOptionBaseURL(baseURL), HTTPOptionLanguage(lang))
	if c == nil {
		t.FailNow()
	}

	if c.Language != lang {
		t.Errorf("expected %s got %s", lang, c.Language)
	}
}

func TestHTTPOptionForceURL(t *testing.T) {
	url := "https://other.domain.tld/free_bike_status"

	c := getHTTPClient(t, HTTPOptionBaseURL(baseURL), HTTPOptionForceURL(FeedKeyFreeBikeStatus, url))
	if c == nil {
		t.FailNow()
	}

	if s := len(c.urls); s != 1 {
		t.Errorf("expected 1 url got %d", s)
	}

	if c.urls[FeedKeyFreeBikeStatus] != url {
		t.Errorf("expected %s got %s", url, c.urls[FeedKeyFreeBikeStatus])
	}
}

func TestHTTPOptionClient(t *testing.T) {
	h := http.Client{Timeout: time.Second}

	c := getHTTPClient(t, HTTPOptionBaseURL(baseURL), HTTPOptionClient(h))
	if c == nil {
		t.FailNow()
	}

	if c.client.Timeout != h.Timeout {
		t.Errorf("expected %d got %d", h.Timeout, c.client.Timeout)
	}
}

func TestNewHTTPClient_NoBaseURL(t *testing.T) {
	_, err := NewHTTPClient()

	if err == nil {
		t.Errorf("expected err got nil")
	}

	if err != ErrBaseURLNotSet {
		t.Errorf("expected %s got %s", ErrBaseURLNotSet, err)
	}
}

func TestHTTPClient_url(t *testing.T) {
	urls := map[FeedKey]string{
		FeedKeyFreeBikeStatus:     "https://other.domain.tld/free_bike_status.json",
		FeedKeyStationInformation: "https://another.domain.org/fr/station_information",
	}

	opts := []HTTPOption{
		HTTPOptionBaseURL(baseURL),
		HTTPOptionLanguage(lang),
	}

	for k, v := range urls {
		opts = append(opts, HTTPOptionForceURL(k, v))
	}

	c := getHTTPClient(t, opts...)
	if c == nil {
		t.FailNow()
	}

	expected := map[FeedKey]string{
		FeedKeyAutoDiscovery:      fmt.Sprintf("%s/%s/%s.json", baseURL, lang, FeedKeyAutoDiscovery),
		FeedKeyStationInformation: "https://another.domain.org/fr/station_information",
		FeedKeyStationStatus:      fmt.Sprintf("%s/%s/%s.json", baseURL, lang, FeedKeyStationStatus),
		FeedKeyFreeBikeStatus:     "https://other.domain.tld/free_bike_status.json",
	}

	for k, v := range expected {
		if u := c.url(k); u != v {
			t.Errorf("expected %s got %s", v, u)
		}
	}
}

func TestHTTPClient_url_noLanguage(t *testing.T) {
	urls := map[FeedKey]string{
		FeedKeyFreeBikeStatus:     "https://other.domain.tld/free_bike_status.json",
		FeedKeyStationInformation: "https://another.domain.org/fr/station_information",
	}

	opts := []HTTPOption{HTTPOptionBaseURL(baseURL)}

	for k, v := range urls {
		opts = append(opts, HTTPOptionForceURL(k, v))
	}

	c := getHTTPClient(t, opts...)
	if c == nil {
		t.FailNow()
	}

	expected := map[FeedKey]string{
		FeedKeyAutoDiscovery:      fmt.Sprintf("%s/%s.json", baseURL, FeedKeyAutoDiscovery),
		FeedKeyStationInformation: "https://another.domain.org/fr/station_information",
		FeedKeyStationStatus:      fmt.Sprintf("%s/%s.json", baseURL, FeedKeyStationStatus),
		FeedKeyFreeBikeStatus:     "https://other.domain.tld/free_bike_status.json",
	}

	for k, v := range expected {
		if u := c.url(k); u != v {
			t.Errorf("expected %s got %s", v, u)
		}
	}
}

func TestHTTPClient_url_cache(t *testing.T) {
	c := getHTTPClient(t, HTTPOptionBaseURL(baseURL))
	if c == nil {
		t.FailNow()
	}

	if s := len(c.urls); s > 0 {
		t.Errorf("expected no url to be cached got %d", s)
	}

	_ = c.url(FeedKeyStationInformation)

	if s := len(c.urls); s != 1 {
		t.Errorf("expected 1 url to be cached got %d", s)
	}

	if _, ok := c.urls[FeedKeyStationInformation]; !ok {
		t.Errorf("expected key %s to be set", FeedKeyStationInformation)
	}
}

func TestHTTPClient_Get_success(t *testing.T) {
	c := getHTTPClient(t, serverHTTPOptions...)
	if c == nil {
		t.FailNow()
	}

	j, err := c.Get(FeedKeyAutoDiscovery)

	if err != nil {
		t.Errorf("expected nil got %s", err)
	}

	if j == nil {
		t.Fatalf("expected *JSON got nil")
	}

	if j.LastUpdated != 1545053599 {
		t.Errorf("expected 1545053599 got %d", j.LastUpdated)
	}
}

func TestHTTPClient_Get_notFound(t *testing.T) {
	c := getHTTPClient(t, serverHTTPOptions...)
	if c == nil {
		t.FailNow()
	}

	_, err := c.Get(FeedKey("toto"))

	if err != ErrFeedNotExist {
		t.Errorf("expected %s got %s", ErrFeedNotExist, err)
	}
}

func TestHTTPClient_Get_invalidJSON(t *testing.T) {
	c := getHTTPClient(t, serverHTTPOptions...)
	if c == nil {
		t.FailNow()
	}

	_, err := c.Get(FeedKey("bad"))

	if err == nil {
		t.Errorf("expected err got nil")
	}
}

func TestHTTPClient_Get_serverError(t *testing.T) {
	c := getHTTPClient(t, serverHTTPOptions...)
	if c == nil {
		t.FailNow()
	}

	_, err := c.Get(FeedKey("500"))

	if err.Error() != "invalid status code (500)" {
		t.Errorf("expected \"invalid status code (500)\" got \"%s\"", err)
	}
}

func TestHTTPClient_Get_invalidURL(t *testing.T) {
	c := getHTTPClient(t, serverHTTPOptions...)
	if c == nil {
		t.FailNow()
	}

	c.urls[FeedKeyAutoDiscovery] = "-"
	_, err := c.Get(FeedKeyAutoDiscovery)

	if err.Error() != "Get -: unsupported protocol scheme \"\"" {
		t.Errorf("expected \"Get -: unsupported protocol scheme \"\"\" got \"%s\"", err)
	}
}

func TestHTTPClient_AutoDiscovery(t *testing.T) {
	c := getHTTPClient(t, serverHTTPOptions...)
	if c == nil {
		t.FailNow()
	}

	a, err := c.AutoDiscovery()
	if err != nil {
		t.Fatalf("expected nil got %s", err)
	}

	if s := len(a.Languages); s != 2 {
		t.Errorf("expected 2 got %d", s)
	}

	if s := len(a.Languages["en"].Feeds); s != 5 {
		t.Errorf("expected 5 got %d", s)
	}
}

func TestHTTPClient_AutoDiscovery_notFound(t *testing.T) {
	c := getHTTPClient(t, serverHTTPOptions...)
	if c == nil {
		t.FailNow()
	}

	c.urls[FeedKeyAutoDiscovery] = fmt.Sprintf("%s/en/404.json", server.URL)

	_, err := c.AutoDiscovery()
	if err != ErrFeedNotExist {
		t.Errorf("expected ErrFeedNotExist got %s", err)
	}
}

func TestHTTPClient_SystemInformation(t *testing.T) {
	c := getHTTPClient(t, serverHTTPOptions...)
	if c == nil {
		t.FailNow()
	}

	si, err := c.SystemInformation()
	if err != nil {
		t.Fatalf("expected nil got %s", err)
	}

	if si.SystemID != "BA" {
		t.Errorf("expected BA got %s", si.SystemID)
	}
}

func TestHTTPClient_SystemInformation_notFound(t *testing.T) {
	c := getHTTPClient(t, serverHTTPOptions...)
	if c == nil {
		t.FailNow()
	}

	c.urls[FeedKeySystemInformation] = fmt.Sprintf("%s/en/404.json", server.URL)

	_, err := c.SystemInformation()
	if err != ErrFeedNotExist {
		t.Errorf("expected ErrFeedNotExist got %s", err)
	}
}

func TestHTTPClient_StationsInformation(t *testing.T) {
	c := getHTTPClient(t, serverHTTPOptions...)
	if c == nil {
		t.FailNow()
	}

	si, err := c.StationsInformation()
	if err != nil {
		t.Fatalf("expected nil got %s", err)
	}

	if s := len(si.Stations); s != 10 {
		t.Errorf("expected 10 got %d", s)
	}

	if si.Stations[0].StationID != "74" {
		t.Errorf("expected 74 got %s", si.Stations[0].StationID)
	}

	if si.Stations[0].Capacity != 27 {
		t.Errorf("expected 27 got %d", si.Stations[0].Capacity)
	}
}

func TestHTTPClient_StationsInformation_notFound(t *testing.T) {
	c := getHTTPClient(t, serverHTTPOptions...)
	if c == nil {
		t.FailNow()
	}

	c.urls[FeedKeyStationInformation] = fmt.Sprintf("%s/en/404.json", server.URL)

	_, err := c.StationsInformation()
	if err != ErrFeedNotExist {
		t.Errorf("expected ErrFeedNotExist got %s", err)
	}
}

func TestHTTPClient_StationsStatus(t *testing.T) {
	c := getHTTPClient(t, serverHTTPOptions...)
	if c == nil {
		t.FailNow()
	}

	si, err := c.StationsStatus()
	if err != nil {
		t.Fatalf("expected nil got %s", err)
	}

	if s := len(si.Stations); s != 10 {
		t.Errorf("expected 10 got %d", s)
	}

	if si.Stations[0].StationID != "74" {
		t.Errorf("expected 74 got %s", si.Stations[0].StationID)
	}

	if si.Stations[0].NumBikesAvailable != 13 {
		t.Errorf("expected 13 got %d", si.Stations[0].NumBikesAvailable)
	}
}

func TestHTTPClient_StationsStatus_notFound(t *testing.T) {
	c := getHTTPClient(t, serverHTTPOptions...)
	if c == nil {
		t.FailNow()
	}

	c.urls[FeedKeyStationStatus] = fmt.Sprintf("%s/en/404.json", server.URL)

	_, err := c.StationsStatus()
	if err != ErrFeedNotExist {
		t.Errorf("expected ErrFeedNotExist got %s", err)
	}
}

func TestHTTPClient_SystemRegions(t *testing.T) {
	c := getHTTPClient(t, serverHTTPOptions...)
	if c == nil {
		t.FailNow()
	}

	sr, err := c.SystemRegions()
	if err != nil {
		t.Fatalf("expected nil got %s", err)
	}

	if s := len(sr.Regions); s != 6 {
		t.Errorf("expected 6 got %d", s)
	}

	if sr.Regions[1].Name != "San Jose" {
		t.Errorf("expected \"San Jose\" got \"%s\"", sr.Regions[1].Name)
	}
}

func TestHTTPClient_SystemRegions_notFound(t *testing.T) {
	c := getHTTPClient(t, serverHTTPOptions...)
	if c == nil {
		t.FailNow()
	}

	c.urls[FeedKeySystemRegions] = fmt.Sprintf("%s/en/404.json", server.URL)

	_, err := c.SystemRegions()
	if err != ErrFeedNotExist {
		t.Errorf("expected ErrFeedNotExist got %s", err)
	}
}

func TestHTTPClient_SystemHours(t *testing.T) {
	c := getHTTPClient(t, serverHTTPOptions...)
	if c == nil {
		t.FailNow()
	}

	sh, err := c.SystemHours()
	if err != nil {
		t.Fatalf("expected nil got %s", err)
	}

	if s := len(sh.RentalHours); s != 1 {
		t.Errorf("expected 1 got %d", s)
	}

	if s := len(sh.RentalHours[0].UserTypes); s != 2 {
		t.Errorf("expected 2 got %d", s)
	}

	if s := len(sh.RentalHours[0].Days); s != 7 {
		t.Errorf("expected 7 got %d", s)
	}

	if sh.JSON.TTL != 600 {
		t.Errorf("expected 600 got %d", sh.JSON.TTL)
	}
}

func TestHTTPClient_SystemHours_notFound(t *testing.T) {
	c := getHTTPClient(t, serverHTTPOptions...)
	if c == nil {
		t.FailNow()
	}

	c.urls[FeedKeySystemHours] = fmt.Sprintf("%s/en/404.json", server.URL)

	_, err := c.SystemHours()
	if err != ErrFeedNotExist {
		t.Errorf("expected ErrFeedNotExist got %s", err)
	}
}

func TestHTTPClient_SystemAlerts(t *testing.T) {
	c := getHTTPClient(t, serverHTTPOptions...)
	if c == nil {
		t.FailNow()
	}

	sa, err := c.SystemAlerts()
	if err != nil {
		t.Fatalf("expected nil got %s", err)
	}

	if s := len(sa.Alerts); s != 1 {
		t.Errorf("expected 1 got %d", s)
	}

	if id := sa.Alerts[0].AlertID; id != "1" {
		t.Errorf("expected '1' got '%s'", id)
	}

	if s := sa.Alerts[0].Type; AlertType(s) != AlertTypeSystemClosure {
		t.Errorf("expected '%s' got '%s'", AlertTypeSystemClosure, s)
	}

	if sa.JSON.TTL != 10 {
		t.Errorf("expected 10 got %d", sa.JSON.TTL)
	}
}

func TestHTTPClient_SystemAlerts_notFound(t *testing.T) {
	c := getHTTPClient(t, serverHTTPOptions...)
	if c == nil {
		t.FailNow()
	}

	c.urls[FeedKeySystemAlerts] = fmt.Sprintf("%s/en/404.json", server.URL)

	_, err := c.SystemAlerts()
	if err != ErrFeedNotExist {
		t.Errorf("expected ErrFeedNotExist got %s", err)
	}
}
