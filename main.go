package main

import (
	"fmt"
	"github.com/go-aws-utility/util"
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

}
