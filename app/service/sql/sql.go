package sql

import (
	_ "github.com/go-sql-driver/mysql"
	//"github.com/joho/godotenv"
	"log"
	"database/sql"
)

type Handler struct {
}

func (h *Handler)FetchTableSchema ()([]string,[]string){

	//loadEnv();
	db := "mysql"
	host :="localhost"
	port := "3306"
	dbuser := "root"
	dbname := "SASUKE_TEST"
	password := "password"

	mysql, error := sql.Open(db,dbuser+":"+password+"@tcp("+host+":"+port+")/"+"information_schema")
	if error != nil {
		log.Fatal("open erro: %v", error)
	}

		main_query :=`select table_name,column_name from columns where table_schema="`+dbname+ `";`
		rows, err := mysql.Query(main_query)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	type TableSchema struct {
		table_name string
		column_name string
	}

	var table_name_slice []string
	var column_name_slice []string

	for rows.Next() {
		var table_name string
		var column_name string
		if err := rows.Scan(&table_name, &column_name); err != nil {
			log.Fatal(err)
		}

		table_name_slice = append(table_name_slice,table_name)
		column_name_slice = append(column_name_slice,column_name)
	}

	return table_name_slice,column_name_slice
}



//func loadEnv()  {
//	err := godotenv.Load()
//	if err != nil {
//		log.Fatal("Error loading .env file")
//	}
//}
