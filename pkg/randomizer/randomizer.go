package randomizer

import (
	"math/rand"
	"time"
)

type Randomizer struct {
	Code        string
	ValueInt    int
	ValueString string
}

func NewRandomizer(code string) *Randomizer {
	return &Randomizer{
		Code: code,
	}
}

func (r *Randomizer) Generate() *Randomizer {
	rand.NewSource(time.Now().UnixNano())
	r.ValueInt = rand.Intn(19)
	return r

}
