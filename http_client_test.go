package gbfs

import (
	"errors"
	gbfsspec "github.com/Eraac/gbfs/spec/v2.0"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var server *httptest.Server

func init() {
	server = httptest.NewServer(http.FileServer(http.Dir("test/gbfs")))
}

func TestNewHTTPClient(t *testing.T) {
	baseURL := "https://domain.tld"

	c, err := NewHTTPClient(HTTPOptionBaseURL(baseURL))

	if err != nil {
		t.Errorf("expect 'nil' got '%s'", err)
		t.FailNow()
	}

	ch, ok := c.(*HTTPClient)
	if !ok {
		t.Errorf("expect type *HTTPClient, got '%T'", c)
		t.FailNow()
	}

	if ch.baseURL != baseURL {
		t.Errorf("expect '%s' got '%s'", baseURL, ch.baseURL)
	}

	if l := len(ch.urls); l != 0 {
		t.Errorf("expect '0' got '%d'", l)
	}
}

func TestNewHTTPClientError(t *testing.T) {
	_, err := NewHTTPClient()

	if !errors.Is(err, ErrBaseURLMissing) {
		t.Errorf("expect '%s' got '%s'", ErrBaseURLMissing, err)
	}
}

func TestHTTPClient_ForceURLs(t *testing.T) {
	key1, key2, key3 := "https://domain.tld/key1.json", "https://domain.tld/key2.json", "https://domain.tld/key3.json"

	c, err := NewHTTPClient(
		HTTPOptionBaseURL("https://domain.tld"),
		HTTPOptionForceURL("key1", key1),
	)
	if err != nil {
		t.Errorf("expect 'nil' got '%s'", err)
		t.FailNow()
	}

	c.ForceURLs(map[string]string{
		"key2": key2,
	}, false)

	ch, ok := c.(*HTTPClient)
	if !ok {
		t.Errorf("expect type *HTTPClient, got '%T'", c)
		t.FailNow()
	}

	if ch.urls["key1"] != key1 {
		t.Errorf("expect '%s' got '%s'", key1, ch.urls["key1"])
	}

	if l := len(ch.urls); l != 2 {
		t.Errorf("expect '2' got '%d'", l)
	}

	ch.ForceURLs(map[string]string{
		"key3": key3,
	}, true)

	if ch.urls["key3"] != key3 {
		t.Errorf("expect '%s' got '%s'", key3, ch.urls["key3"])
	}

	if l := len(ch.urls); l != 1 {
		t.Errorf("expect '1' got '%d'", l)
	}
}

func TestHTTPClient_GetAndRefresh(t *testing.T) {
	c, err := NewHTTPClient(
		HTTPOptionBaseURL(server.URL),
		HTTPOptionLanguage("en"),
	)

	if err != nil {
		t.Errorf("expect 'nil' got '%s'", err)
		t.FailNow()
	}

	var si gbfsspec.FeedSystemInformation
	if err := c.Get(gbfsspec.FeedKeySystemInformation, &si); err != nil {
		t.Errorf("expect 'nil' got '%s'", err)
		t.FailNow()
	}

	if si.Data.StartDate != "2017-05-14" {
		t.Errorf("expect '2017-05-14' got '%s'", si.Data.StartDate)
	}

	if si.LastUpdated != 1589230640 {
		t.Errorf("expect '1589230640' got '%d'", si.LastUpdated)
	}

	if err := c.Refresh(&si, false); err != nil {
		t.Errorf("expect 'nil' got '%s'", err)
	}

	if err := c.Refresh(&si, true); err != nil {
		t.Errorf("expect 'nil' got '%s'", err)
	}
}

func TestHTTPClient_GetInvalid(t *testing.T) {
	c, err := NewHTTPClient(
		HTTPOptionBaseURL(server.URL),
		HTTPOptionLanguage("en"),
		HTTPOptionClient(http.Client{Timeout: 10 * time.Second}),
	)

	if err != nil {
		t.Errorf("expect 'nil' got '%s'", err)
		t.FailNow()
	}

	var si gbfsspec.FeedSystemInformation

	err = c.Get(gbfsspec.FeedKeyGBFSVersions, &si)
	if !errors.Is(err, ErrInvalidFeed) {
		t.Errorf("expect '%s' got '%s'", ErrInvalidFeed, err)
	}

	var g gbfsspec.FeedGBFSVersions
	err = c.Get(gbfsspec.FeedKeyGBFSVersions, &g)
	if !errors.Is(err, ErrFeedNotExist) {
		t.Errorf("expect '%s' got '%s'", ErrFeedNotExist, err)
	}
}

func TestHTTPClient_url_language(t *testing.T) {
	ii := []struct{
		in, out string
	}{
		{in: "key1", out: "https://domain.tld/en/key1.json"},
		{in: "key2", out: "https://other.domain.tld/key2.json"},
	}

	c, err := NewHTTPClient(
		HTTPOptionBaseURL("https://domain.tld"),
		HTTPOptionLanguage("en"),
		HTTPOptionForceURL("key2", "https://other.domain.tld/key2.json"),
	)
	if err != nil {
		t.Errorf("expect 'nil' got '%s'", err)
		t.FailNow()
	}

	ch, ok := c.(*HTTPClient)
	if !ok {
		t.Errorf("expect type *HTTPClient, got '%T'", c)
		t.FailNow()
	}

	for _, i := range ii {
		if o := ch.url(i.in); o != i.out {
			t.Errorf("expect '%s' got '%s", i.out, o)
		}
	}
}

func TestHTTPClient_url_noLanguage(t *testing.T) {
	ii := []struct{
		in, out string
	}{
		{in: "key1", out: "https://domain.tld/key1.json"},
	}

	c, err := NewHTTPClient(HTTPOptionBaseURL("https://domain.tld"))
	if err != nil {
		t.Errorf("expect 'nil' got '%s'", err)
		t.FailNow()
	}

	ch, ok := c.(*HTTPClient)
	if !ok {
		t.Errorf("expect type *HTTPClient, got '%T'", c)
		t.FailNow()
	}

	for _, i := range ii {
		if o := ch.url(i.in); o != i.out {
			t.Errorf("expect '%s' got '%s", i.out, o)
		}
	}
}
