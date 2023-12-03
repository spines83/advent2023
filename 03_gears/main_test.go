package main

import "testing"

func TestScan01(t *testing.T) {

	schematic := [][]rune{
		{'.', '.', '.', '.', '.'},
		{'.', '1', '2', '3', '.'},
		{'.', '.', '&', '.', '.'},
	}

	got, _ := CalculateSchematicSum(schematic)
	want := 123

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestScan02(t *testing.T) {

	schematic := [][]rune{
		{'.', '1', '2', '3', '.'},
		{'&', '.', '.', '.', '.'},
	}

	got, _ := CalculateSchematicSum(schematic)
	want := 123

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestScan03(t *testing.T) {

	schematic := [][]rune{
		{'&', '.', '.', '.', '.'},
		{'.', '1', '2', '3', '.'},
	}

	got, _ := CalculateSchematicSum(schematic)
	want := 123

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestScan04(t *testing.T) {

	schematic := [][]rune{
		{'.', '.', '&', '.', '.'},
		{'.', '1', '2', '3', '4'},
		{'.', '.', '.', '.', '.'},
	}

	got, _ := CalculateSchematicSum(schematic)
	want := 1234

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestScan05(t *testing.T) {

	schematic := [][]rune{
		{'.', '.', '.', '.', '.'},
		{'0', '1', '2', '3', '&'},
		{'.', '.', '.', '.', '.'},
	}

	got, _ := CalculateSchematicSum(schematic)
	want := 123

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestScan06(t *testing.T) {

	schematic := [][]rune{
		{'.', '.', '.', '.', '.'},
		{'.', '.', '2', '3', '&'},
		{'.', '.', '.', '.', '.'},
	}

	got, _ := CalculateSchematicSum(schematic)
	want := 23

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestScan07(t *testing.T) {

	schematic := [][]rune{
		{'.', '.', '.', '.', '.'},
		{'&', '1', '2', '.', '.'},
		{'.', '.', '.', '.', '.'},
	}

	got, _ := CalculateSchematicSum(schematic)
	want := 12

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestScan08(t *testing.T) {

	schematic := [][]rune{
		{'.', '.', '.', '.', '.'},
		{'&', '.', '2', '1', '.'},
		{'.', '.', '.', '.', '.'},
	}

	got, _ := CalculateSchematicSum(schematic)
	want := 0

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestScan09(t *testing.T) {

	schematic := [][]rune{
		{'.', '.', '.', '.', '.'},
		{'.', '1', '2', '.', '&'},
		{'.', '.', '.', '.', '.'},
	}

	got, _ := CalculateSchematicSum(schematic)
	want := 0

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
func TestScan10(t *testing.T) {

	schematic := [][]rune{
		{'.', '.', '.', '.', '.'},
		{'.', '.', '3', '8', '8'},
		{'/', '.', '.', '.', '.'},
	}

	got, _ := CalculateSchematicSum(schematic)
	want := 0

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}