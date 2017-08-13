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
	return c.Render()
}


func (c App) Result() revel.Result {
	return c.Render()
}
