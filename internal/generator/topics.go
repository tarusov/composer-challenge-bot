package generator

import (
	"math/rand"
	"strings"
)

const (
	keyTopics = `{{.TOPICS}}`
)

func (g *Generator) Topics() string {
	return strings.ReplaceAll(g.topicText(), keyTopics, g.topicWords())
}

func (g *Generator) topicText() string {
	if t := randomElem(g.dict.TextTopics); t != "" {
		return t
	}
	return "Topics of the day is: {{.TOPICS}}"
}

func (g *Generator) topicWords() string {

	words, err := g.rwAPI.Words(rand.Intn(2) + 3)
	if err != nil {
		return "peace, loving and understanding"
	}

	var result string
	for _, word := range words {
		result = result + "\n• " + word
	}

	return result
}

func (g *Generator) Tips() string {

	var (
		used      = map[int]struct{}{}
		l         = len(g.dict.Tips)
		tipsCount = rand.Intn(3) + 1
		result    = "Tips:"
	)

	if tipsCount > l {
		tipsCount = l
	}

	for i := 0; i < tipsCount; {

		v := rand.Intn(l)
		if _, ok := used[v]; ok {
			continue
		}

		result = result + "\n• " + g.dict.Tips[v]
		used[v] = struct{}{}

		i++
	}

	return result
}
