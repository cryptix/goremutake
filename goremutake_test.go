package goremutake

import (
	"fmt"
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
	{128 << 7, "bebaba"},
	{(128 << 7) - 1, "tretre"},
	{128 << 14, "bebababa"},
	{(128 << 14) - 1, "tretretre"},
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
	fmt.Println(Encode(10610353957))
	fmt.Println(Encode(5778515749))

	// Output: koremutake
	// goremutake
}

func TestDecodeValidInput(t *testing.T) {
	for _, tcase := range testCases {
		actual, err := Decode(tcase.EncString)
		if err != nil {
			t.Errorf("'%s' should be %d\n Should not return an error:%v", tcase.EncString, tcase.Number, err)
		}
		if actual != tcase.Number {
			t.Errorf("'%s' should be %d\nActual:%d Expected:%d", tcase.EncString, tcase.Number, actual)
		}
	}
}

func TestDecodeLength(t *testing.T) {
	var lengthCases = []string{"", "a", "foo"}
	for _, tcase := range lengthCases {
		_, err := Decode(tcase)
		if err.Error() != ErrorInputLength {
			t.Errorf("Should return an error for input length. '%s'", tcase)
		}
	}
}

func TestDecodeInvalidSyllable(t *testing.T) {
	var invalidCases = []string{"xyz", "qy"}
	for _, tcase := range invalidCases {
		_, err := Decode(tcase)
		if err.Error() != ErrorInputSyllable {
			t.Errorf("Should return an error for invalid syllable. '%s'\nError:%v", tcase, err)
		}
	}
}

func ExampleDecode() {
	fmt.Println(Decode("koremutake"))
	fmt.Println(Decode("goremutake"))

	// Output: 10610353957 <nil>
	// 5778515749 <nil>
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Decode("goremutake")
	}
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encode(5778515749)
	}
}
