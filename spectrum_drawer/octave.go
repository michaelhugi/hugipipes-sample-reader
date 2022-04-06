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

const Octave8Factor = 256.0

const C0Freq = 4186.01 / Octave8Factor
const Cis0Freq = 4434.92 / Octave8Factor
const D0Freq = 4698.63 / Octave8Factor
const Dis0Freq = 4978.03 / Octave8Factor
const E0Freq = 5274.04 / Octave8Factor
const F0Freq = 5587.65 / Octave8Factor
const Fis0Freq = 5919.91 / Octave8Factor
const G0Freq = 6271.93 / Octave8Factor
const Gis0Freq = 6644.88 / Octave8Factor
const A0Freq = 7040.00 / Octave8Factor
const Ais0Freq = 7458.62 / Octave8Factor
const H0Freq = 7902.13 / Octave8Factor

var Octave1 = newOctave(0)
var Octave2 = newOctave(1)
var Octave3 = newOctave(2)
var Octave4 = newOctave(3)
var Octave5 = newOctave(4)
var Octave6 = newOctave(5)
var Octave7 = newOctave(6)
var Octave8 = newOctave(7)
var Octave9 = newOctave(8)

func newOctave(nr uint8) *Octave {
	low := C0Freq * math.Pow(2, float64(nr))
	high := C0Freq * math.Pow(2, float64(nr+1))
	return &Octave{
		low:  low,
		high: high,
		nr:   nr,
	}
}

func (o *Octave) getNotes() []Note {
	if o.nr < 3 {
		return make([]Note, 0)
	}
	notes := make([]Note, 11)
	//notes[5] = newNote("C", o.nr, 30.87, 4)
	notes[0] = newNote("C#", o.nr, Cis0Freq, 24, false)
	notes[1] = newNote("D", o.nr, D0Freq, 25, true)
	notes[2] = newNote("D#", o.nr, Dis0Freq, 26, false)
	notes[3] = newNote("E", o.nr, E0Freq, 27, true)
	notes[4] = newNote("F", o.nr, F0Freq, 28, true)
	notes[5] = newNote("F#", o.nr, Fis0Freq, 29, false)
	notes[6] = newNote("G", o.nr, G0Freq, 30, true)
	notes[7] = newNote("G#", o.nr, Gis0Freq, 31, false)
	notes[8] = newNote("A", o.nr, A0Freq, 32, true)
	notes[9] = newNote("A#", o.nr, Ais0Freq, 33, false)
	notes[10] = newNote("H", o.nr, H0Freq, 34, true)

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
