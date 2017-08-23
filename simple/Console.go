package simple

import "fmt"

type Console struct {
	verbose bool
}

func (t *Console) Verbose(args ...interface{}) {
	if t == nil {
		return
	}
	if t.verbose {
		t.Always(args...)
	}
}

func (t *Console) Always(args ...interface{}) {
	if t == nil {
		return
	}
	fmt.Println(args)
}

func (t *Console) IsVerbose() bool {
	if t == nil {
		return false
	}
	return t.verbose
}

func NewConsole(verbose bool) *Console {
	return &Console{verbose: verbose}
}
