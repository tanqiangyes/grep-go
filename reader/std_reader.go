package reader

import (
	"bufio"
	"io"
)

type StdReader struct {
	finder []Finder
	reader io.Reader

	Output []MatchRes
}

func (s *StdReader) Result() []MatchRes {
	return s.Output
}

func (s *StdReader) Run() {
	//TODO implement me
	panic("implement me")
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
