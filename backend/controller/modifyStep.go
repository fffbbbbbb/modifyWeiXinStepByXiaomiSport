package controller

import (
	"log"
	"modifyStep/util"

	"github.com/gin-gonic/gin"
)

type ModifyStepDto struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Step     int64  `json:"step"`
}

//ModifyStep 修改步数
func ModifyStep(c *gin.Context) {
	var req ModifyStepDto
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("err:", err)
		FailResp(c, "错误的请求参数")
		return
	}
	log.Printf("request:%+v\n", req)
	if req.Account == "" || req.Password == "" {
		FailResp(c, "账号或者密码不能为空")
		return
	}
	err := util.ModifyStep(req.Account, req.Password, req.Step)
	if err != nil {
		FailResp(c, err.Error())
		return
	}
	SuccessResp(c, nil)
	return
}

type Response struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Router  string      `json:"router"`
	Data    interface{} `json:"data"`
}

//FailResp 错误返回
func FailResp(c *gin.Context, Msg string) {
	resp := Response{
		Success: false,
		Router:  c.Request.URL.RequestURI(),
		Msg:     Msg,
	}
	c.JSON(200, &resp)
	c.Abort()
}

//SuccessResp 错误返回
func SuccessResp(c *gin.Context, data interface{}) {
	resp := Response{
		Success: true,
		Router:  c.Request.URL.RequestURI(),
		Data:    data,
	}
	c.JSON(200, &resp)
	c.Abort()
}
