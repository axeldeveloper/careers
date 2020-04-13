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

	CreateAndImport(db, false)

	//defer db.Close()

}

func CreateAndImport(db *gorm.DB, exec bool) {

	if exec {
		db.Debug().AutoMigrate(&Super{}, &Connection{}) //Database migration
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

}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}

func CloseDb() {
	defer db.Close()
}
