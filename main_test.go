package main

import "testing"

func TestNumToBin(t *testing.T) {
	for _, tc := range []struct {
		num  int64
		want string
	}{
		{0, "0"},
		{1, "1"},
		{2, "10"},
		{3, "11"},
		{4, "100"},
		{5, "101"},
		{6, "110"},
		{7, "111"},
		{8, "1000"},
		{9, "1001"},
		{16, "10000"},
		{-1, "1111111111111111111111111111111111111111111111111111111111111111"},
		{1000_0000_0000_0000, "11100011010111111010100100110001101000000000000000"},
		{-1000_0000_0000_0001, "1111111111111100011100101000000101011011001110010111111111111111"},
	} {
		t.Run(tc.want, func(t *testing.T) {
			if got := NumToBin(tc.num); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestBinNumSum(t *testing.T) {
	for _, tc := range []struct {
		a, b string
		want string
	}{
		{"1", "11", "100"},
		{"11001" /* 25 */, "1111111111111111111111111111111111111111111111111111111111110010" /* -14 */, "1011" /* 11 */},
		{"1111111111111111111111111111111111111111111111111111111111110110" /* -10 */, "111" /* 7 */, "1111111111111111111111111111111111111111111111111111111111111101" /* -3 */},
		{"11" /* 3 */, "1111111111111111111111111111111111111111111111111111111111111101" /* -3 */, "0" /* 0 */},
	} {
		t.Run(tc.want, func(t *testing.T) {
			if got := BinNumSum(tc.a, tc.b); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
