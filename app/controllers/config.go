package controllers

import (
	"os"
	"bufio"
	//"log"

	"github.com/revel/revel"
)

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

	// Formパラメータの取得
	db 		 := c.Params.Form.Get("db")
	host	 := c.Params.Form.Get("host")
	dbuser	 := c.Params.Form.Get("dbuser")
	dbname 	 := c.Params.Form.Get("dbname")
	password := c.Params.Form.Get("password")

	// ToDo: テスト接続して接続情報の有効性を確認する。

	// .envファイルへの書き込み文生成
	dbinfoString := "db="+ db + "\n" +
					"host=" + host + "\n" +
					"dbuser=" + dbuser + "\n" +
					"dbname=" + dbname + "\n" +
					"password=" + password
	CreateWriteFile(".env", dbinfoString)
	return c.Redirect(App.Index)
}

// ファイルを新規作成（既に存在する場合は更新）し、文字列を書き込む
func CreateWriteFile(filename string, write_string string){
	_, cerr := os.Create(filename) //既に存在する場合は上書き
	if cerr != nil{
		panic(cerr)
	}

	write_file, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	bw := bufio.NewWriter(write_file)
	_, werr := bw.WriteString(write_string)
	if werr != nil {
		panic(werr)
	}
	bw.Flush()
}
