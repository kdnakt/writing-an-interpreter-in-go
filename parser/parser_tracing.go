package parser

import (
	"fmt"
	"os"
	"strings"
)

var isTrace = false
func init() {
	for _, a := range os.Args {
		if a == "--trace" || a == "-t" {
			isTrace = true
			break
		}
	}
}

var traceLevel int

const traceIdentPlaceholder string = "\t"

func identLevel() string {
	return strings.Repeat(traceIdentPlaceholder, traceLevel-1)
}

func tracePrint(fs string) {
	if isTrace {
		fmt.Printf("%s%s\n", identLevel(), fs)
	}
}

func incIdent() { traceLevel = traceLevel + 1 }
func decIdent() { traceLevel = traceLevel - 1 }

func trace(msg string) string {
	incIdent()
	tracePrint("BEGIN " + msg)
	return msg
}

func untrace(msg string) {
	tracePrint("END " + msg)
	decIdent()
}
