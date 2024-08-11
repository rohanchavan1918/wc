package utils

import (
	"flag"
	"fmt"

	"github.com/rohanchavan1918/wc/internal/models"
)

// Checks if a flag is passed or not
func FlagUsed(s ...string) []string {
	usedFlags := []string{}

	for _, fn := range s {
		flag.Visit(func(f *flag.Flag) {
			if f.Name == fn {
				usedFlags = append(usedFlags, fn)
				return
			}
		})
	}

	return usedFlags
}

func PrintResult(flags []string, o models.Output, fn string) {
	if len(flags) == 0 {
		fmt.Printf("%d %d %d %s\n", o.Lines, o.Words, o.Bytes, fn)
		return
	}

	op := ""
	for _, f := range flags {
		if f == "l" || f == "lines" {
			op = op + fmt.Sprintf("%d ", o.Lines)
		}

		if f == "w" || f == "words" {
			op = op + fmt.Sprintf("%d ", o.Words)
		}

		if f == "c" || f == "bytes" {
			op = op + fmt.Sprintf("%d ", o.Bytes)
		}

		if f == "m" || f == "chars" {
			op = op + fmt.Sprintf("%d ", o.Chars)
		}

	}

	op = op + fn
	fmt.Println(op)

}

func PrintFilesResult(flags []string, o []models.FileOp) {
	for _, fn := range o {
		PrintResult(flags, fn.Output, fn.File)
	}
}
