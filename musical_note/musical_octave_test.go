package musical_note

import (
	tu "github.com/informaticon/lib.go.base.test-utils"
	"testing"
)

func TestOctaveFromFreq(t *testing.T) {
	o1, err := musicalOctaveFromFrequency(16.35)
	tu.AssertNErr(err)
	tu.AssertEq(o1.octaveNumber, 0, t)

	o2, err := musicalOctaveFromFrequency(7902.13)
	tu.AssertNErr(err)
	tu.AssertEq(o2.octaveNumber, 8, t)

	o3, err := musicalOctaveFromFrequency(987.77)
	tu.AssertNErr(err)
	tu.AssertEq(o3.octaveNumber, 5, t)

	o4, err := musicalOctaveFromFrequency(995.77)
	tu.AssertNErr(err)
	tu.AssertEq(o4.octaveNumber, 5, t)

	o5, err := musicalOctaveFromFrequency(1030.0)
	tu.AssertNErr(err)
	tu.AssertEq(o5.octaveNumber, 6, t)
}
