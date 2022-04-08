package musical_note

import (
	"errors"
	"fmt"
	"math"
)

type MusicalOctave struct {
	lowCFreq     float64
	nextCFreq    float64
	octaveNumber int
}

var octave0 = newMusicalOctave(0)
var octave1 = newMusicalOctave(1)
var octave2 = newMusicalOctave(2)
var octave3 = newMusicalOctave(3)
var octave4 = newMusicalOctave(4)
var octave5 = newMusicalOctave(5)
var octave6 = newMusicalOctave(6)
var octave7 = newMusicalOctave(7)
var octave8 = newMusicalOctave(8)

func newMusicalOctave(octaveNumber int) *MusicalOctave {
	low := c0Freq * math.Pow(2, float64(octaveNumber))
	high := c0Freq * math.Pow(2, float64(octaveNumber+1))
	return &MusicalOctave{
		lowCFreq:     low,
		nextCFreq:    high,
		octaveNumber: octaveNumber,
	}
}

func musicalOctaveFromFrequency(freq float64) (*MusicalOctave, error) {
	if octave0.FreqIsInOctave(freq) {
		return newMusicalOctave(octave0.octaveNumber), nil
	}
	if octave1.FreqIsInOctave(freq) {
		return newMusicalOctave(octave1.octaveNumber), nil
	}
	if octave2.FreqIsInOctave(freq) {
		return newMusicalOctave(octave2.octaveNumber), nil
	}
	if octave3.FreqIsInOctave(freq) {
		return newMusicalOctave(octave3.octaveNumber), nil
	}
	if octave4.FreqIsInOctave(freq) {
		return newMusicalOctave(octave4.octaveNumber), nil
	}
	if octave5.FreqIsInOctave(freq) {
		return newMusicalOctave(octave5.octaveNumber), nil
	}
	if octave6.FreqIsInOctave(freq) {
		return newMusicalOctave(octave6.octaveNumber), nil
	}
	if octave7.FreqIsInOctave(freq) {
		return newMusicalOctave(octave7.octaveNumber), nil
	}
	if octave8.FreqIsInOctave(freq) {
		return newMusicalOctave(octave8.octaveNumber), nil
	}
	return nil, errors.New(fmt.Sprintf("Freq %fHz can not be assigned to octave. Freq must be within %fHz and %fHz", freq, octave0.lowCFreq, octave1.nextCFreq))
}

func (s *MusicalOctave) FreqIsInOctave(freq float64) bool {
	return s.lowCFreq <= freq && s.nextCFreq > freq
}
