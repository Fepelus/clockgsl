.PHONY: clean, run

clock: clock.go
	go build $<

clock.go: model.xml modelToGo.gsl
	gsl $<

clean:
	rm clock.go clock

run: clock
	./clock
