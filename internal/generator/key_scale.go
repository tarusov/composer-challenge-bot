package generator

import "strings"

const (
	keyKey   = `{{.KEY}}`
	keyScale = `{{.SCALE}}`
)

// KeyScale
func (g *Generator) KeyScale() string {

	var (
		text  = g.ksText()
		key   = g.key()
		scale = g.scale()
	)

	text = strings.ReplaceAll(text, keyKey, key)
	text = strings.ReplaceAll(text, keyScale, scale)

	return text
}

func (g *Generator) ksText() string {
	if t := randomElem(g.dict.TextKeyScale); t != "" {
		return t
	}
	return "Write song in {{.KEY}} {{.SCALE}}"
}

func (g *Generator) key() string {
	if k := randomElem(g.dict.Keys); k != "" {
		return k
	}
	return "C"
}

func (g *Generator) scale() string {
	if s := randomElem(g.dict.Scales); s != "" {
		return s
	}
	return "Major"
}
