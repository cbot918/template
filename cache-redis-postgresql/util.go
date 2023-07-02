package main

import (
	"encoding/json"
	"fmt"
)

var log = fmt.Println

var logf = fmt.Printf

var logj = func(obj interface{}) {
	bytes, _ := json.MarshalIndent(obj, "", "  ")
	fmt.Println(string(bytes))
}
