package controllers

import (
	"github.com/revel/revel"
	"strings"
	"fmt"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Execute() revel.Result {
	// todo query結果を結果画面に移す?
	table     := c.Params.Form.Get("table_name")
	relation  := c.Params.Form.Get("relation_name")
	columns   := c.Params.Form.Get("column_names")
	relation_columns := c.Params.Form.Get("relation_column_names")
	c.Validation.Required(table).Message("table_name is required.")
	if (c.Validation.HasErrors()) {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}
	// todo 消さない。他でfmt使う時に消す。
	fmt.Print(relation)
	// joinなしの場合
	query := ""
	select_columns := ""
	// joinなしの場合
	if (relation == "") {
		// カラム指定がない場合
		if (columns == "") {
			select_columns = "*"
		// カラム指定がある場合
		} else {
			select_columns = columns
		}
	// 全選択の場合
		query = "select " + select_columns + " from " + table
	// joinありの場合
	} else {
		// カラム指定がない場合
		if (columns == "" && relation_columns == "") {
			select_columns = "*"
		// カラム指定がある場合
		} else {
			col_sl := strings.Split(columns, ",")
			for _, col := range col_sl {
				select_columns += table + "." + col + " as " + table + "_" + col + ","
			}
			r_col_sl := strings.Split(relation_columns, ",")
			for _, col := range r_col_sl {
				select_columns += relation + "." + col + " as " + relation + "_" + col + ","
			}
			csc := []rune(select_columns)
			select_columns = string(csc[:(len(csc) - 1)])
		}
		// 外部キーを取得
		sc := []rune(table)
		foreign_key := string(sc[:(len(sc) - 1)]) + "_id"
		query = "select " + select_columns + " from " + table + " left join " + relation + " on " + table + ".id " + "= " + relation + "." + foreign_key
	}
	// todo sqlを投げて、その結果を渡す
	return c.Render(query, columns)
	// return c.Redirect(App.Result, query)
}


func (c App) Result() revel.Result {
	return c.Render()
}
