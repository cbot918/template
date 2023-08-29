package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/cbot918/liby/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	DB_URL = "root:12345@tcp(localhost:3308)/users"
)

var log = fmt.Println
var logf = fmt.Printf

func InsertUserObj(users []User) {
	if len(users) == 0 {
		log("no user data, quit program")
		os.Exit(1)
	}
	var stmt string
	/* insert users table*/
	stmt = GetCatStmt(users, "users")
	log(stmt)
	execInsert(stmt)

	/* insert follow table */
	stmt = GetCatStmt(users, "follow")
	log(stmt)
	execInsert(stmt)
	log("insert complete")
}

func InsertPostObj(posts []Post) {
	if len(posts) == 0 {
		log("no user data, quit program")
		os.Exit(1)
	}
	var stmt string
	/* insert posts table*/
	stmt = GetCatStmtPosts(posts, "posts")
	log(stmt)
	execInsert(stmt)

	/* insert post_like table*/
	stmt = GetCatStmtPosts(posts, "post_like")
	log(stmt)
	execInsert(stmt)

	/* insert comments table*/
	stmt = GetCatStmtPosts(posts, "comments")
	log(stmt)
	execInsert(stmt)

}

func GetCatStmtPosts(posts []Post, tableType string) (str string) {
	switch tableType {
	case "posts":
		{
			str = "INSERT INTO posts (id, title, body, posted_by, photo, created_at, updated_at) VALUES"
			for _, post := range posts {
				str += fmt.Sprintf("((UNHEX('%s'), '%s', '%s', (UNHEX('%s'), '%s', '%s','%s'),",
					util.GetUuidFill(post.ID.Oid, 32),
					post.Title,
					post.Body,
					util.GetUuidFill(post.PostedBy.Oid, 32),
					post.Photo,
					util.GetParsedTime(post.CreatedAt.Date.String()),
					util.GetParsedTime(post.UpdatedAt.Date.String()),
				)
			}
			str = str[:len(str)-1]
			str += ";"
		}
	case "post_like":
		{
			str = "INSERT INTO post_like (liked_user, target_post) VALUES"
			for _, post := range posts {
				if len(post.Likes) > 0 {
					for _, like := range post.Likes {
						str += fmt.Sprintf("((UNHEX('%s'), (UNHEX('%s')),",
							util.GetUuidFill(like.Oid, 32),
							util.GetUuidFill(post.ID.Oid, 32),
						)
					}
				}
			}
			str = str[:len(str)-1]
			str += ";"
		}
	case "comments":
		{
			str = "INSERT INTO comments (id, texts, posted_by, target_post) VALUES"
			for _, post := range posts {
				if len(post.Comments) > 0 {
					for _, comment := range post.Comments {
						str += fmt.Sprintf("((UNHEX('%s'), '%s', (UNHEX('%s'), (UNHEX('%s')),",
							util.GetUuidFill(comment.ID.Oid, 32),
							comment.Text,
							util.GetUuidFill(comment.PostedBy.Oid, 32),
							util.GetUuidFill(post.ID.Oid, 32),
						)
					}
				}
			}
			str = str[:len(str)-1]
			str += ";"
		}
	}
	return
}

func GetCatStmt(users []User, tableType string) (str string) {
	switch tableType {
	case "users":
		{
			str = "INSERT INTO users (id, email, name, password, pic) VALUES"
			for _, user := range users {
				str += fmt.Sprintf("(UNHEX('%s'),'%s','%s','%s','%s'),", util.GetUuidFill(user.ID.Oid, 32), user.Email, user.Name, user.Password, user.Pic)
			}
			str = str[:len(str)-1]
			str += ";"
		}
	case "follow":
		{
			str = "INSERT INTO follow (from_user,to_user) VALUES"
			for _, user := range users {
				if len(user.Followers) > 0 {
					for _, follower := range user.Followers {
						str += fmt.Sprintf("((UNHEX('%s'),(UNHEX('%s')),", util.GetUuidFill(user.ID.Oid, 32), util.GetUuidFill(follower.Oid, 32))
					}
				}
				if len(user.Following) > 0 {
					for _, following := range user.Following {
						str += fmt.Sprintf("((UNHEX('%s'),(UNHEX('%s')),", util.GetUuidFill(following.Oid, 32), util.GetUuidFill(user.ID.Oid, 32))
					}
				}
			}
			str = str[:len(str)-1]
			str += ";"
		}
	}
	return
}

func execInsert(stmt string) sql.Result {
	db := setupDb(DB_URL)
	err := db.Ping()
	if err != nil {
		fmt.Println("ping failed")
	}

	// stage1
	res, err := db.Exec(stmt)
	if err != nil {
		log("insert exec failed")
		panic(err)
	}
	log("insert stage1 success")
	return res
}

// func ReduceDup(t User) string {
// 	set := mapset.NewSet()
// 	for _, item := range t.Followers {
// 		set.Add(item.Oid)
// 	}
// 	g := set.ToSlice()
// 	return g[0].(string)
// }

func HaveFollower(target *User) bool {
	return len(target.Followers) > 0
}

func setupDb(dburl string) *sqlx.DB {
	db, err := sqlx.Connect("mysql", dburl)
	if err != nil {
		log("sqlx connect failed")
		panic(err)
	}
	return db
}
