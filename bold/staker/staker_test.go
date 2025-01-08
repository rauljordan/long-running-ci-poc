//go:build challengetest && !race

package staker_test

import (
	"testing"

	"github.com/rauljordan/long-running-ci-poc/bold/staker"
)

func TestStaker(t *testing.T) {
	stkr := staker.New()
	if stkr.Bar() != "Bar" {
		t.Fatal("Bar() should return 'Bar'")
	}
}
