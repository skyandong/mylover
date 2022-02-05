package main

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
)

const catalogue = "files"

func main() {
	// 获取目录文件
	fis, err := ioutil.ReadDir(catalogue)
	if err != nil {
		log.Fatalf("read catalogue error: Directory name: %v, err: %v", catalogue, err)
	}

	//创建文件
	document, err := os.Create("佩奇小猪猪.csv")
	if err != nil {
		log.Fatalf("create aggregate documents error: %v", err)
	}
	defer func() {
		if err = document.Close(); err != nil {
			log.Fatalf("close aggregate document error")
		}
	}()
	writer := csv.NewWriter(document)

	// 获取 csv 文件
	for _, f := range fis {
		fullName := path.Join(catalogue, f.Name())
		csvFile, err := os.Open(fullName)
		if err != nil {
			log.Fatalf("open file error: file name: %v, err: %v", f.Name(), err)
		}
		reader := csv.NewReader(csvFile)
		if err = writeFiles(writer, reader); err != nil {
			log.Fatalf("write file error: %v", err)
		}
		if err = csvFile.Close(); err != nil {
			log.Fatalf("close csv file error: %v", err)
		}
	}
}

func writeFiles(writer *csv.Writer, reader *csv.Reader) error {
	for {
		row, err := reader.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		if err = writer.Write(row); err != nil {
			log.Fatalf("can not write, err is %+v", err)
		}
	}
	writer.Flush()
	return nil
}
