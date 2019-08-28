run:
	present -base $(GOPATH)/src/golang.org/x/tools/cmd/present

deps:
	go get -u golang.org/x/tools/cmd/present
