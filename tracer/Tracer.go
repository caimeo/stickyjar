package tracer

import "fmt"

type Tracer struct {
	verbose bool
}

func (t *Tracer) Verbose(s string) {
	if t.verbose {
		t.Always(s)
	}
}

func (t *Tracer) Always(s string) {
	fmt.Println(s)
}

func New(verbose bool) *Tracer {
	return &Tracer{verbose: verbose}
}
