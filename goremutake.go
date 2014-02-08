/*
Convert to and from Koremutake Memorable Random Strings.

Koremutake is a 128bit MeRS encoding algorithm that can convert any large,
unsigned number into a memorable sequence of phonetically unique syllables.

See http://shorl.com/koremutake.php for details

This implementation is based on Patrick Schork's implementation in Python.
*/
package goremutake

import (
	"errors"
)

// known errors
const (
	ErrorInputLength   = "Invalid length of input string"
	ErrorInputSyllable = "Invalid syllable in input string"
)

// number of syllables
const syllables = 128

var phonemes = []string{
	"ba", "be", "bi", "bo", "bu", "by", "da", "de", "di", "do", "du", "dy", "fa",
	"fe", "fi", "fo", "fu", "fy", "ga", "ge", "gi", "go", "gu", "gy", "ha", "he",
	"hi", "ho", "hu", "hy", "ja", "je", "ji", "jo", "ju", "jy", "ka", "ke", "ki",
	"ko", "ku", "ky", "la", "le", "li", "lo", "lu", "ly", "ma", "me", "mi", "mo",
	"mu", "my", "na", "ne", "ni", "no", "nu", "ny", "pa", "pe", "pi", "po", "pu",
	"py", "ra", "re", "ri", "ro", "ru", "ry", "sa", "se", "si", "so", "su", "sy",
	"ta", "te", "ti", "to", "tu", "ty", "va", "ve", "vi", "vo", "vu", "vy", "bra",
	"bre", "bri", "bro", "bru", "bry", "dra", "dre", "dri", "dro", "dru", "dry",
	"fra", "fre", "fri", "fro", "fru", "fry", "gra", "gre", "gri", "gro", "gru",
	"gry", "pra", "pre", "pri", "pro", "pru", "pry", "sta", "ste", "sti", "sto",
	"stu", "sty", "tra", "tre",
}

// Encode unsigned integer value to Koremutake string
func Encode(value uint) string {
	if value == 0 {
		return phonemes[0]
	}

	var key string
	for value > 0 {
		key = phonemes[value%syllables] + key
		value /= syllables
	}

	return key
}

//Decode Koremutake string to unsigned integer value
func Decode(input string) (value uint, err error) {

	// construct map of indexes to ease lookup
	phonemesMap := make(map[string]uint, 128)
	for k, v := range phonemes {
		phonemesMap[v] = uint(k)
	}

	if len(input) < 2 {
		return 0, errors.New(ErrorInputLength)
	}

	// iterate over input and calculate value
	var (
		idx uint
		ok  bool
	)
	for input != "" {
		// fmt.Println(input, len(input))
		if len(input) < 2 {
			return 0, errors.New(ErrorInputLength)
		}

		if idx, ok = phonemesMap[input[:2]]; ok {
			input = input[2:]
			value = value*syllables + idx
			continue
		}

		if len(input) >= 3 {
			if idx, ok = phonemesMap[input[:3]]; ok {
				input = input[3:]
				value = value*syllables + idx
				continue
			}
		}

		return 0, errors.New(ErrorInputSyllable)

	}
	return
}
