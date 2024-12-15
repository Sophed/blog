package main

import "math/rand"

var (
	ADJECTIVES = []string{
		"elegant", "vibrant", "mysterious", "serene",
		"graceful", "ancient", "bold", "delicate",
		"whimsical", "fiery", "charming", "rugged",
		"playful", "luminous", "majestic", "radiant",
		"melodic", "dynamic", "peaceful", "frosted",
	}
	ANIMALS = []string{
		"elephant", "giraffe", "dolphin", "panda",
		"tiger", "kangaroo", "penguin", "lion",
		"zebra", "owl", "octopus", "koala",
		"shark", "wolf", "rabbit", "falcon",
		"crocodile", "peacock", "bison", "butterfly",
	}
)

func randomAdjective() string {
	return ADJECTIVES[rand.Intn(len(ADJECTIVES))]
}

func randomAnimal() string {
	return ANIMALS[rand.Intn(len(ANIMALS))]
}
