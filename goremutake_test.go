package goremutake

import (
	"fmt"
	"math"
	"testing"
)

var testCases = []struct {
	Number    uint
	EncString string
}{
	{0, "ba"},
	{127, "tre"},
	{128, "beba"},
	{256, "biba"},
	{uint(math.Pow(128, 2)), "bebaba"},
	{uint(math.Pow(128, 2)) - 1, "tretre"},
	{uint(math.Pow(128, 3)), "bebababa"},
	{uint(math.Pow(128, 3)) - 1, "tretretre"},
	{10610353957, "koremutake"},
	{4398046511103, "tretretretretretre"},
}

func TestEncode(t *testing.T) {
	for _, tcase := range testCases {
		actual := Encode(tcase.Number)
		if actual != tcase.EncString {
			t.Errorf("%d should be '%s' - Actual:%s", tcase.Number, tcase.EncString, actual)
		}
	}
}

func ExampleEncode() {
	fmt.Println(Encode(12345))
	// Output: drano
}

func TestDecode(t *testing.T) {
	for _, tcase := range testCases {
		actual := Decode(tcase.EncString)
		if actual != tcase.Number {
			t.Errorf("'%s' should be %d - Actual:%d", tcase.EncString, tcase.Number, actual)
		}
	}
}

func ExampleDecode() {
	fmt.Println(Decode("babebibobu"))
	// Output: 2130308
}
