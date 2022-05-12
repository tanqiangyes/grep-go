package reader

import (
	"bufio"
	"io"
)

type StdReader struct {
	finder Finder
	reader io.Reader

	Output []MatchRes
}

func (s *StdReader) Print() {
	//TODO implement me
	panic("implement me")
}

func (s *StdReader) Run() {
	//TODO implement me
	panic("implement me")
}

func (s *StdReader) Close() {
	//TODO implement me
	panic("implement me")
}

func NewStdReader(reader io.Reader, search string, exact, sensitive bool) (Reader, error) {
	finder, err := NewFinder(search, exact, sensitive)
	if err != nil {
		return nil, err
	}
	read := bufio.NewReader(reader)

	return &StdReader{
		finder: finder,
		reader: read,
		Output: nil,
	}, nil
}
