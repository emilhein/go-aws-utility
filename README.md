[![Go Report Card](https://goreportcard.com/badge/github.com/emilhein/go-aws-utility)](https://goreportcard.com/report/github.com/emilhein/go-aws-utility)
[![Build Status](https://travis-ci.org/emilhein/go-aws-utility.svg?branch=master)](https://travis-ci.org/emilhein/go-aws-utility)

# go-aws-utility

This project acts like a wrapper for the AWS go SDK.
More functionality will be following...

## How to use

```
package main

import (
	"fmt"
	"github.com/emilhein/go-aws-utility/util"
)
func main() {
	// Get credentials

	credentials := util.GetAccountInfo()
	fmt.Printf("We found your acocunt Key: %v", credentials)

	//list buckets
	buckets, err := util.GetS3Buckets()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	buckets.ListBuckets()

	// read files in parrallel
	input := util.FilesInput{Bucket: "[YOUR-BUCKET]", FileNames: []string{"Filepath1", "filepath2"}}
	util.GetS3Files(input)
	fmt.Println("All files read!")

```
