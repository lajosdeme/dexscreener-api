package dexscreenerapi

import (
	"testing"
)

func TestGetLatestTokenProfiles(t *testing.T) {
	_, err := GetLatestTokenProfiles()
	if err != nil {
		t.Error("Error: ", err)
	}
}
