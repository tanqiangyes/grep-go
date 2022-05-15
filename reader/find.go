package reader

import (
	"github.com/fatih/color"
	"github.com/tanqiangyes/grep-go/in_errors"
	"math/rand"
	"regexp"
	"strings"
	"time"
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
			return makeIExact(search), nil
		} else {
			return makeExact(search), nil
		}
	} else {
		return makeRegexp(search)
	}
}

// ExactFinder ： a exact Finder for searching， Case sensitive。
type ExactFinder struct {
	Search string
	// color for print
	color *color.Color
}

func makeExact(search string) *ExactFinder {
	rand.Seed(int64(time.Now().Nanosecond()))
	index := rand.Int()%len(colorMap) + 1
	c := color.New(colorMap[int64(index)])
	return &ExactFinder{Search: search, color: c}
}

func (e *ExactFinder) Find(input string) (string, bool) {
	if strings.Contains(input, e.Search) {
		// do color replaceAll
		return strings.ReplaceAll(input, e.Search, e.color.Sprint(e.Search)), true
	}
	return "", false
}

// IExactFinder ： a exact Finder for searching，not Case-sensitive。
type IExactFinder struct {
	Search string
	// color for print
	color *color.Color
}

func makeIExact(search string) *IExactFinder {
	rand.Seed(int64(time.Now().Nanosecond()))
	index := rand.Int()%len(colorMap) + 1
	c := color.New(colorMap[int64(index)])
	return &IExactFinder{Search: search, color: c}
}

func (i *IExactFinder) Find(input string) (string, bool) {
	if strings.Contains(strings.ToLower(input), strings.ToLower(i.Search)) {
		re := regexp.MustCompile(`(?i)` + i.Search)
		return re.ReplaceAllString(input, i.color.Sprint(i.Search)), true
	}
	return "", false
}

// RegexpFinder ： a regexp finder for searching string.
type RegexpFinder struct {
	*regexp.Regexp
	// color for print
	color *color.Color
}

func makeRegexp(reg string) (*RegexpFinder, error) {
	compile, err := regexp.Compile(reg)
	if err != nil {
		return nil, in_errors.ErrCreateRegexpFinder
	}

	rand.Seed(int64(time.Now().Nanosecond()))
	index := rand.Int()%len(colorMap) + 1
	c := color.New(colorMap[int64(index)])
	return &RegexpFinder{Regexp: compile, color: c}, err
}

func (r RegexpFinder) Find(input string) (string, bool) {
	if r.MatchString(input) {
		// use first suitable string as replace string.
		return r.ReplaceAllString(input, r.color.Sprint("$1")), true
	}
	return "", false
}
