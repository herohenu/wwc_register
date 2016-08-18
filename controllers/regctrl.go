package controllers

import (
	"encoding/json"
	//"errors"
	"github.com/astaxie/beego"
	"log"
	"wwc_register/helper"
	"wwc_register/models"
)

type RegController struct {
	beego.Controller
}

func (c *RegController) AddUser() {
	var student models.Person
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &student); err != nil {
		log.Println("student", student)
		resp := GetErrorResp(4002)
		c.Data["json"] = &resp
		log.Println("err:", err, "resp:", resp)
		c.ServeJSON()
	}
	log.Println("AddUser Request:", student)
	err := student.AddUser()

	if err != nil {
		resp := GetErrorResp(5004)
		c.Data["json"] = &resp
		log.Println("err:", err, "resp:", resp)
		c.ServeJSON()
	}
	resp := RespOK()
	c.Data["json"] = &resp
	//c.Data["json"] = "{\"ObjectId\":\"" + "123" + "\"}"
	log.Println("Respons OK:", resp)
	log.Println("Data:", c.Data["json"])
	c.ServeJSON()

}

func RespOK() models.RespMsg {
	var resp models.RespMsg
	resp.Code = 200
	resp.Message = "200 OK"
	return resp

}

func GetErrorResp(code int) models.RespMsg {
	var resp models.RespMsg
	desc, err := helper.GetErrorDesc(code)
	if err != nil {
		resp.Code = 9000
		resp.Message = err.Error()
		return resp
	}
	resp.Code = code
	resp.Message = desc
	return resp
}
