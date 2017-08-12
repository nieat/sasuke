package sql

import (
	//"github.com/gocraft/dbr"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/joho/godotenv"
	//"log"
	"fmt"
	"log"
	"database/sql"
)

type Handler struct {
}

//var (
//	db			string
//	host		string
//	dbuser		string
//	dbname		string
//	password	string
//	port        string
//)

func (h *Handler)FetchTableSchema () {
	//loadEnv();
	db := "mysql"
	host :="localhost"
	port := "3306"
	dbuser := "root"
	dbname := "SASUKE_TEST"
	password := "password"

	//conn, err := dbr.Open(db,dbuser+":"+password+"@tcp("+host+":"+port+")/use information_schema",nil)
	//if err != nil {
	//	panic(err)
	//}
	//session := conn.NewSession(nil)
	//var tables []string
	//session.SelectBySql("show tables").Load(&tables)

	//var table_schema [][]string
	mysql, error := sql.Open(db,dbuser+":"+password+"@tcp("+host+":"+port+")/"+"information_schema")
	if error != nil {
		log.Fatal("open erro: %v", error)
	}

		//use_query := "use information_schema;"
		main_query := "select table_name,column_name from columns where table_schema=' "+dbname+ `';`

		//mysql.Query(use_query)
		rows, err := mysql.Query(main_query)

	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns() // カラム名を取得
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))
	fmt.Print(len(values))
	//  rows.Scan は引数に `[]interface{}`が必要.

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
		//defer rows.Close()
		//
		//if err != nil {
		//	log.Fatal( err)
		//}
		//
		//columns, err := rows.Columns()
		//fmt.Print(columns)
		//
		//values := make([]sql.RawBytes, len(columns))
		//scanArgs := make([]interface{}, len(values))
		//
		//for i := range values{
		//	scanArgs[i] = &values[i]
		//}
		//
		//for rows.Next() {
		//	err = rows.Scan(scanArgs...)
		//	if err != nil{
		//		panic(err.Error())
		//	}
		//
		//	var value string
		//	for _, col := range values{
		//		if col == nil{
		//			value = "NULL"
		//		}else{
		//			value = string(col)
		//		}
		//		fmt.Println(value)
		//	}
		//}
		//fmt.Println("---------------")



	//fmt.Println(table_columns[:]);


	// 内容を整形する



	return
}



//func loadEnv()  {
//	err := godotenv.Load()
//	if err != nil {
//		log.Fatal("Error loading .env file")
//	}
//}
