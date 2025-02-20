package main

import (
	"time"

	"gorm.io/gorm"
)

type Recipe struct {
	ID             uint           `gorm:"primaryKey"`
	Name           string         `gorm:"not null"`
	ImageUrls      []string       `gorm:"type:json"`
	CategoryID     uint           `gorm:"not null"`
	Category       CategoryType   `gorm:"type:text;not null;index"`
	Instructions   []Instruction  `gorm:"foreignKey:RecipeID"`
	Ingredients    []Ingredient   `gorm:"foreignKey:RecipeID"`
	PrepTime       uint           
	CookTime       uint           
	TotalTime      uint           
	Servings       uint
	DifficultyID   uint           `gorm:"not null"`
	Difficulty     DifficultyType `gorm:"type:text;not null;index"`
	Calories       uint
	Macronutrients []string       `gorm:"type:json"`
	Rating         float32        `gorm:"default:0"`
	ReviewCount    uint           `gorm:"default:0"`
	Author         string         `gorm:"not null"`
	Notes          string         `gorm:"type:text"`
	Equipment      []string       `gorm:"type:json"`
	DietaryInfo    []string       `gorm:"type:json"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

type Ingredient struct {
	ID        uint   `gorm:"primaryKey"`
	RecipeID  uint   `gorm:"not null"`
	Name      string `gorm:"type:text;not null"`
}

type Instruction struct {
	ID           uint   `gorm:"primaryKey"`
	RecipeID     uint   `gorm:"not null"`
	StepNumber   uint   `gorm:"not null"`
	Description  string `gorm:"type text;not null"`
}

type Category struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"not null"`
	Recipes []Recipe `gorm:"foriegnKey:CategoryID"`
}

type Difficulty struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Recipes []Recipe `gorm:"foriegnKey:DifficultyID"`
}