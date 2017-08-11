package controllers

import (
	"github.com/revel/revel"
	// "log"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Execute(table string, relation string, columns []string, options []string) revel.Result {
	return c.Redirect(App.Result)
}


func (c App) Result() revel.Result {
	return c.Render()
}
