package main

type DifficultyType string

const (
	Easy   DifficultyType = "Easy"
	Medium DifficultyType = "Medium"
	Hard   DifficultyType = "Hard"
)

var Difficulties = []DifficultyType{Easy, Medium, Hard}

type CategoryType string

const (
	Breakfast  CategoryType = "Breakfast"
	Lunch      CategoryType = "Lunch"
	Dinner     CategoryType = "Dinner"
	Dessert    CategoryType = "Dessert"
	Snack      CategoryType = "Snack"
	Beverage   CategoryType = "Beverage"
)

var Categories = []CategoryType{Breakfast, Lunch, Dinner, Dessert, Snack, Beverage}