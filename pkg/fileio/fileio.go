package fileio

import (
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"
)

/*Test folder*/
var TestFolder = "testdata"

/*Init function runs when package is imported*/
func init() {
	if err := os.Mkdir(TestFolder, 0755); err != nil {
		log.Println(err.Error())
	}
}

/*Create file*/
func createFile() {
	f, err := os.Create(filepath.Join(TestFolder, "test.txt"))
	defer f.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	f.WriteString("Hello Go")
}

/*read data from file*/
func readFile() {
	in, err := os.Open(filepath.Join(TestFolder, "test.txt"))
	if err != nil {
		log.Fatal(err.Error())
	}
	data, err := io.ReadAll(in)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(string(data))
}

/*Read file using bufio*/
func readFileWithBuffer() {
	in, err := os.Open(filepath.Join(TestFolder, "test.txt"))
	if err != nil {
		log.Fatal(err.Error())
	}

	bufferReader := bufio.NewScanner(in)

	for bufferReader.Scan() {
		log.Println(bufferReader.Text())
	}

}

/*run all function*/
func RunFileio() {
	createFile()
	readFile()
	readFileWithBuffer()
}
