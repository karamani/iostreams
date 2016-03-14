package iostreams

import (
	"bufio"
	"io"
	"os"
)

func ReadStdin(process func(row []byte) error) error {

	var input []byte = nil

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
