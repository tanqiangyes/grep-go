package reader

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

type StdReader struct {
	finder []Finder
	reader io.Reader

	Output []MatchRes
	Error  error
}

func (s *StdReader) Result() []MatchRes {
	return s.Output
}

func (s *StdReader) Run() {
	br := bufio.NewReader(s.reader)
	var line int64 = 1
	var match MatchRes
	match.Filename = "stdin"
	for {
		readString, err := br.ReadString('\n')
		if err != nil {
			// EOF, we should break this loop, and return data.
			if errors.Is(err, io.EOF) {
				break
			}
			// something wrong, write the error and return.
			s.Error = err
			return
		}

		// do some string processing
		readString = strings.TrimSpace(readString)
		readString = strings.Trim(readString, " ")
		readString = strings.Trim(readString, "\n")

		if s.find(readString) {
			// we found, add it into match.
			match.Lines = append(match.Lines, line)
			match.MatchString = append(match.MatchString, readString)
		}
		line++
	}
	s.Output = []MatchRes{match}
	return
}

func (s *StdReader) find(str string) bool {
	for _, finder := range s.finder {
		if _, ok := finder.Find(str); ok {
			return ok
		}
	}
	return false
}

func (s *StdReader) Close() {
	return
}

func (s *StdReader) IsError() error {
	return s.Error
}

func NewStdReader(reader io.Reader, finder []Finder) (Reader, error) {
	read := bufio.NewReader(reader)
	return &StdReader{
		finder: finder,
		reader: read,
		Output: nil,
	}, nil
}
