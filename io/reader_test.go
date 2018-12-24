package io

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {

	reader := strings.NewReader("this is a good movie")
	limitReader := io.LimitReader(reader, 6)
	if _, err := io.Copy(os.Stdout, limitReader); err != nil {
		t.Fatalf("%s", err)
	}
}

func TestMultiReader(t *testing.T) {
	r1 := strings.NewReader("first reader ")
	r2 := strings.NewReader("second reader ")
	r3 := strings.NewReader("third reader\n")
	r := io.MultiReader(r1, r2, r3)

	if _, err := io.Copy(os.Stdout, r); err != nil {
		t.Fatal(err)
	}
}

func TestCreatTempFile(t *testing.T) {
	content := []byte("temporary file's content")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tmpfile.Name())
	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}


func TestSearchFileConet(t *testing.T) {

	folders:= []string{`D:\userdata\dingliu\git\nbi`,
		`D:\userdata\dingliu\git\NBI_3GPPNBIPM`,
		`D:\userdata\dingliu\git\NBI_INVENTORY`,
		`D:\userdata\dingliu\git\NBI_NBI3GC`,
		`D:\userdata\dingliu\git\NBI_NBI3GCOM`,
		`D:\userdata\dingliu\git\NBI_PMFILEMERGER`,
		`D:\userdata\dingliu\git\NBI_RESTDA`,
		`D:\userdata\dingliu\git\nbives`,
	}

	for _,root:=range folders{
		filepath.Walk(root, visit)

	}
}

func visit(path string, f os.FileInfo, err error) error{

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil{
		log.Fatal("open file failed", err)
	}

	defer file.Close()

	fmt.Println("visit:", path)
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		if strings.Contains(scanner.Text(), "cm-notification-lib"){
			fmt.Println(path, scanner.Text())
		}
	}

	if scanner.Err()!=nil{
		return scanner.Err()
	}

	return nil

}
