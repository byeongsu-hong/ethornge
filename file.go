package ethornge

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

type pathParam struct {
	p string
}

func createFile(path *pathParam) {
	var _, err = os.Stat(path.p)

	if os.IsNotExist(err) {
		var file, err = os.Create(path.p)
		if err != nil {
			return
		}
		defer file.Close()
	}
}

func writeFile(path *pathParam, data []byte) {
	var err = ioutil.WriteFile(path.p, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func writeJSON(path *pathParam, v interface{}) {
	createFile(path)
	var data, err = json.MarshalIndent(v, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	writeFile(path, data)
}

func removeFile(path *pathParam) {
	if err := os.Remove(path.p); err != nil {
		log.Fatal(err)
	}
}

func readFile(path *pathParam) []byte {
	var data, err = ioutil.ReadFile(path.p)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func readJSON(path *pathParam, v interface{}) {
	if err := json.Unmarshal(readFile(path), &v); err != nil {
		log.Fatal(err)
	}
}

func readCSV(path *pathParam, mock interface{}) (vs []interface{}) {
	var csvReader = csv.NewReader(bytes.NewReader(readFile(path)))
	var m = reflect.ValueOf(mock)

	for {
		var v = make([]interface{}, m.NumField())
		var record, err = csvReader.Read()
		if len(record) != m.NumField() {
			log.Fatal("Number of fields does not match [record]")
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < m.NumField(); i++ {
			var f = m.Field(i)
			if f.Kind() != reflect.String {
				log.Fatal("Field type does not match [string]")
			}
			v[i] = record[i]
		}
		vs = append(vs, v)
	}
	return
}
