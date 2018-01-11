package main

import (
	"github.com/jdef/plugins1/pkg/foo"
)

func Foo(_ foo.F) { println("foo!") }

func main() {} // noop for plugins
