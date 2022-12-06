package utils

import (
	"math/rand"
	"time"
)

var colours = [...]string{
	"\033[30m",
	"\033[31m",
	"\033[32m",
	"\033[33m",
	"\033[34m",
	"\033[35m",
	"\033[36m",
	"\033[37m",
}

func GetRandomColour() string {

	seed := rand.NewSource((time.Now().UnixNano()))
	random := rand.New(seed)
	return colours[random.Intn(len(colours))]
}
