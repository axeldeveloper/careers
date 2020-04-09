package models

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

	db.Debug().AutoMigrate(&Heroe{}, &Connection{}) //Database migration
	db.Debug().Model(&Connection{}).AddForeignKey("heroe_id", "Heroes(id)", "CASCADE", "CASCADE")

	Morcego := Heroe{
		Name:         "Batman",
		Intelligence: 100,
		FullName:     "Bruce Wayne",
		Power:        47,
		Occupation:   "Businessman",
		Image:        "https://firebasestorage.googleapis.com/v0/b/axeldbcloud.appspot.com/o/Batman-lista.jpg?alt=media&token=5a452bbc-8151-4714-b32e-2147b5239b62",
		Relatives:    4,
		Type:         "Heroes",
		Connection: []Connection{
			{Group: "Justice League"},
			{Group: "Wayne Enterprises"},
			{Group: "Club of Heroes"},
			{Group: "formerly White Lantern Corps"},
			{Group: "Sinestro Corps"}}}

	db.Debug().Create(&Morcego)

}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
