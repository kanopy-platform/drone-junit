package main

import (
	"github.com/kanopy-platform/drone-junit/internal/card"
	"github.com/kanopy-platform/drone-junit/internal/junit"
)

func main() {
	suites, err := junit.Read("/tmp/unit-tests.xml") // TODO customize path with arg
	if err != nil {
		panic(err)
	}

	if err := card.WriteCard(suites); err != nil {
		panic(err)
	}
}
