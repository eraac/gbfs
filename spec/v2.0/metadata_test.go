package gbfsspec

import (
	"testing"
	"time"
)

func TestMetadata_IsExpired(t *testing.T) {
	ii := []struct{
		in Metadata
		out bool
	}{
		{
			in: Metadata{
				LastUpdated: Timestamp(time.Now().Unix()),
				TTL:         10000,
			}, out: false,
		},
		{
			in: Metadata{
				LastUpdated: Timestamp(time.Now().Unix()),
				TTL:         0,
			}, out: true,
		},
	}

	for _, i := range ii {
		if got := i.in.IsExpired(); got != i.out {
			t.Errorf("expect '%t' got '%t'", i.out, got)
		}
	}
}
