package main

import "strings"

func IsFirstMsg(msg []byte) bool {
	flag := "_init_user_name"
	return strings.Contains(string(msg), "\""+flag+"\":")
}
