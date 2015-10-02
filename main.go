package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/daviddengcn/ljson"
)

func main() {
	os.Exit(jsonfmt())
}

func jsonfmt() int {
	if len(os.Args) != 2 {
		usage()
		return 1
	}
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("err: %s", err)
		return 1
	}

	fi, err := file.Stat()
	if err != nil {
		fmt.Printf("err: %s", err)
		return 1
	}

	j := ljson.NewDecoder(file)
	data := map[string]interface{}{}
	if err := j.Decode(&data); err != nil {
		fmt.Printf("err: %s", err)
		return 1
	}

	if err := file.Close(); err != nil {
		fmt.Printf("err: %s", err)
		return 1
	}

	jsonStr, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("err: %s", err)
		return 1
	}

	if err := ioutil.WriteFile(path, jsonStr, fi.Mode()); err != nil {
		fmt.Printf("err: %s", err)
		return 1
	}

	return 0
}

func usage() {
	fmt.Printf(
		"jsonfmt <file>\n\n" +
			"  takes a loosely formatted json file and formats it to strict json\n")
}
