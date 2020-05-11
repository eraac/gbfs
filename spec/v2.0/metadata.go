package gbfsspec

import "time"

type (
	Metadata struct {
		// Last time the data in the feed was updated.
		LastUpdated Timestamp `json:"last_updated"`

		// Number of seconds before the data in the feed will be updated again (0 if the data should always be refreshed).
		TTL int `json:"ttl"`

		// GBFS version number to which the feed confirms, according to the versioning framework.
		Version string `json:"version"`
	}
)

// IsExpired return true if TTL has been reached
func (m Metadata) IsExpired() bool {
	if m.TTL == 0 {
		return true
	}

	return time.Now().Unix() > int64(m.LastUpdated)+int64(m.TTL)
}
