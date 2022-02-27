package handlers

import (
	"math/rand"
	"time"
)

func Get_Next_Word() string {
	things := make([]string, 0)
	things = append(things,
		"George Washington",
		"Abraham Lincon",
		"Nestle",
		"Moira Rose")

	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	return things[rand.Intn(len(things))]
}
