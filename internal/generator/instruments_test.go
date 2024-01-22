package generator_test

import (
	"fmt"
	"testing"
)

func TestGeneratorInstruments(t *testing.T) {

	gen, err := mkTestGenerator()
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(gen.Instruments())
	}
}
