package controllers

import (
	"github.com/beego/beego/v2/client/orm"
)

type BaseController struct {
	//rep repository.Repository
	DB orm.Ormer
}
