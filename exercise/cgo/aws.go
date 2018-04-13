package main

import (
	"C"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//export singleton
type singleton struct {
	ddb *dynamodb.DynamoDB
}

var instance *singleton
var once sync.Once

//export getInstance
func getInstance(accessKeyId string, secretKey string, region string) *singleton {

	once.Do(func() {
		//credentials
		creds := credentials.NewStaticCredentials(accessKeyId, secretKey, "")
		//create Session
		sess := session.Must(session.NewSession(&aws.Config{
			Region:      aws.String(region),
			Credentials: creds,
		},
		))
		// Create DynamoDB client
		ddb := dynamodb.New(sess)
		instance = &singleton{ddb}
	})
	return instance
}

//export getItem
func getItem(tableName string, keyName string, keyValue string) (*dynamodb.GetItemOutput, error) {
	result, err := instance.ddb.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			keyName: {
				S: aws.String(keyValue),
			},
		},
	})
	fmt.Println(result)
	return result, err
}

func main() {}
