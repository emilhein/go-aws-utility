[![Go Report Card](https://goreportcard.com/badge/github.com/emilhein/go-aws-utility)](https://goreportcard.com/report/github.com/emilhein/go-aws-utility)
[![Build Status](https://travis-ci.org/emilhein/go-aws-utility.svg?branch=master)](https://travis-ci.org/emilhein/go-aws-utility)

# go-aws-utility

This project acts like a wrapper for the AWS go SDK.

## How to use

```
package main

import (
	"fmt"
	"github.com/go-aws-utility/util"
)
func main() {
	buckets, err := util.GetS3Buckets()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	buckets.ListBuckets()
}

```
