package generator

import (
	"math/rand"
	"strings"
)

const (
	keyFirstName = `{{.FIRST_NAME}}`
	keyLastName  = `{{.LAST_NAME}}`
	keyUserName  = `{{.USER_NAME}}`

	defaultHello = "Hi!"
)

// Hello generates random greetings message.
func (g *Generator) Hello(fn, ln, un string) string {
	switch rand.Intn(3) {
	case 0:
		return g.helloFn(fn)
	case 1:
		return g.helloFnLn(fn, ln)
	case 2:
		return g.helloUn(un)
	default:
		return defaultHello
	}
}

func (g *Generator) helloFn(fn string) string {
	if hw := randomElem(g.dict.TextHelloFn); hw != "" {
		return strings.ReplaceAll(hw, keyFirstName, fn)
	}
	return defaultHello
}

func (g *Generator) helloFnLn(fn, ln string) string {
	if hw := randomElem(g.dict.TextHelloFnLn); hw != "" {
		hw = strings.ReplaceAll(hw, keyFirstName, fn)
		return strings.ReplaceAll(hw, keyLastName, ln)
	}
	return defaultHello
}

func (g *Generator) helloUn(un string) string {
	if hw := randomElem(g.dict.TextHelloUn); hw != "" {
		return strings.ReplaceAll(hw, keyUserName, un)
	}
	return defaultHello
}
