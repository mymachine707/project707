package bigint

import "project/clean"

type Bigint struct {
	Value string
}

func NewInt(num string) (Bigint, error) {

	return Bigint{clean.Clean(num)}, nil
}

/*
func (a *Bigint) Set {

}*/
