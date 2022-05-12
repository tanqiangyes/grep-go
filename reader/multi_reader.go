package reader

import (
	"fmt"
	"github.com/tanqiangyes/grep-go/tools"
)

type MultiReader struct {
	finder Finder
	files  []string

	Output []MatchRes
}

func (m *MultiReader) Print() {
	//TODO implement me
	panic("implement me")
}

func (m *MultiReader) Run() {
	//TODO implement me
	panic("implement me")
}

func (m *MultiReader) Close() {
	//TODO implement me
	panic("implement me")
}

func NewMultiReader(path []string, search string, isRecursive, exact, sensitive bool) (Reader, error) {
	finder, err := NewFinder(search, exact, sensitive)
	if err != nil {
		return nil, err
	}
	fmt.Println(1111, finder)
	files, err := tools.Files(isRecursive, path...)
	if err != nil {
		return nil, err
	}
	fmt.Println(1111, files)
	return &MultiReader{
		finder: finder,
		files:  files,
		Output: nil,
	}, nil
}
