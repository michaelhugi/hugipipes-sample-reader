package hugipipes_sample

import (
	tu "github.com/informaticon/lib.go.base.test-utils"
	"testing"
)

func TestNewSignal(t *testing.T) {
	sig, err := NewSignalFromWav("samples/stecker/Ma_57_st.wav")
	tu.AssertNErr(err)

	tu.AssertEq(sig.Wav.NumChannels, 1, t)
	sig2, err := NewSignalFromWav("samples/LRMonoPhase4.wav")
	tu.AssertNErr(err)
	tu.AssertEq(sig2.Wav.NumChannels, 2, t)
	tu.AssertEq(len(sig2.GetSamples(Left)), len(sig2.GetSamples(Right)), t)
	tu.AssertVNErr(sig2.MonoSpectrum(Left))

	specL, err := sig.MonoSpectrum(Left)
	tu.AssertNErr(err)
	tu.AssertNErr(DrawSpectrum(specL, "spectrum", true))

	filterSignal, err := sig.BandpassAtBaseFrequency()
	tu.AssertNErr(err)

	peakToPeakSignal, _, err := filterSignal.GetPeakToPeakBaseFilteredSignal()
	tu.AssertNErr(err)

	signals := make([]Signal, 3)
	signals[0] = *sig
	signals[1] = *filterSignal
	signals[2] = *peakToPeakSignal

}

func TestFreq440(t *testing.T) {
	sig, err := NewSignalFromWav("samples/stecker/Ma_57_st.wav")
	tu.AssertNErr(err)
	exactFrequency, err := sig.calcExactBaseFrequency()
	tu.AssertNErr(err)
	tu.AssertEq(exactFrequency, 440.19740294541276, t)
}

func TestFreq659(t *testing.T) {
	sig, err := NewSignalFromWav("samples/stecker/Ma_64_st.wav")
	tu.AssertNErr(err)
	exactFrequency, err := sig.calcExactBaseFrequency()
	tu.AssertNErr(err)
	tu.AssertEq(exactFrequency, 659.4640706543557, t)
}
