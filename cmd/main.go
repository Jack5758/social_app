package main

import (
	"github.com/Jack5758/di"
)

func main() {

	err := di.RunApp()

	if err != nil {
		panic(err)
	}
}
