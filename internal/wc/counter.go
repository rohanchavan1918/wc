package wc

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"

	"github.com/rohanchavan1918/wc/internal/models"
	"github.com/rohanchavan1918/wc/internal/utils"
)

// Main entry point of the word counter
func Process(input models.Input, usedFlags []string) {
	// Check if the input is coming from files or stdIn
	if len(input.Files) > 0 {
		ProcessFiles(input, usedFlags)
	} else {
		ProcessStdin(input, usedFlags)
	}

}

func ProcessFiles(input models.Input, usedFlags []string) error {

	fileResults := []models.FileOp{}
	totalResult := models.Output{}

	for _, file := range input.Files {
		fileResult := models.FileOp{File: file}
		// wc does not give line count, but newLine count thus we need to initialize Lines to -1
		op := models.Output{Lines: -1}
		f, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
		if err != nil {
			fmt.Println("Failed to process file ", file)
			fmt.Println("Error  ", err.Error())
			continue
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			if err := ProcessLine(input, scanner.Bytes(), &op); err == nil {
				op.AddCount("lines", 1)
			}
		}

		if err := scanner.Err(); err != nil {
			return nil
		}

		fileResult.Output = op
		fileResults = append(fileResults, fileResult)
		totalResult.AddCount("bytes", op.Bytes)
		totalResult.AddCount("chars", op.Chars)
		totalResult.AddCount("lines", op.Lines)
		totalResult.AddCount("words", op.Words)
	}

	for _, op := range fileResults {
		utils.PrintResult(usedFlags, op.Output, op.File)
	}
	if len(fileResults) > 1 {
		utils.PrintResult(usedFlags, totalResult, "total")
	}

	return nil
}

func ProcessStdin(input models.Input, usedFlags []string) error {
	scanner := bufio.NewScanner(os.Stdin)
	output := models.Output{Lines: -1}

	for scanner.Scan() {
		if err := ProcessLine(input, scanner.Bytes(), &output); err == nil {
			output.AddCount("lines", 1)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	utils.PrintResult(usedFlags, output, "")
	return nil
}

func ProcessLine(ip models.Input, line []byte, op *models.Output) error {
	if ip.Bytes {
		op.AddCount("bytes", len(line))
	}

	if ip.Chars {
		op.AddCount("chars", utf8.RuneCount(line))
	}

	if ip.Words {
		// Calculate words
		wSlc := string(line)
		words := 0
		wordStarted := false
		for _, s := range wSlc {
			if string(s) != " " && !wordStarted {
				wordStarted = true
				words++
			} else if string(s) == " " && wordStarted {
				wordStarted = false
			}
		}

		op.AddCount("words", words)
	}

	return nil
}
