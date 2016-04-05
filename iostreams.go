package iostreams

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var ThroughMode bool = false

func ProcessStdin(process func(row []byte) error) error {

	var input []byte = nil

	stat, err := os.Stdin.Stat()
	if err != nil {
		return err
	}

	if (stat.Mode() & os.ModeCharDevice) == 0 {
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
				if ThroughMode {
					fmt.Println(string(input))
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
