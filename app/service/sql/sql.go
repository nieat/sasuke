package sql

import (
	"github.com/gocraft/dbr"
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

	conn, err := dbr.Open(db,dbuser+":"+password+"@tcp("+host+":"+port+")/"+dbname,nil)
	if err != nil {
		panic(err)
	}
	session := conn.NewSession(nil)
	var tables []string
	session.SelectBySql("show tables").Load(&tables)

	//var table_schema [][]string
	mysql, error := sql.Open(db,dbuser+":"+password+"@tcp("+host+":"+port+")/"+dbname)
	if error != nil {
		log.Fatal("open erro: %v", err)
	}

	for i := 0;i <len(tables) ;i ++  {
		query := "SHOW COLUMNS FROM "+tables[i]

		rows, err := mysql.Query(query)
		defer rows.Close()
		if err != nil {
			log.Fatal( err)
		}

		fmt.Print(rows);
		for rows.Next() {
			var Field string
			fmt.Printf(Field)
		}

	}

	//fmt.Println(table_columns[:]);


	// 内容を整形する



	return
}

	func getDB() (db *sql.DB) {
	db, err := sql.Open("mysql", "USER:PASSWORD/DB_NAME")
		if err != nil {
		log.Fatal("open erro: %v", err)
		}
	return db
	}

//func loadEnv()  {
//	err := godotenv.Load()
//	if err != nil {
//		log.Fatal("Error loading .env file")
//	}
//}
