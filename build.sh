#!/bin/bash

go clean -i -r ./...
rm -v -rf *.so

go install ./cmd/...
go build -buildmode=plugin -o plugin.so ./plugin

# works
go test -v ./cmd/app/
go test -v -cover ./cmd/app/
go test -v -tags e2e github.com/jdef/plugins1/pkg/foo ./cmd/app/
go test -v -tags e2e -cover github.com/jdef/plugins1/pkg/foo ./cmd/app/

# broken
echo EXECUTING WITH coverpkg
go test -v -cover -coverpkg github.com/jdef/plugins1/pkg/foo ./cmd/app/
