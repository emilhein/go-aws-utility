package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Input struct {
	Key         string
	Bucket      string
	FileChannel chan []byte
	Wg          *sync.WaitGroup
}

type FilesInput struct {
	Bucket    string
	FileNames []string
}
type BucketList struct {
	Names []string
}

type S3JSONFiles struct {
	Files []interface{}
}

func (b *BucketList) ListBuckets() {
	for i := 0; i < len(b.Names); i++ {
		fmt.Println("", b.Names[i])
	}
}

func GetS3Buckets() (BucketList, error) {
	config := GetConfig()
	client := s3.New(config)
	input := &s3.ListBucketsInput{}

	res, err := client.ListBuckets(input)
	if err == nil {
		var ListBuckets BucketList
		for _, bucket := range res.Buckets {
			// specificBucket := Bucket{Name: *bucket.Name, CreationDate: *bucket.CreationDate}
			ListBuckets.Names = append(ListBuckets.Names, *bucket.Name)
		}
		return ListBuckets, nil
	}
	return BucketList{}, errors.New("Could not get any buckets")
}

func GetS3Files(f FilesInput) S3JSONFiles {
	fileChannel := make(chan []byte, len(f.FileNames))
	var wg sync.WaitGroup

	for w := 0; w < len(f.FileNames); w++ {
		wg.Add(1)
		input := S3Input{Bucket: f.Bucket, Key: f.FileNames[w], FileChannel: fileChannel, Wg: &wg}
		go ReadFile(input)
	}
	wg.Wait()
	close(fileChannel)

	var fileList []interface{}

	for elem := range fileChannel {
		fmt.Printf("Parsing file to JSON... \n")
		var randomObject interface{}
		json.Unmarshal(elem, &randomObject)
		fileList = append(fileList, randomObject)

	}
	return S3JSONFiles{Files: fileList}
}

func ReadFile(h S3Input) {
	defer h.Wg.Done()
	fmt.Printf("Reading file.. %v/%v  \n", h.Bucket, h.Key)
	sess, err := session.NewSession(&aws.Config{Region: aws.String(S3_REGION)})
	if err != nil {
		fmt.Println("Error")
		// Handle error
	}
	results, err := s3.New(sess).GetObject(&s3.GetObjectInput{
		Bucket: aws.String(h.Bucket),
		Key:    aws.String(h.Key),
	})
	if err != nil {
		fmt.Println("Error getting file")
	}
	defer results.Body.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, results.Body); err != nil {
		fmt.Println("Error copying to buffer")
	}
	h.FileChannel <- buf.Bytes() //send file to channel
}
