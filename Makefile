.PHONY: clean, test

all: test cli/cli server/server

test: clock_test.go clock.go
	go test

%.go: model.xml modelToGo.gsl
	gsl $<

cli/cli: cli/cli.go clock.go
	pushd cli; go build; popd

server/server: server/server.go clock.go
	pushd server; go build; popd

clean:
	rm clock.go clock_test.go 

install: server/server cli/cli
	cp server/server /usr/local/bin/clocksrv
	cp cli/cli ${HOME}/Dropbox/bin/clockgsl

