package io

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)


func prepareTestDirTree(tree string) (string, error) {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		return "", fmt.Errorf("error creating temp directory: %v\n", err)
	}

	err = os.MkdirAll(filepath.Join(tmpDir, tree), 0755)
	if err != nil {
		os.RemoveAll(tmpDir)
		return "", err
	}

	return tmpDir, nil
}


func TestSearch(t *testing.T) {

	tmpDir, err := prepareTestDirTree("dir/to/walk/skip")
	if err != nil {
		fmt.Printf("unable to create test dir tree: %v\n", err)
		return
	}
	defer os.RemoveAll(tmpDir)
	os.Chdir(tmpDir)

	subDirToSkip := "skip"

	fmt.Println("On Unix:")
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() && info.Name() == subDirToSkip {
			fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
			return filepath.SkipDir
		}
		fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", tmpDir, err)
		return
	}
}



func TestSearchFileContent(t *testing.T) {

	folders:= []string{`D:\userdata\dingliu\git\nbi`,
		`D:\userdata\dingliu\git\NBI_3GPPNBIPM`,
		`D:\userdata\dingliu\git\NBI_INVENTORY`,
		`D:\userdata\dingliu\git\NBI_NBI3GC`,
		`D:\userdata\dingliu\git\NBI_NBI3GCOM`,
		`D:\userdata\dingliu\git\NBI_PMFILEMERGER`,
		`D:\userdata\dingliu\git\NBI_RESTDA`,
		`D:\userdata\dingliu\git\NBI_NBISNMP`,
		`D:\userdata\dingliu\git\nbives`,
		`D:\userdata\dingliu\git\restda_common`,
		`D:\userdata\dingliu\git\restda_fm`,
		`D:\userdata\dingliu\git\restda_interface`,
		`D:\userdata\dingliu\git\restda_pm`,
	}

	for _,folder:=range folders{
		if cherr := os.Chdir(folder);cherr!=nil{
			panic("Chdir" + folder + "failed" )
		}

		err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				//fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
				return nil
			}

			if info.IsDir() || !strings.HasSuffix(path, "pom.xml") {
				//fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
				return nil
			}


			//fmt.Printf("visited file: %q\n", path)
			file, err := os.OpenFile(path, os.O_RDONLY, 0644)
			if err != nil{
				log.Fatal("open file failed", err)
			}
			scanner := bufio.NewScanner(file)
			for scanner.Scan(){
				if strings.Contains(scanner.Text(), "cm-notification-lib"){
					fmt.Println(path, scanner.Text())
				}
			}
			file.Close()
			return nil
		})

		if err != nil {
			fmt.Printf("error walking the path %q: %v\n", folder, err)
			//return
		}
	}
}
