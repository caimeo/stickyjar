package simple

import "fmt"

type Console struct {
	verbose bool
	debug   bool
}

func (t *Console) Debug(args ...interface{}) {
	if t == nil {
		return
	}
	if t.debug {
		t.Always(args...)
	}
}

func (t *Console) Verbose(args ...interface{}) {
	if t == nil {
		return
	}
	if t.verbose || t.debug {
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

func (t *Console) IsDebug() bool {
	if t == nil {
		return false
	}
	return t.debug
}

func NewConsole(verbose bool, debug bool) *Console {
	return &Console{verbose: verbose, debug: debug}
}
