all : golist errcheck gofmt

golist :
	go list ./...
errcheck :
	errcheck ./...
gofmt :
	gofmt -l -w -s .

