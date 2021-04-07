package services

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	. "go-service/internal/models"
)

type DynamodbUserService struct {
	DB *dynamodb.DynamoDB
}

func NewUserService(db *dynamodb.DynamoDB) *DynamodbUserService {
	return &DynamodbUserService{DB: db}
}

func (m *DynamodbUserService) GetAll() (*[]User, error) {
	params := &dynamodb.ScanInput{
		TableName: aws.String(User{}.GetTableName()),
	}
	result, err := m.DB.Scan(params)
	if err != nil {
		return nil, err
	}
	users := make([]User, 0)
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &users)
	return &users, err
}

func (m *DynamodbUserService) Load(id string) (*User, error) {
	key, err := dynamodbattribute.MarshalMap(UserKey{
		Id:  id,
	})
	if err != nil {
		return nil, err
	}
	var user User
	result, err := m.DB.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(user.GetTableName()),
		Key: key,
	})
	if err != nil {
		return nil, err
	}
	err = dynamodbattribute.UnmarshalMap(result.Item, &user)
	return &user, err
}

func (m *DynamodbUserService) Insert(user *User) (int64, error) {
	avUser, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return 0, err
	}
	input := &dynamodb.PutItemInput{
		Item:      avUser,
		TableName: aws.String(user.GetTableName()),
	}
	_, err = m.DB.PutItem(input)
	return 1, err
}

func (m *DynamodbUserService) Update(user *User) (int64, error) {
	key, err := dynamodbattribute.MarshalMap(UserKey{
		Id:  user.Id,
	})
	if err != nil {
		return 0, err
	}
	update, err := dynamodbattribute.MarshalMap(UserUpdate{
		Username:  user.Username,
		Phone: user.Phone,
		Email: user.Email,
		DateOfBirth: user.DateOfBirth,
	})
	if err != nil {
		return 0, err
	}
	fmt.Printf("%v %v", key, update)
	input := &dynamodb.UpdateItemInput{
		Key:                       key,
		TableName:                 aws.String(user.GetTableName()),
		UpdateExpression:          aws.String("set username=:u, email=:e, phone=:p, dateOfBirth=:d"),
		ExpressionAttributeValues: update,
		ReturnValues:              aws.String("UPDATED_NEW"),
	}
	_, err = m.DB.UpdateItem(input)
	if err != nil{
		return 0, err
	}
	return 1, err
}

func (m *DynamodbUserService) Delete(id string) (int64, error) {
	key, err := dynamodbattribute.MarshalMap(UserKey{
		Id:  id,
	})
	if err != nil {
		return 0, err
	}
	input := &dynamodb.DeleteItemInput{
		Key: key,
		TableName: aws.String(User{}.GetTableName()),
	}
	_, err = m.DB.DeleteItem(input)
	if err != nil {
		return 0, err
	}
	return 1, err
}