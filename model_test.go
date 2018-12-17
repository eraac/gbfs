package gbfs

import (
	"testing"
	"time"
)

type jsonTest struct {
	in  *JSON
	out bool
}

var (
	n = time.Now().Unix()

	jsonTests = []jsonTest{
		{in: &JSON{LastUpdated: n, TTL: 0}, out: true},
		{in: &JSON{LastUpdated: n, TTL: 100}, out: false},
		{in: &JSON{LastUpdated: n - 1000, TTL: 100}, out: true},
	}
)

func TestJSONIsOutdated(t *testing.T) {
	if len(jsonTests) == 0 {
		t.Errorf("no test value available for testing")
	}

	for _, tt := range jsonTests {
		isOutdated := tt.in.IsOutdated()

		if isOutdated != tt.out {
			t.Errorf("expected %t got %t for values %d, %d", tt.out, isOutdated, tt.in.LastUpdated, tt.in.TTL)
		}
	}
}
