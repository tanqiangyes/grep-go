package reader

import (
	"bufio"
	"errors"
	"io"
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
		line++
		if err != nil {
			// EOF, we should break this loop, and return data.
			if errors.Is(err, io.EOF) {
				break
			}
			// something wrong, write the error and return.
			s.Error = err
			return
		}
		if s.find(readString) {
			// we found, add it into match.
			match.Lines = append(match.Lines, line)
			match.MatchString = append(match.MatchString, readString)
		}
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
	//TODO implement me
	panic("implement me")
}

func NewStdReader(reader io.Reader, finder []Finder) (Reader, error) {
	read := bufio.NewReader(reader)
	return &StdReader{
		finder: finder,
		reader: read,
		Output: nil,
	}, nil
}
