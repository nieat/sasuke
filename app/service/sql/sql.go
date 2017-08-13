package sql

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"database/sql"
	"sasuke/app/service/file"
	"os"
)

type Handler struct {
}


func (h *Handler)FetchTableSchema ()([]string,[]string){

	// Todo: Connect()
	f := &file.Handler{}
	f.LoadEnv();

	db 			:= os.Getenv("db")
	host 		:= os.Getenv("host")
	port 		:= os.Getenv("port")
	dbuser 		:= os.Getenv("dbuser")
	dbname 		:= os.Getenv("dbname")
	password 	:= os.Getenv("password")

	mysql, error := sql.Open(db,dbuser+":"+password+"@tcp("+host+":"+port+")/"+"information_schema")
	if error != nil {
		log.Fatal("open erro: %v", error)
	}
	defer mysql.Close()


	main_query :=`select table_name,column_name from columns where table_schema="`+dbname+ `";`
	rows, err := mysql.Query(main_query)
	defer rows.Close()

	if err != nil {
		panic(err.Error())
	}

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



// SQL Queryを実行し、カラム名・結果レコードを返す
func(h *Handler) ExecuteQuery(query string)([]string, [][]interface{}){

	// Todo: Connect()
	f := &file.Handler{}
	f.LoadEnv();

	db 			:= os.Getenv("db")
	host 		:= os.Getenv("host")
	port 		:= os.Getenv("port")
	dbuser 		:= os.Getenv("dbuser")
	dbname 		:= os.Getenv("dbname")
	password 	:= os.Getenv("password")

	mysql, err := sql.Open(db,dbuser+":"+password+"@tcp("+host+":"+port+")/"+dbname)
	if err != nil {
		log.Fatal("open erro: %v", err)
	}
	defer mysql.Close()


	// query実行
	rows, err := mysql.Query(query)
	if err != nil{
		panic(err.Error())
	}
	defer rows.Close()

	// オブジェクトから値を取り出す
	// rows.Scanで値を取り出すために[]interface{}を定義
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}
	count := len(columns)
	valuePtrs :=  make([]interface{}, count)

	records := make([][]interface{}, 0)


	for rows.Next() {
		values := make([]interface{}, count)

		for i,_ := range columns{
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)

		for i,_ := range columns{
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if(ok){
				v = string(b)
			}else{
				v = val
			}
			values[i] = v
		}
		records = append(records, values)
	}

	return columns, records;
}

