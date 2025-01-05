package controllers

import (
	dao "GODI/models/DAO"
	"GODI/services"
)

func NewDITestController() {
	test := dao.DITestImpl{}
	services.NewDiTest(test)
}
