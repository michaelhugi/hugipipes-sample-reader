package hugipipes_sample

type Octave struct {
	low  float64
	high float64
	nr   uint8
}

const startFirstOctave = 15.8925

var Octave0 = newOctave(0, startFirstOctave, startFirstOctave*2)
var Octave1 = newOctave(1, startFirstOctave*2, startFirstOctave*4)
var Octave2 = newOctave(2, startFirstOctave*4, startFirstOctave*8)
var Octave3 = newOctave(3, startFirstOctave*8, startFirstOctave*16)
var Octave4 = newOctave(4, startFirstOctave*16, startFirstOctave*32)
var Octave5 = newOctave(5, startFirstOctave*32, startFirstOctave*64)
var Octave6 = newOctave(6, startFirstOctave*64, startFirstOctave*128)
var Octave7 = newOctave(7, startFirstOctave*128, startFirstOctave*256)
var Octave8 = newOctave(8, startFirstOctave*256, startFirstOctave*512)
var Octave9 = newOctave(9, startFirstOctave*512, startFirstOctave*1024)

func newOctave(nr uint8, low float64, high float64) *Octave {
	return &Octave{
		low:  low,
		high: high,
		nr:   nr,
	}
}
