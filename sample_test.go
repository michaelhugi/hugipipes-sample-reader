package hugipipes_sample

import (
	"fmt"
	tu "github.com/informaticon/lib.go.base.test-utils"
	wv "github.com/mjibson/go-dsp/wav"
	"os"
	"testing"
	"time"
)

func TestCalculateSample(t *testing.T) {
	file, err := os.Open("samples/stecker/Ma2_55_st.wav")
	tu.AssertNErr(err)
	f, err := LoadSample(file)
	tu.AssertNErr(err)

	tu.AssertNErr(file.Close())

	file2, err := os.Open("samples/stecker/Ma2_55_st.wav")
	tu.AssertNErr(err)
	defer func() {
		tu.AssertNErr(file2.Close())
	}()

	wav, err := wv.New(file2)
	tu.AssertNErr(err)

	fmt.Printf("%+v\n", f.Spectrum.Points[0].String())
	fmt.Printf("%+v\n", f.Spectrum.Points[len(f.Spectrum.Points)-1].String())
	tu.AssertEq(f.Spectrum.Points[0].LowFrequency, 0, t)
	tu.AssertEq(f.Spectrum.Points[0].CenterFrequency, 0, t)
	tu.AssertEq(f.Spectrum.Points[0].HighFrequency, 0.22050441008820176, t)

	tu.AssertEq(f.Spectrum.Points[1].LowFrequency, 0.22050441008820176, t)
	tu.AssertEq(f.Spectrum.Points[1].CenterFrequency, 0.4410088201764035, t)
	tu.AssertEq(f.Spectrum.Points[1].HighFrequency, 0.6615132302646053, t)

	tu.AssertEq(f.Spectrum.Points[len(f.Spectrum.Points)-1].LowFrequency, 22049.77949558991, t)
	tu.AssertEq(f.Spectrum.Points[len(f.Spectrum.Points)-1].CenterFrequency, 22050, t)
	tu.AssertEq(f.Spectrum.Points[len(f.Spectrum.Points)-1].HighFrequency, 22050, t)

	tu.AssertEq(f.TimeSpectrum.Items[0].Time, time.Duration(0), t)

	timeDelta := wav.Duration.Nanoseconds() / int64(wav.Samples-1)
	timeLastItemsNs := int64(wav.Samples-fftWindowSize-1) * timeDelta
	timeLastItem := time.Duration(timeLastItemsNs) * time.Nanosecond

	tu.AssertEq(f.TimeSpectrum.Items[len(f.TimeSpectrum.Items)-1].Time, timeLastItem, t)

}
