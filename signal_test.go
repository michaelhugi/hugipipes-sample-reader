package hugipipes_sample

import (
	tu "github.com/informaticon/lib.go.base.test-utils"
	"testing"
)

func TestNewSignal(t *testing.T) {
	sig, err := NewSignal("samples/stecker/Ma_66_st.wav")
	tu.AssertNErr(err)
	tu.AssertEq(sig.Wav.NumChannels, 1, t)
	sig2, err := NewSignal("samples/LRMonoPhase4.wav")
	tu.AssertNErr(err)
	tu.AssertEq(sig2.Wav.NumChannels, 2, t)
	tu.AssertEq(len(sig2.GetSamples(Left)), len(sig2.GetSamples(Right)), t)

	specL, err := sig.MonoSpectrum(Left)
	tu.AssertNErr(err)
	println(specL.Points[0].String())
}
