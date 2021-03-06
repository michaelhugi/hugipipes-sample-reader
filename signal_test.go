package hugipipes_sample

import (
	tu "github.com/informaticon/lib.go.base.test-utils"
	"image"
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
	//temp := mn.NewMTemperamentEqual(440)
	//lowNote := temp.Octave(3).Note(mn.C)
	//highNote := temp.Octave(6).Note(mn.C)
	tu.AssertNErr(err)

	filterSignal, err := sig.BandpassAtBaseFrequency()
	tu.AssertNErr(err)

	peakToPeakSignal, _, err := filterSignal.GetPeakToPeakBaseFilteredSignal()
	tu.AssertNErr(err)

	drawer := NewPngDrawer().PlotWidth(int(sig.SampleCount))
	drawer.AddSpectrum(NewSpectrumPngDrawer(specL, "Spectrum"))
	wv1 := NewWavePngDrawer("Signal unfiltered / filtered", NewWavePngDrawerItem(sig, Left, blue))
	wv1.AddWave(NewWavePngDrawerItem(filterSignal, Left, image.White.C))
	drawer.AddWave(wv1)
	drawer.AddWave(NewWavePngDrawer("Peak to peak filtered", NewWavePngDrawerItem(peakToPeakSignal, Left, image.White.C)))
	tu.AssertNErr(drawer.Draw("testspectrum"))
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
