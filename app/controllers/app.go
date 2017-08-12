package controllers

import (
	"github.com/revel/revel"
	"sasuke/app/service/sql"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	s := &sql.Handler{}
	s.FetchTableSchema()
	return c.Render()
}

func (c App) Execute() revel.Result {
	return c.Render()
}


func (c App) Result() revel.Result {
	return c.Render()
}
