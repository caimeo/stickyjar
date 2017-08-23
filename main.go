// mirthapi project main.go
package main

import (
	"flag"

	"github.com/caimeo/iniflags"
	"github.com/caimeo/stickyjar/curljar"
	"github.com/caimeo/stickyjar/simple"
)

var verboseMode = flag.Bool("verbose", false, "Verbose trace output.")
var cookieFileP = flag.String("file", "", "Cookiejar file")
var cookieFile string

var t simple.Console

func main() {
	iniflags.SetConfigFile(".settings")
	iniflags.SetAllowMissingConfigFile(true)
	iniflags.Parse()

	cookieFile = *cookieFileP

	t := simple.NewConsole(*verboseMode)

	t.Always("Cookie Curl")
	t.Always(cookieFile)

	cj, _ := curljar.New(cookieFile, nil)

	t.Verbose(cj.String())

}
