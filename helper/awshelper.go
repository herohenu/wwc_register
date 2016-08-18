package helper

import (
	//"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"reflect"
)

const (
	API_REGION       = "us-west-2"
	S3_BUCKET_REGION = "us-west-2"
	//CRED_FILE_PATH   = "../../../.aws/credentials"
	//CRED_FILE_PATH = "cygdrive/c/Users/pumpkin pai/.aws/credentials"
	CRED_FILE_PATH = "C:\\Users\\pumpkin pai\\.aws\\credentials"
	PROFILE        = "wwc"
)

/*******
	Credentials
********/
func InitConfig(region string) (*aws.Config, error) {
	creds := GetCredentialShared(PROFILE)
	conf := aws.NewConfig().WithRegion(region).WithCredentials(creds)
	//creds, conf := GetCredentialChain()
	val, err := creds.Get()
	if err != nil {
		log.Println("InitConfig:", err)
	}
	log.Println("Cred ProviderName:", val.ProviderName)
	conf.WithRegion(region).WithCredentials(creds)
	return conf, nil
}
func GetCredentialShared(profile string) *credentials.Credentials {
	return credentials.NewSharedCredentials(CRED_FILE_PATH, profile)
}
func GetCredentialChain() (*credentials.Credentials, *aws.Config) {
	config := aws.NewConfig()
	ec2m := ec2metadata.New(session.New(), config)
	var ProviderList []credentials.Provider = []credentials.Provider{
		&credentials.EnvProvider{},
		&credentials.SharedCredentialsProvider{Filename: CRED_FILE_PATH, Profile: PROFILE},
		&ec2rolecreds.EC2RoleProvider{
			Client: ec2m,
		},
	}
	creds := credentials.NewChainCredentials(ProviderList)
	return creds, config
	//return credentials.NewStaticCredentials(accessKey, secretKey, ``)
}

/********
AWS Operation
**********/

func DynOp(input interface{}) (output interface{}, err error) {
	itype := reflect.TypeOf(input)
	log.Println("DynOp Name:", itype.Name())
	conf, err := InitConfig(API_REGION)
	if err != nil {
		log.Println("GetSchdDyn Err:", err)
		return nil, err
	}
	svc := dynamodb.New(session.New(), conf)
	switch itype.Name() {
	case "QueryInput":
		qi := input.(dynamodb.QueryInput)
		resp, operr := svc.Query(&qi)
		output = *resp
		err = operr
	case "GetItemInput":
		geti := input.(dynamodb.GetItemInput)
		resp, operr := svc.GetItem(&geti)
		output = *resp
		err = operr
	case "PutItemInput":
		puti := input.(dynamodb.PutItemInput)
		resp, operr := svc.PutItem(&puti)
		output = *resp
		err = operr
	case "DeleteItemInput":
		deli := input.(dynamodb.DeleteItemInput)
		resp, operr := svc.DeleteItem(&deli)
		output = *resp
		err = operr
	case "UpdateItemInput":
		log.Println("UpdateItemInput")
		updi := input.(dynamodb.UpdateItemInput)
		resp, operr := svc.UpdateItem(&updi)
		output = *resp
		err = operr
	case "BatchGetItemInput":
		log.Println("DynIp BatchGetItemInput")
		bti := input.(dynamodb.BatchGetItemInput)
		resp, operr := svc.BatchGetItem(&bti)
		output = *resp
		err = operr
	}
	if err != nil {
		log.Println("DynOp Err:", err)
		return nil, err
	}

	return output, nil

}

func S3Op(input interface{}) (output interface{}, err error) {
	itype := reflect.TypeOf(input)
	log.Println("Name:", itype.Name())

	conf, err := InitConfig(S3_BUCKET_REGION)
	if err != nil {
		log.Println("GetSchdDyn Err:", err)
		return nil, err
	}
	svc := s3.New(session.New(), conf)
	switch itype.Name() {
	case "PutObjectInput":
		po := input.(s3.PutObjectInput)
		resp, operr := svc.PutObject(&po)
		output = *resp
		err = operr
	case "DeleteObjectInput":
		do := input.(s3.DeleteObjectInput)
		resp, operr := svc.DeleteObject(&do)
		output = *resp
		err = operr
	}
	if err != nil {
		log.Println("S3Op Err:", err)
		return nil, err
	}
	return output, nil
}
