package util

import (
	"os"
	"testing"
)

func SkipCI(t *testing.T) bool {
	if os.Getenv("GITHUB_ACTIONS") != "" {
		t.Skip("Skipping testing in CI environment")
		return true
	}
	return false
}
