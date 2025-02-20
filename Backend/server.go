package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Connect to an in-memory SQLite database with foreign keys enabled
	db, err := gorm.Open(sqlite.Open("local.db?_pragma=foreign_keys(1)"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	
	//Auto migrate models
	err = db.AutoMigrate(&Recipe{}, &Ingredient{}, &Instruction{}, &Category{}, &Difficulty{})
	
	fmt.Println("Database connected successfully:", db)
	
	seed(db)

	router := mux.NewRouter()
	router.HandleFunc("/", mainHandler)

	log.Fatal(http.ListenAndServe(":3005", router))
}

func mainHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello World!")
}