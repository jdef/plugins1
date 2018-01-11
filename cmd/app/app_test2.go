// +build e2e

package main

import "testing"

func TestFoo2(t *testing.T) {
	RunPluginAt("../../plugin.so")
}
