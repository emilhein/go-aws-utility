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
	"github.com/go-aws-utility/util/services"
)

func main() {
	// Get credentials
	credentials := services.GetAccountInfo()
	fmt.Printf("We found your acocunt Key: %v \n", credentials)

	// //list buckets
	buckets, err := services.GetS3Buckets()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	buckets.ListBuckets()

	// read files in parrallel
	input := services.FilesInput{Bucket: "jubii-bi-inbox", FileNames: []string{"reports/fees.json"}}
	returnValues := services.GetS3Files(input)

	fmt.Println("All files read!")
	for i := 0; i < len(returnValues.Files); i++ {
		fmt.Println("", returnValues.Files[i])

	}

}


```
