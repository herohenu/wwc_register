package helper

import (
	"errors"
	"strconv"
)

func GetError(code int) error {
	codestr := strconv.Itoa(code)
	val, exist := ErrorTable[codestr]
	if exist {
		return errors.New(val)
	}
	return errors.New("error code not exist")
}

func GetErrorDesc(code int) (des string, err error) {
	codestr := strconv.Itoa(code)
	val, exist := ErrorTable[codestr]
	if exist {
		return val, nil
	}
	return "", errors.New("error code not exist")
}

var ErrorTable map[string]string = map[string]string{
	"5000": "oops, server was done",
	"5002": "device command can not to translate",
	"5003": "DDB query failed, please try again",
	"5004": "dynamodb query failed, please try again",
	"4000": "invalid access token",
	"4002": "Json Parsing Fail",
}
