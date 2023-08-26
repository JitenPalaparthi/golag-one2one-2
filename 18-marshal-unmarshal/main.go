package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/BurntSushi/toml"
	yaml "gopkg.in/yaml.v3"
)

var (
	filename string
)

func main() {

	flag.StringVar(&filename, "filename", "sample.json", "--filename=sample.yaml || -filename=sample.yaml")
	//fmt.Println(os.Args)
	flag.Parse()
	//filename := "sample.yaml"
	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	samplMap := make(map[string]any)

	if filepath.Ext(filename) == ".json" {
		err = json.Unmarshal(bytes, &samplMap)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else if filepath.Ext(filename) == ".yaml" || filepath.Ext(filename) == ".yml" {
		err = yaml.Unmarshal(bytes, &samplMap)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else if filepath.Ext(filename) == ".toml" {
		err = toml.Unmarshal(bytes, &samplMap)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	for k, v := range samplMap {
		fmt.Println("Key:", k /*"Value:",, v*/, "Type:", reflect.TypeOf(v))
	}

	inmap, _ := samplMap["widget"]
	val, ok2 := inmap.(map[string]any)["image"]
	fmt.Println(reflect.TypeOf(val))
	if ok2 {
		switch val.(type) {
		case map[string]any:
			for k, v := range val.(map[string]any) {
				fmt.Println("Key:", k, "Value:", v, "Type:", reflect.TypeOf(v))
			}

		case []string:
			for _, v := range val.([]string) {
				fmt.Println(v)
			}

		case []any:
			for _, v := range val.([]any) {
				fmt.Println(v)
			}

		default:
			fmt.Println("Undefined type")
		}

	}

}
