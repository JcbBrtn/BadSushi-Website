package handlers

import (
	"fmt"
	"math/rand"
	"time"
)

func Get_Next_Word(lastWord string) string {
	things := make([]string, 0)
	things = append(things,
		"George Washington",
		"Abraham Lincon",
		"Nestle",
		"Moira Rose",
		"Bailey Grimm",
	)

	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator

	nextWord := things[rand.Intn(len(things))]
	for nextWord == lastWord {
		nextWord = things[rand.Intn(len(things))]
	}
	fmt.Println("lastWord : " + lastWord)
	fmt.Println("nextWord : " + nextWord)

	return nextWord
}
