package service

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func CreateTableUsers(db *dynamodb.DynamoDB) {
	tableName := "users"
	users := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}
	_, err := db.CreateTable(users)
	if err != nil {
		//log.Fatalf("Got error calling CreateTable: %s", err)
		fmt.Println(err)
		return
	}
	fmt.Println("Created the table :", tableName)
}
