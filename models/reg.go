package models

import (
	"log"
	"wwc_register/helper"
)

func (p *Person) AddUser() (err error) {
	param := AddUseInput(p)
	resp, err := helper.DynOp(param)
	log.Println("resp", resp)
	return err
}
