package models

type Input struct {
	Bytes bool
	Chars bool
	Lines bool
	Words bool
	Files []string
}

type Output struct {
	Bytes int
	Chars int
	Lines int
	Words int
}

type FileOp struct {
	File   string
	Output Output
}

// returns a new Input
func NewInput(usedFlags, files []string) Input {

	if len(usedFlags) == 0 {
		// 	Return Input with all as true
		return Input{true, true, true, true, files}
	}

	input := Input{}
	for _, fn := range usedFlags {
		if fn == "c" || fn == "bytes" {
			input.Bytes = true
		}

		if fn == "m" || fn == "chars" {
			input.Chars = true
		}

		if fn == "l" || fn == "lines" {
			input.Lines = true
		}

		if fn == "w" || fn == "words" {
			input.Words = true
		}
	}
	input.Files = files
	return input
}

// Add file info/counts,
// field :  one of bytes,chars,lines,ml,words
func (o *Output) AddCount(field string, n int) {

	switch field {
	case "bytes":
		o.Bytes += n
		return

	case "chars":
		o.Chars += n
		return

	case "lines":
		o.Lines += n
		return

	case "words":
		o.Words += n
		return

	}

}
