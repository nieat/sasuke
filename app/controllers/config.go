package controllers

import ("github.com/revel/revel")

type Config struct {
	*revel.Controller
}

// 初期設定入力ページ
func (c Config) Index() revel.Result {
	return c.Render()
}

// 初期設定確認ページ
func (c Config) Confirm() revel.Result{
	return c.Render()
}

// 初期設定保存
func (c Config) Save() revel.Result {
	return  c.Render()
}
