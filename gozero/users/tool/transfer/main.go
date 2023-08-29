package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	jsonUser = "users.json"
	jsonPost = "posts.json"
)

func main() {
	users := GetStructFromUserJson(jsonUser, []User{})
	InsertUserObj(users)

	posts := GetStructFromPostJson(jsonPost, []Post{})
	InsertPostObj(posts)

	// PrintUserJson(users)
	// PrintPostJson(posts)
}

func GetStructFromUserJson(file string, tstruct []User) []User {
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("read file failed")
		panic(err)
	}
	return userMarshaler(content, tstruct)
}
func userMarshaler(content []byte, target []User) []User {
	err := json.Unmarshal(content, &target)
	if err != nil {
		fmt.Println("json unmarshal error")
		panic(err)
	}
	return target
}
func PrintUserJson(target []User) {
	jsonData, err := json.MarshalIndent(target, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}

func GetStructFromPostJson(file string, tstruct []Post) []Post {
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("read file failed")
		panic(err)
	}
	return postMarshaler(content, tstruct)
}
func postMarshaler(content []byte, target []Post) []Post {
	err := json.Unmarshal(content, &target)
	if err != nil {
		fmt.Println("json unmarshal error")
		panic(err)
	}
	return target
}
func PrintPostJson(target []Post) {
	jsonData, err := json.MarshalIndent(target, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}
