package main

import "time"

type Post struct {
	ID struct {
		Oid string `json:"$oid"`
	} `json:"_id"`
	Likes []struct {
		Oid string `json:"$oid"`
	} `json:"likes"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Photo    string `json:"photo"`
	PostedBy struct {
		Oid string `json:"$oid"`
	} `json:"postedBy"`
	Comments []struct {
		ID struct {
			Oid string `json:"$oid"`
		} `json:"_id"`
		Text     string `json:"text"`
		PostedBy struct {
			Oid string `json:"$oid"`
		} `json:"postedBy"`
	} `json:"comments"`
	CreatedAt struct {
		Date time.Time `json:"$date"`
	} `json:"createdAt"`
	UpdatedAt struct {
		Date time.Time `json:"$date"`
	} `json:"updatedAt"`
	V int `json:"__v"`
}
