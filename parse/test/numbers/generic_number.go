package numbers

import "github.com/metabition/genny/generic"

type NumberType generic.Number

func NumberTypeMax(a, b NumberType) NumberType {
	if a > b {
		return a
	}
	return b
}