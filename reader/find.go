package reader

import (
	"github.com/tanqiangyes/grep-go/in_errors"
	"regexp"
	"strings"
)

// Finder is a interface for searching.
// it can find text string based on a plain text or regexp.
// if cannot found it returns an empty string.
type Finder interface {
	Find(input string) (string, bool)
}

func NewFinder(search string, regexp, sensitive bool) (Finder, error) {
	if !regexp {
		if sensitive {
			return makeExact(search), nil
		} else {
			return makeIExact(search), nil
		}
	} else {
		return makeRegexp(search)
	}
}

// ExactFinder ： a exact Finder for searching， Case sensitive。
type ExactFinder struct {
	Search string
}

func makeExact(search string) *ExactFinder {
	return &ExactFinder{Search: search}
}

func (e *ExactFinder) Find(input string) (string, bool) {
	if strings.Contains(input, e.Search) {
		return input, true
	}
	return "", false
}

// IExactFinder ： a exact Finder for searching，not Case-sensitive。
type IExactFinder struct {
	Search string
}

func makeIExact(search string) *IExactFinder {
	return &IExactFinder{Search: search}
}

func (i *IExactFinder) Find(input string) (string, bool) {
	if strings.Contains(strings.ToLower(input), strings.ToLower(i.Search)) {
		return input, true
	}
	return "", false
}

// RegexpFinder ： a regexp finder for searching string.
type RegexpFinder struct {
	*regexp.Regexp
}

func makeRegexp(reg string) (*RegexpFinder, error) {
	compile, err := regexp.Compile(reg)
	if err != nil {
		return nil, in_errors.ErrCreateRegexpFinder
	}
	return &RegexpFinder{compile}, err
}

func (r RegexpFinder) Find(input string) (string, bool) {
	if r.MatchString(input) {
		return input, true
	}
	return "", false
}
