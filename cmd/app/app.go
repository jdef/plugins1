package main

import (
	"os"
	"plugin"

	"github.com/jdef/plugins1/pkg/foo"
)

func main() {
	for _, path := range os.Args {
		RunPluginAt(path)
	}
}

func RunPluginAt(path string) {
	p, err := plugin.Open(path)
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("Foo")
	if err != nil {
		panic(err)
	}
	f.(func(foo.F))(struct{}{})
}
