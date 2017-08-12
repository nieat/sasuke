package sql

import (
	"github.com/gocraft/dbr"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/joho/godotenv"
	//"log"
	"fmt"
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

	conn, err := dbr.Open(db,dbuser+":"+password+"@tcp("+host+":"+port+")/"+dbname,nil)
	if err != nil {
		panic(err)
	}
	sess := conn.NewSession(nil)
	var tables []string
	sess.SelectBySql("show tables").Load(&tables)

	//var builder [][]string
	//var table_columns [][]string
	for i := 0;i <len(tables) ;i ++  {
		//sess.SelectBySql("SHOW COLUMNS FROM"+tables[i]).Load(&builder)
		a := sess.SelectBySql("SHOW COLUMNS FROM"+tables[i])
		fmt.Print(tables[i],a)
	}


	// 内容を整形する



	return
}

//func loadEnv()  {
//	err := godotenv.Load()
//	if err != nil {
//		log.Fatal("Error loading .env file")
//	}
//}
