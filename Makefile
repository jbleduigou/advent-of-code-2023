vet:
	go vet ./...

test: 
	go test -v ./... 

0: 
	cd 0-calorie-counting && go test -v ./... 