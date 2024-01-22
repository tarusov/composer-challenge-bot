package generator_test

import (
	"fmt"
	"testing"
)

func TestGeneratorHello(t *testing.T) {

	const (
		fn = "Franz"
		ln = "Kafka"
		un = "PussyDestroyer777"
	)

	gen, err := mkTestGenerator()
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(gen.Hello(fn, ln, un))
	}
}
