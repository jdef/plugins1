## DEMO

```sh
$ ./build.sh 
removed ‘plugin.so’
=== RUN   TestFoo
foo!
--- PASS: TestFoo (0.00s)
PASS
ok      github.com/jdef/plugins1/cmd/app        0.003s
=== RUN   TestFoo
foo!
--- PASS: TestFoo (0.00s)
PASS
coverage: 55.6% of statements
ok      github.com/jdef/plugins1/cmd/app        0.004s
?       github.com/jdef/plugins1/pkg/foo        [no test files]
=== RUN   TestFoo
foo!
--- PASS: TestFoo (0.00s)
PASS
ok      github.com/jdef/plugins1/cmd/app        0.003s
?       github.com/jdef/plugins1/pkg/foo        [no test files]
=== RUN   TestFoo
foo!
--- PASS: TestFoo (0.00s)
PASS
coverage: 50.0% of statements
ok      github.com/jdef/plugins1/cmd/app        0.003s  coverage: 50.0% of statements
EXECUTING WITH coverpkg
=== RUN   TestFoo
--- FAIL: TestFoo (0.00s)
panic: plugin.Open: plugin was built with a different version of package github.com/jdef/plugins1/pkg/foo [recovered]
        panic: plugin.Open: plugin was built with a different version of package github.com/jdef/plugins1/pkg/foo

goroutine 5 [running]:
testing.tRunner.func1(0xc4200980f0)
        /usr/local/go-1.9.2/src/testing/testing.go:711 +0x2d2
panic(0x5ecce0, 0xc420040400)
        /usr/local/go-1.9.2/src/runtime/panic.go:491 +0x283
github.com/jdef/plugins1/cmd/app.RunPluginAt(0x617e4f, 0xf)
        /home/vagrant/pluginbugs/src/github.com/jdef/plugins1/cmd/app/app.go:19 +0x1fb
github.com/jdef/plugins1/cmd/app.TestFoo(0xc4200980f0)
        /home/vagrant/pluginbugs/src/github.com/jdef/plugins1/cmd/app/app_test.go:6 +0x36
testing.tRunner(0xc4200980f0, 0x61faa8)
        /usr/local/go-1.9.2/src/testing/testing.go:746 +0xd0
created by testing.(*T).Run
        /usr/local/go-1.9.2/src/testing/testing.go:789 +0x2de
exit status 2
FAIL    github.com/jdef/plugins1/cmd/app        0.007s
```
