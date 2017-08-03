// mirthapi project main.go
package main

import (
	"flag"

	"github.com/caimeo/iniflags"
	"github.com/caimeo/stickyjar/curljar"
	"github.com/caimeo/stickyjar/tracer"
)

var verboseMode = flag.Bool("verbose", false, "Verbose trace output.")
var cookieFileP = flag.String("file", "", "Cookiejar file")
var cookieFile string

var t tracer.Tracer

func main() {
	iniflags.SetConfigFile(".settings")
	iniflags.SetAllowMissingConfigFile(true)
	iniflags.Parse()

	cookieFile = *cookieFileP

	t := tracer.New(false)

	t.Always("Cookie Curl")
	t.Always(cookieFile)

	curlcookie.New(cookieFile, nil)

	//	fmt.Printf(cj.String())

}
