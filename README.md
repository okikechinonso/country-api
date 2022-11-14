# Countries-API

Run the application using:
```bigquery
docker compose up
```

### Create Country (POST)

```
localhost:8081/countries
```
Payload
```
{
    "name": "Nigeria",
    "short_name": "NGN",
    "continent": "Africa",
    "is_operational": true
}
```

### Update Country (POST)
Endpoint
```
localhost:8081/countries/:id
```
payload
```
{
    "name": "Nigeria",
    "short_name": "NGN",
    "continent": "Africa",
    "is_operational": false
}
```
### Get Country By Id (GET)
Endpoint
```
localhost:8081/countries/:id
```
### Get Countries (GET)
Endpoint
```
localhost:8081/countries
```

## Tests
Testing is done using the GoMock framework. The ``gomock`` package and the ``mockgen``code generation tool are used for this purpose.
If you installed the dependencies using the command given above, then the packages would have been installed. Otherwise, installation can be done using the following commands:
```
go get github.com/golang/mock/mockgen@v1.6.0
go install github.com/golang/mock/mockgen
mockgen -source=domain/repository.go -destination=test/mock_db.go -package=test
```
run all the test files using:
```bigquery
go test -v ./...
```
