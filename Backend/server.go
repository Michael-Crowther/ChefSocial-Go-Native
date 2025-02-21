package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Declare a global DB variable
var db *gorm.DB

func main() {
	var err error
	// Connect to an in-memory SQLite database with foreign keys enabled
	db, err = gorm.Open(sqlite.Open("local.db?_pragma=foreign_keys(1)"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	
	// Auto migrate models
	err = db.AutoMigrate(&Recipe{}, &Ingredient{}, &Instruction{}, &Category{}, &Difficulty{})
	if err != nil {
		panic("failed to auto migrate: " + err.Error())
	}
	
	fmt.Println("Database connected successfully:", db)
	
	seed(db)

	router := mux.NewRouter()
	router.HandleFunc("/", mainHandler)
	router.HandleFunc("/recipes", recipeHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":3005", router))
}

func mainHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello World!")
}

func recipeHandler(w http.ResponseWriter, r *http.Request) {
	var recipes []Recipe
	
	// Fetch recipes from the database
	if err := db.Find(&recipes).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set header and return the recipes as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}