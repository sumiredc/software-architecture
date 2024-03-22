package controllers

import (
	"net/http"

	"github.com/Greek-Academy/software-architecture/clean-architecture/application/use_cases"
	"github.com/Greek-Academy/software-architecture/clean-architecture/application/use_cases/requests"
	"github.com/Greek-Academy/software-architecture/clean-architecture/application/use_cases/responses"
)

type MuhohoControllerI interface {
	Handle() responses.MuhohoResponseI
}

type muhohoController struct {
	req *http.Request
}

func NewMuhohoController(req *http.Request) MuhohoControllerI {
	return &muhohoController{req}
}

func (c *muhohoController) Handle() responses.MuhohoResponseI {
	// 入力データの生成
	req := requests.NewMuhohoRequest(c.req)

	// ビジネスロジック
	mf := use_cases.NewMakeFace(req)

	// 出力データを返却
	return mf.Handle()
}
