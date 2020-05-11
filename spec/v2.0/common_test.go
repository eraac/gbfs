package gbfsspec

import (
	"testing"
	"time"
)

func TestDate_ToTime(t *testing.T) {
	ii := []struct{
		in Date
		tz string
		want time.Time
		shouldReturnErr bool
	}{
		{
			in: Date("2020-01-02"),
			tz: "UTC",
			want: time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
			shouldReturnErr: false,
		},
		{
			in: Date("2020-012"),
			tz: "UTC",
			want: time.Date(2020, 1, 2, 0, 0, 0, 0, time.Local),
			shouldReturnErr: true,
		},
		{
			in: Date("2020-012"),
			tz: "invalid-timezone",
			want: time.Date(2020, 1, 2, 0, 0, 0, 0, time.Local),
			shouldReturnErr: true,
		},
	}

	for _, i := range ii {
		got, err := i.in.ToTime(i.tz)
		if !i.shouldReturnErr && got != i.want {
			t.Errorf("expect '%s' got '%s'", i.want, got)
		}

		if err == nil && i.shouldReturnErr {
			t.Errorf("expect error got nil")
		}

		if err != nil && !i.shouldReturnErr {
			t.Errorf("expect 'nil' got '%s'", err)
		}
	}
}

func TestTime_Clock(t *testing.T) {
	ii := []struct{
		in Time
		h, m, s int
	}{
		{in: Time("23:43:01"), h: 23, m: 43, s: 1},
		{in: Time("10:01:00"), h: 10, m: 1, s: 0},
		{in: Time("nope"), h: 0, m: 0, s: 0},
		{in: Time("1:59:00"), h: 1, m: 59, s: 0},
	}

	for _, i := range ii {
		h, m, s := i.in.Clock()

		if i.h != h {
			t.Errorf("expect '%d' got '%d'", i.h, h)
		}

		if i.m != m {
			t.Errorf("expect '%d' got '%d'", i.m, m)
		}

		if i.s != s {
			t.Errorf("expect '%d' got '%d'", i.s, s)
		}
	}
}

func TestTimestamp_ToTime(t *testing.T) {
	ii := []struct{
		in Timestamp
		out int64
	}{
		{in: Timestamp(1589226008), out: 1589226008},
		{in: Timestamp(0), out: 0},
	}

	for _, i := range ii {
		if got := i.in.ToTime().Unix(); got != i.out {
			t.Errorf("expect '%d' got '%d'", i.out, got)
		}
	}
}

func TestPrice_UnmarshalJSON(t *testing.T) {
	ii := []struct{
		in []byte
		out Price
	}{
		{in: []byte("\"12\""), out: Price("12")},
		{in: []byte("12.43"), out: Price("12.43")},
		{in: []byte(""), out: Price("")},
	}

	for _, i := range ii {
		var p Price
		if err := p.UnmarshalJSON(i.in); err != nil {
			t.Errorf("expect 'nil' got '%s'", err)
			continue
		}

		if p != i.out {
			t.Errorf("expect '%s' got '%s'", i.out, p)
		}
	}
}

func TestBoolean_UnmarshalJSON(t *testing.T) {
	ii := []struct{
		in []byte
		out Boolean
	}{
		{in: []byte("1"), out: true},
		{in: []byte("\"1\""), out: true},
		{in: []byte("true"), out: true},
		{in: []byte("TRUE"), out: true},
		{in: []byte("\"TRUe\""), out: true},
		{in: []byte(""), out: false},
		{in: []byte("O"), out: false},
		{in: []byte("\"0\""), out: false},
		{in: []byte("false"), out: false},
		{in: []byte("FALSE"), out: false},
		{in: []byte("\"FASLe\""), out: false},
	}

	for _, i := range ii {
		var b Boolean
		if err := b.UnmarshalJSON(i.in); err != nil {
			t.Errorf("expect 'nil' got '%s'", err)
			continue
		}

		if b != i.out {
			t.Errorf("expect '%t' got '%t', for '%s'", i.out, b, i.in)
		}
	}
}
