package reader

type Reader interface {
	//a fun to start
	Run()

	// a func to close the reader
	Close()

	Result() []MatchRes

	IsError() error
}

// a return value
type MatchRes struct {
	Filename    string
	Lines       []int64
	MatchString []string
}
