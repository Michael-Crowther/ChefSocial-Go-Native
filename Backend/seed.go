package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

func seed(db *gorm.DB){
	// Categories
	var categoryIDs []uint
	for _, catName := range Categories {
		category := Category{Name: string(catName)}
		db.Create(&category)
		categoryIDs = append(categoryIDs, category.ID)
	}

	// Difficulties
	var difficultyIDs []uint
	for _, diffName := range Difficulties {
		difficulty := Difficulty{Name: string(diffName)}
		db.Create(&difficulty)
		difficultyIDs = append(difficultyIDs, difficulty.ID)
	}

	//Recipes
	for i := 0; i < 10; i++ {
		recipe := Recipe{
			Name:       faker.Word() + " Recipe",
			ImageUrls:  faker.URL(),
			CategoryID:   categoryIDs[rand.Intn(len(categoryIDs))],
			DifficultyID: difficultyIDs[rand.Intn(len(difficultyIDs))],
			PrepTime:   uint(rand.Intn(20) + 10),
			CookTime:   uint(rand.Intn(40) + 20),
			TotalTime:  uint(rand.Intn(60) + 30),
			Servings:   uint(rand.Intn(6) + 1),
			Calories:   uint(rand.Intn(700) + 100),
			Rating:     rand.Float32() * 5,
			Author:     faker.Name(),
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		for i := 0; i < 5; i++ {
			recipe.Ingredients = append(recipe.Ingredients, Ingredient{
				Name: faker.Word(),
			})
		}

		for i := 1; i < 5; i++ {
			recipe.Instructions = append(recipe.Instructions, Instruction{
				Description: faker.Sentence(),
				StepNumber: uint(i),
			})
		}

		db.Create(&recipe)
	}
	log.Printf("Seeded DB")
}