package gbfs

import (
	"encoding/json"
	"testing"
)

type unmarshalBooleanTest struct {
	in  []byte
	out Boolean
}

var unmarshalBooleanTests = []unmarshalBooleanTest{
	{in: []byte("1"), out: true},
	{in: []byte("true"), out: true},
	{in: []byte("false"), out: false},
	{in: []byte("0"), out: false},
	{in: []byte("\"string\""), out: false},
}

func TestBoolean_UnmarshalJSON(t *testing.T) {
	if len(unmarshalBooleanTests) == 0 {
		t.Errorf("no test value available for testing")
	}

	for _, tt := range unmarshalBooleanTests {
		var b Boolean

		if err := json.Unmarshal(tt.in, &b); err != nil {
			t.Fatalf("json unmarshal err: %s", err)
		}

		if tt.out != b {
			t.Errorf("expected %t got %t, for value %s", tt.out, b, tt.in)
		}
	}
}
