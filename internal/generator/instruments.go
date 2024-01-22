package generator

import (
	"math/rand"
	"strings"
)

const (
	keyInstruments = "{{.INSTRUMENTS}}"
)

func (g *Generator) Instruments() string {

	var (
		used       = map[int]struct{}{}
		l          = len(g.dict.Instruments)
		instrCount = rand.Intn(3) + 1
		text       = randomElem(g.dict.TextInstruments)
		result     = ""
	)

	if instrCount > l {
		instrCount = l
	}

	if text == "" {
		text = "Try to use this instruments: <i><b>{{.INSTRUMENTS}}</b></i>"
	}

	for i := 0; i < instrCount; {

		v := rand.Intn(l)
		if _, ok := used[v]; ok {
			continue
		}

		result = result + "\n• " + g.dict.Instruments[v]
		used[v] = struct{}{}

		i++
	}

	return strings.ReplaceAll(text, keyInstruments, result)
}
