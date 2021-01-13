build-literals: embed
	garble -literals build -o dist *.go
build-literals-tiny: embed
	garble -literals -tiny build -o dist *.go
embed:
	rice embed-go