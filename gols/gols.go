package main

import (
	"os"

	"gols/functions"
)

func main() {
	args := os.Args[1:]

	useColor := functions.IsTerminal(os.Stdout)

	functions.SimpleLS(os.Stdout, args, useColor)
}
