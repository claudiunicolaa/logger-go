package test

import (
	"testing"
)

func TestHello(t *testing.T) {
	testHelper := GetTestHelper()
	if testHelper.DEMO_URL == "" {
		t.Error("Helper DEMO_URL is empty")
	}
}
