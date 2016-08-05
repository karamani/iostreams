package iostreams

import (
	"bufio"
	"io"
	"os"
)

type StdinProcessor struct{}

func NewStdinProcessor() *StdinProcessor {
	return &StdinProcessor{}
}

func (*StdinProcessor) Ready() bool {
	return StdinReady()
}

func (*StdinProcessor) Process(pfunc func(row []byte) error) error {
	return ProcessStdin(pfunc)
}

func stdinReady() (bool, error) {

	stat, err := os.Stdin.Stat()
	if err != nil {
		return false, err
	}

	ready := (stat.Mode() & os.ModeCharDevice) == 0
	return ready, nil
}

func StdinReady() bool {
	ready, err := stdinReady()
	return err == nil && ready
}

func ProcessStdin(process func(row []byte) error) error {

	var input []byte = nil

	stdinReady, err := stdinReady()
	if err != nil {
		return err
	}

	if stdinReady {
		reader := bufio.NewReader(os.Stdin)
		for {
			bytes, hasMoreInLine, err := reader.ReadLine()
			if err != nil {
				if err == io.EOF {
					return nil
				}
				return err
			}
			input = append(input, bytes...)
			if !hasMoreInLine {
				if err := process(input); err != nil {
					return err
				}
				input = nil
			}
		}
	}

	return nil
}

func ChanStdin(line chan []byte) error {

	process := func(row []byte) error {
		line <- row
		return nil
	}

	err := ProcessStdin(process)
	return err
}
