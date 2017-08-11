package controllers

import (
	"bufio"
	"os"
	"log"

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
// ToDo 引数は構造体（？）で定義
func (c Config) Save(db string, host string, dbuser string, dbname string, password string) revel.Result {

	// ToDo: テスト接続して接続情報の有効性を確認する。

	// .envファイルへの書き込み文生成（テスト文字列）
	dbinfoString := "db="+ db + "\n" + "host=" + host + "\n" + "dbuser=" + dbuser + "\n" + "dbname=" + dbname + "\n" + "password" + password

	// .envファイルへの書き込み
	// ToDo: 既に書き込みがある場合、一度リフレッシュする
	write_file, _ := os.OpenFile(".env", os.O_WRONLY|os.O_APPEND, 0644)
	bw := bufio.NewWriter(write_file)
	_, err := bw.WriteString(dbinfoString)
	if err != nil {
		log.Fatal(err)
	}
	bw.Flush()

	return c.Redirect(App.Index) //ToDo: トップページにリダイレクト
}
