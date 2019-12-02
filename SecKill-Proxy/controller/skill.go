package controller

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/guxiaogang/SecKill-Proxy/service"
)

type SkillController struct {
	beego.Controller
}

func (p *SkillController) SecKill() {
	productId, err := p.GetInt("product_id")
	result := make(map[string]interface{})

	result["code"] = 0
	result["message"] = "success"

	defer func() {
		p.Data["json"] = result
		p.ServeJSON()
	}()

	if err != nil {
		result["code"] = service.ErrInvalidRequest
		result["message"] = "invalid product_id"
		return
	}

	source := p.GetString("src")
	authCode := p.GetString("authCode")
	secTime := p.GetString("time")
	nance := p.GetString("nance")

	secRequest := service.NewSecRequest()
	secRequest.AuthCode = authCode
	secRequest.Nance = nance
	secRequest.ProductId = productId
	secRequest.SecTime = secTime
	secRequest.Source = source
}

func (p *SkillController) SecInfo() {
	productId, err := p.GetInt("product_id")

	result := make(map[string]interface{})
	result["code"] = 0
	result["message"] = "success"

	defer func() {
		p.Data["json"] = result
		p.ServeJSON()
	}()

	if err != nil {
		result["code"] = service.ErrInvalidRequest
		result["message"] = "bad request."
		logs.Error("invalid request, get product_id failed, err:%v", err)
		return
	}

	data, code, err := service.SecInfoById(productId)

	if err != nil {
		result["code"] = code
		result["message"] = err.Error()
		logs.Error("invalid request, get product_id failed, err:%v", err)
	}

	result["data"] = data
}
