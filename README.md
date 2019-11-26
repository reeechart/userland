# Userland

## Project Description
Given a list of API Contracts based on [Simukti's Userland APIAry](https://userland.docs.apiary.io), implement all of the API using Golang

## Starting Development Server

### Without Docker
```go
go run main.go
```

### With Docker
To start server, enter the following
```sh
docker-compose up
```

To delete network, enter the following
```sh
docker-compose down
```

## Run Test
```go
go test ./... -v
```

## Author
[Ferdinandus Richard](https://github.com/reeechart) (mentored by Abdi Pratama)
