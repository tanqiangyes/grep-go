package reader

import (
	"bufio"
	"errors"
	"github.com/tanqiangyes/grep-go/tools"
	"io"
	"os"
	"strings"
)

type MultiReader struct {
	finder []Finder
	files  []string

	Output []MatchRes
	Error  error
}

func (m *MultiReader) Result() []MatchRes {
	return m.Output
}

func (m *MultiReader) Run() {
	lens := len(m.files)
	if lens < 1 {
		m.Error = errors.New("no file to grep")
		return
	}
	ch := make(chan MatchRes)
	for _, file := range m.files {
		go m.dealFileMatch(file, ch)
	}

	for {
		select {
		case res := <-ch:
			m.Output = append(m.Output, res)
			if len(m.Output) >= lens {
				return
			}
		}
	}
}

func (m *MultiReader) dealFileMatch(file string, ch chan MatchRes) {
	f, err := os.Open(file)
	if err != nil { //read failed, return
		//fmt.Println(f.Name(), err.Error())
		ch <- MatchRes{
			Filename:    f.Name(),
			Lines:       nil,
			MatchString: nil,
		}
		return
	}
	// close the file reader.
	defer f.Close()
	// buffer read
	br := bufio.NewReader(f)
	var line int64 = 1
	var match MatchRes
	match.Filename = f.Name()
	for {
		readString, err := br.ReadString('\n')
		if err != nil {
			// EOF, we should break this loop, and return data.
			if errors.Is(err, io.EOF) {
				break
			}
			// something wrong, write the error and return.
			m.Error = err
			// send res to chan, make sure res lens.
			ch <- MatchRes{Filename: f.Name()}
			return
		}

		// do some string processing
		readString = strings.TrimSpace(readString)
		readString = strings.Trim(readString, " ")
		readString = strings.Trim(readString, "\n")

		if readString, ok := m.find(readString); ok {
			// we found, add it into match.
			match.Lines = append(match.Lines, line)
			match.MatchString = append(match.MatchString, readString)
		}
		line++
	}
	ch <- match
	return
}

func (m *MultiReader) find(str string) (string, bool) {
	var found bool
	for _, finder := range m.finder {
		if s, ok := finder.Find(str); ok {
			str = s
			found = true
		}
	}
	return str, found
}

func (m *MultiReader) Close() {
	return
}

func (m *MultiReader) IsError() error {
	return m.Error
}

func NewMultiReader(path []string, finder []Finder, isRecursive bool) (Reader, error) {
	// get all files
	files, err := tools.Files(isRecursive, path...)
	if err != nil {
		return nil, err
	}
	return &MultiReader{
		finder: finder,
		files:  files,
		Output: nil,
	}, nil
}
