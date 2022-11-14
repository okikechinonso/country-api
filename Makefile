run: |
	gofmt -w .
	go run main.go 

mock-db:
	mockgen -source=domain/repository.go -destination=test/mock_db.go -package=test