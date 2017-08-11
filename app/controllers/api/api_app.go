package controllers

import (
	"github.com/revel/revel"
	// "log"
)

type ApiApp struct {
	*revel.Controller
}

func (c ApiApp) Execute(table string, relation string, columns []string, options []string) revel.Result {
	return c.Render()
}