package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/cbot918/http-db-template/config"

	"github.com/cbot918/liby/util"
	"github.com/gofiber/fiber"
	_ "github.com/lib/pq"
)

const (
	PORT = ":5455"
)

func main(){
	c := config.NewConfig()

	app := fiber.New()
	
	app.Get("/psql", func (ctx *fiber.Ctx) {
		container_name := "http-db-template_postgres_1"
		network := "http-db-template_template"
		s1 := fmt.Sprintf("docker inspect %s | grep IPA", container_name)
		s2 := fmt.Sprintf( "docker run -it --rm --network %s postgres psql -h [查出來的ip] -U postgres -W",network)
		sendString := fmt.Sprintf("\n以下方式連入資料庫:\n%s\n%s\n\n",s1,s2)
		ctx.SendString(sendString)
	})

	app.Get("/psql/ping", func (ctx *fiber.Ctx){	
		connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",c.Host,c.Port,c.User,c.Password,c.Dbname)
		conn, err := sql.Open("postgres",connStr); util.Checke(err, "sql open failed"); 
		if err := conn.Ping(); err != nil { log.Fatal(err) }
	
		if err :=conn.Ping(); err != nil {
			fmt.Println("ping failed")
			panic(err)
		}
		fmt.Println("local: ping db success")
		
		ctx.SendString(fmt.Sprintf("res: %s created and ping %s success\n",c.Dbname,c.Dbname))
	})
	
	app.Get("/mongo", func (ctx *fiber.Ctx) {
		container_name := "http-db-template_mongo_1"
		user := "yale"
		pw := "12345"
		db := "testdb"
		
		s1 := fmt.Sprintf("$ docker exec -it %s bash",container_name)
		s2 := fmt.Sprintf("$ mongosh -u %s -p %s",user ,pw)
		s3 := fmt.Sprintf("$ use testdb")
		s4 := fmt.Sprintf("$ db.createUser({ user: '%s', 'pwd': '%s', roles: [{ role: 'readWrite', db: '%s' }] })", user, pw, db)
		
		n1 := fmt.Sprintf("開另一個視窗給nodejs")
		n2 := fmt.Sprintf("$ yarn")
		n3 := fmt.Sprintf("$ node mongo.js")

		v1 := fmt.Sprintf("進mongosh視窗")
		v2 := fmt.Sprintf("$ show dbs")
		v3 := fmt.Sprintf("$ show tables")
		v4 := fmt.Sprintf("$ db.user.find()")

		sendString := fmt.Sprintf("\n以下方式連入資料庫:\n%s\n%s\n%s\n%s\n\n%s\n%s\n%s\n\n%s\n%s\n%s\n%s\n\n",s1,s2,s3,s4,n1,n2,n3,v1,v2,v3,v4)
		ctx.SendString(sendString)
	})


	app.Listen(c.WebPort)
}






