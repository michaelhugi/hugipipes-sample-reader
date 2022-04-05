package hugipipes_sample

import (
	"fmt"
	tu "github.com/informaticon/lib.go.base.test-utils"
	"testing"
)

func TestNewSignal(t *testing.T) {
	sig, err := NewSignal("samples/stecker/Ma_57_st.wav")
	tu.AssertNErr(err)
	tu.AssertEq(sig.Wav.NumChannels, 1, t)
	sig2, err := NewSignal("samples/LRMonoPhase4.wav")
	tu.AssertNErr(err)
	tu.AssertEq(sig2.Wav.NumChannels, 2, t)
	tu.AssertEq(len(sig2.GetSamples(Left)), len(sig2.GetSamples(Right)), t)
	tu.AssertVNErr(sig2.MonoSpectrum(Left))

	specL, err := sig.MonoSpectrum(Left)
	tu.AssertNErr(err)
	tu.AssertNErr(DrawSpectrum(specL, "spectrum"))

	fmt.Printf("Max: %f\n", specL.MaxAmpl)
}
