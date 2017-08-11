package controllers

import (
	"github.com/revel/revel"
	"sasuke/app/service/file"
)

type Config struct {
	*revel.Controller
}

// DB接続情報
var(
	db			string
	host		string
	dbuser		string
	dbname		string
	password	string
)

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

	// Formパラメータの取得
	db 		 := c.Params.Form.Get("db")
	host	 := c.Params.Form.Get("host")
	dbuser	 := c.Params.Form.Get("dbuser")
	dbname 	 := c.Params.Form.Get("dbname")
	password := c.Params.Form.Get("password")

	// バリデーション
	c.Validation.Required(host)
	c.Validation.Required(dbuser)
	c.Validation.Required(dbname)
	c.Validation.Required(password)
	if c.Validation.HasErrors(){
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Config.Index)
	}

	// ToDo: テスト接続して接続情報の有効性を確認する。

	// .envファイルへの書き込み文生成
	dbconfig := 	"db="+ db + "\n" +
					"host=" + host + "\n" +
					"dbuser=" + dbuser + "\n" +
					"dbname=" + dbname + "\n" +
					"password=" + password
	f := &file.Handler{}
	f.CreateWriteFile(".env", dbconfig)

	return c.Redirect(App.Index)
}
