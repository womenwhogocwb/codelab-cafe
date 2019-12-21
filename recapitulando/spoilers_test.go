package main

import "testing"

func TestSomaSlices(t *testing.T) {
	t.Run("soma slices de 5 elementos", func(t *testing.T) {
		sliceDoTeste := []int{1, 2, 3, 4, 5}
		got := SomaSlices(sliceDoTeste)
		want := 15

		if got != want {
			t.Errorf("got %d; want %d", got, want)
		}
	})
}