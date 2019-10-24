package services

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDbList struct {
	Names []string
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
