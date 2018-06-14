package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

func CreateDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.MkdirAll(path, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func CreateDirs(paths ...string) (err error) {
	for _, path := range paths {
		if err = CreateDir(path); err != nil {
			return
		}
	}
	return
}

func CreateFile(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		var file *os.File
		if file, err = os.Create(path); err != nil {
			return err
		}
		defer file.Close()
	}
	return nil
}

func CreateFiles(paths ...string) (err error) {
	for _, path := range paths {
		if err = CreateFile(path); err != nil {
			return
		}
	}
	return
}

func RemoveDir(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return err
	}
	return nil
}

func RemoveDirs(paths ...string) (err error) {
	for _, path := range paths {
		if err = RemoveDir(path); err != nil {
			return
		}
	}
	return
}

func RemoveFile(path string) error {
	if err := os.Remove(path); err != nil {
		return err
	}
	return nil
}

func RemoveFiles(paths ...string) (err error) {
	for _, path := range paths {
		if err = RemoveFile(path); err != nil {
			return
		}
	}
	return
}

func WriteFile(path string, data []byte) error {
	var err = ioutil.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadFile(path string) ([]byte, error) {
	var data, err = ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ReadCSV(path string, mock interface{}) (vs []interface{}, err error) {
	var file, _ = os.Open(path)
	var csvReader = csv.NewReader(file)
	var m = reflect.ValueOf(mock)

	for {
		var v = make([]interface{}, m.NumField())
		var record, err = csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if len(record) != m.NumField() {
			return nil, fmt.Errorf("Number of fields does not match [record]")
		}
		for i := 0; i < m.NumField(); i++ {
			var f = m.Field(i)
			if f.Kind() != reflect.String {
				return nil, fmt.Errorf("Field type does not match [string]")
			}
			v[i] = record[i]
		}
		vs = append(vs, v)
	}
	return vs, nil
}

func GetDirElems(path string) ([]string, error) {
	var err error

	var files []os.FileInfo
	if files, err = ioutil.ReadDir(path); err != nil {
		return nil, err
	}

	var sources []string
	for _, file := range files {
		if file != nil && !file.IsDir() {
			sources = append(sources, file.Name())
		}
	}

	return sources, nil
}
