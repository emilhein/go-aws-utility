[![Go Report Card](https://goreportcard.com/badge/github.com/emilhein/golang-starter)](https://goreportcard.com/report/github.com/emilhein/golang-starter)
[![Build Status](https://travis-ci.org/emilhein/golang-starter.svg?branch=master)](https://travis-ci.org/emilhein/golang-starter)

# golang-starter

A starter project for playing around with Golang.

### Serverless function

This project shows how you can setup automatic deployment to AWS using the serverless framework and Travis CI

The serverless function simply echos back what it recieves

### Webserver

cd \$GOPATH/src/golang-starter/src/github.com/emilhein/webserver is a webserver with 3 endpoint

1. GET /simple // Will run some simple logic in the console

2. POST /getS3file // Will retrive a file from S3, parse it and log it

3. GET /interfaces // Will print simple output using interfaces

4. GET /startmining // Will start 3 gophers that communicate over channels

## Get started

```
cd \$GOPATH/src/golang-starter/src/github.com/emilhein/webserver
go build
.\webserver.exe

```

For local development i would use [gin](https://github.com/gin-gonic/gin) and run:

```
cd $GOPATH/src/golang-starter/src/github.com/emilhein/webserver
gin run main.go
```
