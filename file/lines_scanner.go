package file

import (
	"bufio"
	"io"
	"strings"
)

// LinesScanner provides methods to read the entire file line-by-line, omitting newline character
type LinesScanner interface {
	Next() bool
	Get() (string, error)
}

type linesIterator struct {
	rd         io.Reader
	reader     *bufio.Reader
	err        error
	line       string
	isFinished bool
}

func (iterator *linesIterator) Next() bool {
	if iterator.err != nil || iterator.isFinished {
		return false
	}

	line, err := iterator.ReadEntireLine()

	if err == io.EOF || err == nil {
		iterator.line = line
	}

	if err == io.EOF {
		iterator.isFinished = true
	}

	if err != nil && err != io.EOF {
		iterator.err = err
	}

	return true
}

func (iterator *linesIterator) Get() (string, error) {
	return iterator.line, iterator.err
}

// NewLinesScanner starts a new `LinesScanner` for the `io.Reader`
func NewLinesScanner(rd io.Reader) LinesScanner {
	iterator := &linesIterator{}
	iterator.rd = rd
	iterator.reader = bufio.NewReader(iterator.rd)

	return iterator
}

func (iterator *linesIterator) ReadEntireLine() (string, error) {
	reader := iterator.reader
	var entireLine strings.Builder

	for {
		line, isPrefix, err := reader.ReadLine()
		if err != nil && err != io.EOF {
			return "", err
		}

		_, bufErr := entireLine.Write(line)
		if bufErr != nil {
			return "", bufErr
		}

		if isPrefix == false && (err == nil || err == io.EOF) {
			return entireLine.String(), err
		}
	}
}
