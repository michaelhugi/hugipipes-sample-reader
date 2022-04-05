package spectrum_drawer

import (
	"fmt"
	"math"
)

type Octave struct {
	low  float64
	high float64
	nr   uint8
}

func (o *Octave) drawFrequency() bool {
	return o.nr > 0
}

func (o *Octave) frequencyReadable() string {
	l := fmt.Sprint(math.Round(o.low*100) / 100)
	if o.nr <= 1 {
		return l
	}
	return fmt.Sprintf("%sHz", l)
}

const startFirstOctave = 32.70 //16.35

var Octave1 = newOctave(1, startFirstOctave, startFirstOctave*2)
var Octave2 = newOctave(2, startFirstOctave*2, startFirstOctave*4)
var Octave3 = newOctave(3, startFirstOctave*4, startFirstOctave*8)
var Octave4 = newOctave(4, startFirstOctave*8, startFirstOctave*16)
var Octave5 = newOctave(5, startFirstOctave*16, startFirstOctave*32)
var Octave6 = newOctave(6, startFirstOctave*32, startFirstOctave*64)
var Octave7 = newOctave(7, startFirstOctave*64, startFirstOctave*128)
var Octave8 = newOctave(8, startFirstOctave*128, startFirstOctave*256)
var Octave9 = newOctave(9, startFirstOctave*256, startFirstOctave*512)

func newOctave(nr uint8, low float64, high float64) *Octave {
	return &Octave{
		low:  low,
		high: high,
		nr:   nr,
	}
}

func (o *Octave) getNotes() []Note {
	if o.nr < 2 {
		return make([]Note, 0)
	}
	notes := make([]Note, 11)
	//notes[5] = newNote("C", o.nr, 30.87, 4)
	notes[0] = newNote("C#", o.nr, 34.65, 24, false)
	notes[1] = newNote("D", o.nr, 36.71, 25, true)
	notes[2] = newNote("D#", o.nr, 38.89, 26, false)
	notes[3] = newNote("E", o.nr, 41.20, 27, true)
	notes[4] = newNote("F", o.nr, 43.65, 28, true)
	notes[5] = newNote("F#", o.nr, 46.25, 29, false)
	notes[6] = newNote("G", o.nr, 49.00, 30, true)
	notes[7] = newNote("G#", o.nr, 51.91, 31, false)
	notes[8] = newNote("A", o.nr, 55.00, 32, true)
	notes[9] = newNote("A#", o.nr, 58.27, 33, false)
	notes[10] = newNote("H", o.nr, 61.74, 34, true)

	return notes
}

type Note struct {
	freq float64
	name string
	midi string
}

func newNote(baseName string, octave uint8, baseFreq float64, baseMidi int, showText bool) Note {
	factor := math.Pow(2, float64(octave))
	freq := baseFreq * factor
	midi := fmt.Sprint(baseMidi + int((octave-1)*12) + 1)
	name := fmt.Sprintf("%s%d", baseName, octave)
	if !showText {
		midi = ""
		name = ""
	}
	return Note{
		freq: freq,
		name: name,
		midi: midi,
	}
}
