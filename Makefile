.PHONY: clean, test

all: test cli/cli server/server

test: clock_test.go clock.go
	go test

%.go: model.xml modelToGo.gsl
	gsl $<

cli/cli: cli/cli.go
	pushd cli; go build; popd

server/server: server/server.go
	pushd cli; go build; popd

clean:
	rm clock.go clock_test.go 
