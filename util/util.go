package util

import (
	"errors"
	"fmt"

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

func (b *BucketList) ListBuckets() {
	for i := 0; i < len(b.Names); i++ {
		fmt.Println("", b.Names[i])
	}
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
