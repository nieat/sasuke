package controllers

import (
	"github.com/revel/revel"
	"sasuke/app/service/sql"
	"strings"
	"fmt"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	s := &sql.Handler{}
	table_name_slice, column_name_slice := s.FetchTableSchema()
	// tableのダブり値を削除
	m := make(map[string]bool)
	uniq := [] string{}

	for _, ele := range table_name_slice {
		if !m[ele] {
			m[ele] = true
			uniq = append(uniq, ele)
		}
	}

	// var table_schema [ len(uniq) ][len(column_name_slice)]string

	//for i:=0 ;i < len(column_name_slice) ;i++  {
	//
	//	fmt.Printf(table_name_slice[i])
	//	fmt.Printf(column_name_slice[i])
	//	before_index := i -1
	//	// テーブル名変更された時
	//	if table_name_slice[i] != table_name_slice[before_index] {
	//		table_schema[i][0] = table_name_slice[i]
	//		table_schema[i] = append(table_schema[i], column_name_slice[i])
	//	}
	//}
	return c.Render(table_name_slice, column_name_slice)
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
				select_columns = select_columns + table + "." + col + ","
			}
			r_col_sl := strings.Split(relation_columns, ",")
			for _, col := range r_col_sl {
				select_columns = select_columns + relation + "." + col + ","
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
