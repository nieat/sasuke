package filehandler

import (
	"os"
	"bufio"
	"log"
)

type Handler struct{
}

// ファイルを新規作成（既に存在する場合は更新）し、文字列を書き込む
func(h *Handler) CreateWriteFile(filename string, write_string string){
	_, cerr := os.Create(filename) //既に存在する場合は上書き
	if cerr != nil{
		log.Fatal(cerr)
	}

	write_file, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	bw := bufio.NewWriter(write_file)
	_, werr := bw.WriteString(write_string)
	if werr != nil {
		log.Fatal(werr)
	}
	bw.Flush()
}
