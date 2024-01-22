package generator

import "strings"

const (
	keyGenre = "{{.GENRE}}"
)

func (g *Generator) Genre() string {
	return strings.ReplaceAll(g.genreText(), keyGenre, g.genreKeys())
}

func (g *Generator) genreText() string {
	if t := randomElem(g.dict.TextGenre); t != "" {
		return t
	}
	return "Genre of track is {{.GENRE}}"
}

func (g *Generator) genreKeys() string {

	var (
		base   = randomElem(g.dict.GenreBase)
		prefix = randomElem(g.dict.GenrePrefix)
		patch  = randomElem(g.dict.GenrePatch)
	)

	if base != prefix {
		base = prefix + " " + base
	}

	if prefix != patch {
		base = patch + " " + base
	}

	if base == "" {
		return "Russian Rock"
	}

	return base
}
