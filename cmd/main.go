package main

import (
	"flag"

	"github.com/rohanchavan1918/wc/internal/models"
	"github.com/rohanchavan1918/wc/internal/utils"
	"github.com/rohanchavan1918/wc/internal/wc"
)

func main() {
	var c, m, l, w bool
	var bytes, chars, lines, words bool

	flag.BoolVar(&c, "c", false, "print the byte counts")
	flag.BoolVar(&bytes, "bytes", false, "print the byte counts")

	flag.BoolVar(&m, "m", false, "print the character counts")
	flag.BoolVar(&chars, "chars", false, "print the character counts")

	flag.BoolVar(&l, "l", false, "print the character counts")
	flag.BoolVar(&lines, "lines", false, "print the character counts")

	flag.BoolVar(&w, "w", false, "print the character counts")
	flag.BoolVar(&words, "words", false, "print the character counts")

	flag.Parse()

	availableCommands := []string{"l", "lines", "w", "words", "c", "bytes", "m", "chars", "L", "mll"}
	usedFlags := utils.FlagUsed(availableCommands...)
	input := models.NewInput(usedFlags, flag.Args())
	wc.Process(input, usedFlags)
}
