.PHONY: clean, test

test: clock_test.go clock.go
	go test

%.go: model.xml modelToGo.gsl
	gsl $<

clean:
	rm clock.go clock_test.go clock
