# Userland

[![Go Report Card](https://goreportcard.com/badge/github.com/reeechart/userland)](https://goreportcard.com/report/github.com/reeechart/userland)
<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-61%25-brightgreen.svg?longCache=true&style=flat)</a>

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
