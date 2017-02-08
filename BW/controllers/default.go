package controllers

import (
	"github.com/astaxie/beego"

)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	c.TplNames = "home.html"
  
}


func (c *MainController) Citizen() {
	c.TplNames = "citizen.html"
}

func (c *MainController) Hobby() {

	c.TplNames = "hobby.html"
 }

func (c *MainController) Moment() {

	c.TplNames = "moment.html"
 }


func (c *MainController) Sign() {

	c.TplNames = "sign.html"
 }
