package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"ns/filedir/models"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {

	xmlFile, err := os.Open("C:/project/practice/GO/updateFilePath/")
	if err != nil {
		fmt.Println(err.Error())
	}

	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println(err.Error())
	}

	var fileConfig models.FileConfig

	err = xml.Unmarshal(byteValue, &fileConfig)
	if err != nil {
		fmt.Println(err.Error())
	}

	baselocation := fileConfig.BaseLocation
	mainfile, err := os.Create(baselocation + fileConfig.MainFile)
	if err != nil {
		fmt.Println(err)
	}

	var actualfiles []string
	var totalfiles []string
	basepath := path.Base(baselocation)
	err = filepath.Walk(baselocation,
		func(filedirectory string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() || filepath.Ext(filedirectory) == ".log" {

			} else {
				ifdirequal := filepath.Base(filedirectory)
				if ifdirequal == basepath {

				} else if filepath.Base(filedirectory) == fileConfig.BatFile || filepath.Base(filedirectory) == fileConfig.MainFile {

				} else {
					actualfiles = strings.Split(filedirectory, basepath)
					totalfiles = append(totalfiles, basepath+actualfiles[1]+"\n")
					fmt.Println("\\i", basepath+actualfiles[1])
				}
			}
			var fileinbyte []byte
			for _, eachfile := range totalfiles {
				fileinbyte = append(fileinbyte, []byte(eachfile)...)
			}
			err = ioutil.WriteFile(mainfile.Name(), []byte(fileinbyte), 0777)
			if err != nil {
				fmt.Println(err.Error())
			}

			return nil
		})
	if err != nil {
		fmt.Println(err.Error())
	}
}
