package gbfs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	HTTPClient struct {
		client http.Client
		urls   map[FeedKey]string

		BaseURL  string
		Language string
	}
)

var (
	ErrBaseURLNotSet = fmt.Errorf("base url not set")
	ErrFeedNotExist  = fmt.Errorf("feed not found")
)

func NewHTTPClient(opts ...HTTPOption) (Client, error) {
	c := &HTTPClient{urls: map[FeedKey]string{}}

	for _, opt := range opts {
		opt(c)
	}

	if c.BaseURL == "" {
		return nil, ErrBaseURLNotSet
	}

	return c, nil
}

func (c *HTTPClient) Get(key FeedKey) (j *JSON, err error) {
	req := new(http.Request)
	if req, err = http.NewRequest(http.MethodGet, c.url(key), nil); err != nil {
		return
	}

	res := new(http.Response)
	if res, err = c.client.Do(req); err != nil {
		return
	}
	defer func() {_ = res.Body.Close()}()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			return nil, ErrFeedNotExist
		}

		return j, fmt.Errorf("invalid status code (%d)", res.StatusCode)
	}

	j = new(JSON)

	return j, json.NewDecoder(res.Body).Decode(j)
}

func (c *HTTPClient) AutoDiscovery() (g *GBFS, err error) {
	j := new(JSON)
	if j, err = c.Get(FeedKeyAutoDiscovery); err != nil {
		return
	}

	g = new(GBFS)
	g.JSON = j

	return g, json.Unmarshal(j.Data, &g.Languages)
}

func (c *HTTPClient) SystemInformation() (si *SystemInformation, err error) {
	j := new(JSON)
	if j, err = c.Get(FeedKeySystemInformation); err != nil {
		return
	}

	si = new(SystemInformation)
	si.JSON = j

	return si, json.Unmarshal(j.Data, si)
}

func (c *HTTPClient) StationsInformation() (si *StationsInformation, err error) {
	j := new(JSON)
	if j, err = c.Get(FeedKeyStationInformation); err != nil {
		return
	}

	si = new(StationsInformation)
	si.JSON = j

	return si, json.Unmarshal(j.Data, si)
}

func (c *HTTPClient) StationsStatus() (ss *StationsStatus, err error) {
	j := new(JSON)
	if j, err = c.Get(FeedKeyStationStatus); err != nil {
		return
	}

	ss = new(StationsStatus)
	ss.JSON = j

	return ss, json.Unmarshal(j.Data, ss)
}

func (c *HTTPClient) FreeBikeStatus() (fbs *FreeBikeStatus, err error) {
	j := new(JSON)
	if j, err = c.Get(FeedKeyFreeBikeStatus); err != nil {
		return
	}

	fbs = new(FreeBikeStatus)
	fbs.JSON = j

	return fbs, json.Unmarshal(j.Data, fbs)
}

func (c *HTTPClient) SystemHours() (sh *SystemHours, err error) {
	j := new(JSON)
	if j, err = c.Get(FeedKeySystemHours); err != nil {
		return
	}

	sh = new(SystemHours)
	sh.JSON = j

	return sh, json.Unmarshal(j.Data, sh)
}

func (c *HTTPClient) SystemCalendar() (sc *SystemCalendar, err error) {
	j := new(JSON)
	if j, err = c.Get(FeedKeySystemCalendar); err != nil {
		return
	}

	sc = new(SystemCalendar)
	sc.JSON = j

	return sc, json.Unmarshal(j.Data, sc)
}

func (c *HTTPClient) SystemRegions() (sr *SystemRegions, err error) {
	j := new(JSON)
	if j, err = c.Get(FeedKeySystemRegions); err != nil {
		return
	}

	sr = new(SystemRegions)
	sr.JSON = j

	return sr, json.Unmarshal(j.Data, sr)
}

func (c *HTTPClient) SystemPricingPlans() (spp *SystemPricingPlans, err error) {
	j := new(JSON)
	if j, err = c.Get(FeedKeySystemPricingPlans); err != nil {
		return
	}

	spp = new(SystemPricingPlans)
	spp.JSON = j

	return spp, json.Unmarshal(j.Data, spp)
}

func (c *HTTPClient) SystemAlerts() (sa *SystemAlerts, err error) {
	j := new(JSON)
	if j, err = c.Get(FeedKeySystemAlerts); err != nil {
		return
	}

	sa = new(SystemAlerts)
	sa.JSON = j

	return sa, json.Unmarshal(j.Data, sa)
}

func (c *HTTPClient) url(key FeedKey) string {
	if u, ok := c.urls[key]; ok {
		return u
	}

	l := c.Language

	if l != "" {
		l = fmt.Sprintf("%s/", l)
	}

	c.urls[key] = fmt.Sprintf("%s/%s%s.json", c.BaseURL, l, key)

	return c.urls[key]
}

// =========
//  OPTIONS
// =========

func HTTPOptionClient(h http.Client) HTTPOption {
	return func(c *HTTPClient) {
		c.client = h
	}
}

func HTTPOptionBaseURL(url string) HTTPOption {
	return func(c *HTTPClient) {
		c.BaseURL = url
	}
}

func HTTPOptionLanguage(lang string) HTTPOption {
	return func(c *HTTPClient) {
		c.Language = lang
	}
}

func HTTPOptionForceURL(key FeedKey, url string) HTTPOption {
	return func(c *HTTPClient) {
		c.urls[key] = url
	}
}
