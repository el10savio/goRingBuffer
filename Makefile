
fmt:
	@echo "go fmt goRingBuffer"
	go fmt ./...

vet:
	@echo "go vet goRingBuffer"
	go vet ./...

lint:
	@echo "go lint goRingBuffer"
	golint ./...

test:
	@echo "Testing goRingBuffer"
	go test -v -race --cover ./...

codespell:
	@echo "checking goRingBuffer spellings"
	codespell
