package main

import "testing"

func TestCalibration01(t *testing.T) {

	got := getCalibrationValue("1abc2")
	want := 12

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCalibration02(t *testing.T) {

	got := getCalibrationValue("a1b2c3d4e5f")
	want := 15

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCalibration03(t *testing.T) {

	got := getCalibrationValue("aa45sd7a5123sd")
	want := 43

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestReplacement01(t *testing.T) {

	got := getCalibrationValueWithReplace("two1nine", true)
	want := 29

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestReplacement02(t *testing.T) {

	got := getCalibrationValueWithReplace("abcone2threexyz", true)
	want := 13

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestReplacement03(t *testing.T) {

	got := getCalibrationValueWithReplace("eightwothree", true)
	want := 83

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}