package logger

import (
	"testing"
)

func TestHelper(t *testing.T) {
	testHelper := NewTestHelper()
	if testHelper.demoURL == "" {
		t.Error("Helper DEMO_URL is empty")
	}
}
