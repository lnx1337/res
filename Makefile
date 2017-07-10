install:
	rm -rf $$GOPATH/src/resuelve;
	# mkdir -p $$GOPATH/src/;
	ln -sf $$PWD/ $$GOPATH/src/;
	go get -u $$(go list -f '{{ join .Deps " " }}' ./...) || true

	#Runing test ...
	go test resuelve/invoice/invoice
	go test resuelve/invoice/request
	go test resuelve/invoice/utils

build:
	go build -o resuelve main.go

updatedeps:
	go get -u $$(go list -f '{{ join .Deps " " }}' ./...) || true
	