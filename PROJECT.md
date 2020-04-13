
# Project
	Supers API 1.0

## INFORMATION 
	- go version  `go1.12.7 linux/amd64`
	- go mod `controle da versão dos pacote`
		- export GO111MODULE=on 
	- PostgresSQL
	- Docker
	- swagger


# DOCKER POSTGRES
```bash
foo@bar:~$ docker/docker-compose up -d 
```

![PostgresSQL Docker](/screen/pg.png "PostgresSQL Docker")



## UUID POSTGRESQL - TYPES

```sql
-- uuid postgres
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- ENUM postgres
CREATE TYPE super_type AS ENUM (
	'Hero',
	'Villain');

```

# GORM "MODEL/BASE.GO" - DATABASE MIGRATION
```go
import (
	"fmt"
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB //database

func init() {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	dbUri := fmt.Sprintf("host=%s port=%s  user=%s dbname=%s sslmode=disable password=%s",
		dbHost,
		dbPort,
		username,
		dbName,
		password) //Build connection string

	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	//Database migration
	db.Debug().AutoMigrate(&Super{}, &Connection{}) 
	db.Debug().Model(&Connection{}).AddForeignKey("super_id", "Supers(id)", "CASCADE", "CASCADE")

	Morcego := Super{
		Name:         "Batman",
		Intelligence: 100,
		FullName:     "Bruce Wayne",
		Power:        47,
		Occupation:   "Businessman",
		Image:        "https://firebasestorage.googleapis.com/v0/b/axeldbcloud.appspot.com/o/Batman-lista.jpg?alt=media&token=5a452bbc-8151-4714-b32e-2147b5239b62",
		Relatives:    4,
		Type:         "Hero",
		Connection: []Connection{
			{Group: "Justice League"},
			{Group: "Wayne Enterprises"},
			{Group: "Club of Heroes"},
			{Group: "formerly White Lantern Corps"},
			{Group: "Sinestro Corps"}}}

	Jocker := Super{
		Name:         "The Joker",
		FullName:     "Unknown",
		Intelligence: 100,
		Power:        47,
		Occupation:   "Criminal · mass murderer",
		Image:        "https://vignette.wikia.nocookie.net/marvel_dc/images/5/58/Joker_0003.jpg/revision/latest/top-crop/width/360/height/450?cb=20140725171230",
		Relatives:    2,
		Type:         "Villain",
		Connection: []Connection{
			{Group: "League of Villainy"},
			{Group: "Legion of Doom,"},
			{Group: "Red Hood Gang"}}}

	db.Debug().Create(&Morcego)
	db.Debug().Create(&Jocker)

}

// set User's table name to be `Herois`
func (User) TableName() string {
	return "Herois"
 }

// Disable table name's pluralization globally
//db.SingularTable(true)


//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
```




# CHALLENGE  - DESAFIO	
API deve ser possível:


Operação                            | Status
------------                        | -------------
Cadastrar um Super/Vilão            | OK
Listar todos os Super's cadastrados | OK
Listar apenas os Super Heróis       | OK
Listar apenas os Super Vilões       | OK
Buscar por nome                     | OK
Buscar por 'uuid'                   | OK
Remover o Super                     | OK 




# SWAGGO SETUP

## INSTALL PACKAGES
- go get -u github.com/swaggo/swag/cmd/swag
- go get -u github.com/swaggo/http-swagger

## GENERATE SWAGGER DOCUMENTATION
	Let us divide this whole process of API documentation into 3 steps:
	Adding annotations in code
	Generating Swagger specs (swagger.json and swagger.yaml)
	Serving the Swagger UI using the specs generated in the previous step

## GENERATE SWAGGER.JSON

## IN YOUR PROJECT DIR (~/GOPATH/SRC/SUPER-XXXX-API NORMALLY)

```console
foo@bar:~$ swag init -g yourModel.go
Output
	019/12/02 19:19:43 Generate swagger docs....
	2019/12/02 19:19:43 Generate general API Info, search dir:./
	2019/12/02 19:19:43 create docs.go at  docs/docs.go
	2019/12/02 19:19:43 create swagger.json at  docs/swagger.json
	2019/12/02 19:19:43 create swagger.yaml at  docs/swagger.yaml

```
or 

```console
foo@bar:~$ swag init
Output
	019/12/02 19:19:43 Generate swagger docs....
	2019/12/02 19:19:43 Generate general API Info, search dir:./
	2019/12/02 19:19:43 create docs.go at  docs/docs.go
	2019/12/02 19:19:43 create swagger.json at  docs/swagger.json
	2019/12/02 19:19:43 create swagger.yaml at  docs/swagger.yaml

```

## RUN OR SYNCRONIZE
```console
foo@bar:~$ go mod -vendor or go mod -sync
foo@bar:~$ go run main.go
```


## REST API CODE FOR CRUD
Response Code	        | HTTP Operation	    | Description
------------------------| ----------------------|---------------------------------------
200 Success	            | GET POST, PUT, DELETE	| The request was received.
202 Accepted	        | POST, PUT, DELETE		| The request was received.
204 No Content	        | GET, PUT, DELETE		| The request was processed successfully
301 Moved Permanently	| GET	                | Resource has moved.
303 See Other	        | GET	                | Redirection.

## URL
[swagger](http://localhost:8484/swagger/index.html) - http://localhost:8484/swagger/index.html

[api](http://localhost:8484/swagger/index.html) - http://localhost:8484/api/supers

## TEST PROJECT WHITH SWAGGER UI
swagger 
![swagger ui complete](/screen/swagger-index-html.png "swagger ui complete")


## TEST PROCEDURE 

	1 - execute o comando docker 
	1.1 - Criar estensões e tipos no Postgres
	2 - sincronize os pacotes do go 
	3 - execute os comandos do Swagger 2.0 for Go
	4 - execute o projeto atravéz do main.go
	5 - teste avontade