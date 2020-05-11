package gbfs

import "testing"

func TestError_Error(t *testing.T) {
	if ErrInvalidFeed.Error() != "invalid feed" {
		t.Errorf("expect 'invalid feed' got '%s'", ErrInvalidFeed.Error())
	}
}
