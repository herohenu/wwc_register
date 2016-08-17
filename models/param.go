package models

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"time"
)

/*
	GetDevSched get a particular did info
*/

const (
	DYNANODB_PERSON_TABLE = "Person"
)

func AddUseInput(p *Person) (params dynamodb.PutItemInput) {
	ts := fmt.Sprintf("%f", float64(time.Now().UnixNano())/1000000000)
	params = dynamodb.PutItemInput{
		TableName: aws.String(DYNANODB_PERSON_TABLE),
		Item: map[string]*dynamodb.AttributeValue{
			"primary_email": {S: aws.String(p.PrimaryEmail)},
			"name":          {S: aws.String(p.Name)},
			"mobile":        {S: aws.String(p.Mobile)},
			"class":         {SS: aws.StringSlice(p.Class)},
			"prog_language": {SS: aws.StringSlice(p.ProgLanguage)},
			"create_at":     {S: aws.String(ts)},
		},
	}
	return params
}
