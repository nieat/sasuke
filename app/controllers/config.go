package controllers

import ("github.com/revel/revel")

type App struct {
	*revel.Controller
}

// 初期設定入力ページ
func (c App) Index() revel.Result {
	return c.Render()
}

// 初期設定確認ページ
func (c App) Confirm() revel.Result{
	return c.Render()
}

// 初期設定保存
func (c App) Save() revel.Result {
	return  c.Render()
}