package main

type User struct {
	ID struct {
		Oid string `json:"$oid"`
	} `json:"_id"`
	Pic       string `json:"pic"`
	Followers []struct {
		Oid string `json:"$oid"`
	} `json:"followers"`
	Following []struct {
		Oid string `json:"$oid"`
	} `json:"following"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	V        int    `json:"__v"`
}
