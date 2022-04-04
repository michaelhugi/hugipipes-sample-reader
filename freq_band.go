package hugipipes_sample

var FreqBand0 FreqBand = newFreqBand(16.35, 32.70)
var FreqBand1 FreqBand = newFreqBand(32.70, 65.41)
var FreqBand2 FreqBand = newFreqBand(65.41, 130.81)
var FreqBand3 FreqBand = newFreqBand(130.81, 261.63)
var FreqBand4 FreqBand = newFreqBand(261.63, 523.25)
var FreqBand5 FreqBand = newFreqBand(523.25, 1046.50)
var FreqBand6 FreqBand = newFreqBand(1046.50, 2093.00)
var FreqBand7 FreqBand = newFreqBand(2093.00, 4186.01)
var FreqBand9 FreqBand = newFreqBand(4186.01, 8372.02)

var FreqBand01 FreqBand = newFreqBand(29.14, 38.89)
var FreqBand12 FreqBand = newFreqBand(58.2, 77.78)
var FreqBand23 FreqBand = newFreqBand(116.54, 155.56)
var FreqBand34 FreqBand = newFreqBand(233.08, 311.13)
var FreqBand45 FreqBand = newFreqBand(466.16, 622.25)
var FreqBand56 FreqBand = newFreqBand(932.33, 1244.51)
var FreqBand67 FreqBand = newFreqBand(1864.66, 2489.02)
var FreqBand79 FreqBand = newFreqBand(3729.31, 4978.03)

type FreqBand struct {
	lowFreq  float64
	highFreq float64
}

func newFreqBand(lowFreq float64, highFreq float64) FreqBand {
	return FreqBand{
		lowFreq:  lowFreq,
		highFreq: highFreq,
	}
}
