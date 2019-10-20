package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
)

const S3_REGION = "eu-west-1"

type Account struct {
	AccessKeyID string
}
type BucketList struct {
	Names []string
}
type DynamoDbList struct {
	Names []string
}

type FilesInput struct {
	Bucket    string
	FileNames []string
}
type FilesOutput struct {
	Status       bool
	FileContents []byte
}

type S3Input struct {
	Key         string
	Bucket      string
	FileChannel chan []byte
}

var wg sync.WaitGroup

func (b *BucketList) ListBuckets() {
	for i := 0; i < len(b.Names); i++ {
		fmt.Println("", b.Names[i])
	}
}

func GetS3Files(f FilesInput) {
	wg.Add(len(f.FileNames))
	fileChannel := make(chan []byte, len(f.FileNames))
	for w := 0; w < len(f.FileNames); w++ {
		input := S3Input{Bucket: f.Bucket, Key: f.FileNames[w], FileChannel: fileChannel}
		go ReadFile(input)
	}

	for elem := range fileChannel {
		fmt.Println("From Finder: ")
		var randomObject interface{}
		json.Unmarshal(elem, &randomObject)

		fmt.Println("", randomObject)

	}
	close(fileChannel)

	wg.Wait()
	fmt.Println("Files read!")
	// res, err :=
	// if err != nil {
	// 	return FilesOutput{}, err // errors.New("Could not get any tables", err.String())
	// }
	// result := FilesOutput{Status: true, FileContents: res}
	// return result, nil
}
func GetConfig() *session.Session {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(S3_REGION)})
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}
	return sess
}

func GetAccountInfo() Account {
	session := GetConfig()
	cred, err := session.Config.Credentials.Get()
	if err != nil {
		panic("Could not get your credentials, " + err.Error())
	}
	account := Account{AccessKeyID: cred.AccessKeyID}
	return account
}

func GetDynamoDbTables() (DynamoDbList, error) {
	config := GetConfig()
	client := dynamodb.New(config)
	input := &dynamodb.ListTablesInput{}

	res, err := client.ListTables(input)
	if err == nil {
		var DynamoList DynamoDbList
		for _, table := range res.TableNames {
			fmt.Println("", *table)

			DynamoList.Names = append(DynamoList.Names, *table)
		}
		return DynamoList, nil
	}
	return DynamoDbList{}, errors.New("Could not get any tables")
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

func ReadFile(h S3Input) {
	defer wg.Done()
	fmt.Println("Reading file..")
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
		// return nil, err
	}
	defer results.Body.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, results.Body); err != nil {
		fmt.Println("Error copying to buffer")

		// return nil, err
	}
	h.FileChannel <- buf.Bytes() //send file to channel

	// return buf.Bytes(), nil
}
