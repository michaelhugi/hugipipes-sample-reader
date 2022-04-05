package dspttt

import (
	"github.com/eripe970/go-dsp-utils"
	test_utils "github.com/informaticon/lib.go.base.test-utils"
)

func foo() {
	s, err := dsp.ReadSignalFile("asdfa", 3)
	test_utils.AssertNErr(err)
	s.FrequencySpectrum()
	//fft.FFTReal()
}
