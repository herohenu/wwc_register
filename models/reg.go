package models

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
	"wwc_register/helper"
)

func (p *Person) AddUser() (err error) {
	param := AddUseInput(p)
	resp, err := helper.DynOp(param)
	if err != nil {
		return err
	}
	log.Println("AddUser resp", resp.(dynamodb.PutItemOutput))
	return nil
}
