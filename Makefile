all : golist errcheck gofmt

golist :
	go list ./...
errcheck :
	errcheck ./...
gofmt :
	gofmt -l -w -s .

goimports_all :
	find . -print | grep --regex '.*\.go' | xargs goimports -w
