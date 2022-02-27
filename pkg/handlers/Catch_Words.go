package handlers

import (
	"math/rand"
	"time"
)

var things []string = []string{
	"George Washington",
	"Abraham Lincon",
	"Nestle",
	"Moira Rose",
	"Lupin",
	"Potter",
	"Holly",
	"Fonzi",
	"Roxy",
	"Taz",
	"Zeke",
	"Missy Franklin",
	"Roger Federer",
	"Katie Ledecky",
	"Michael Phelps",
	"Bjorn Borg",
	"Philip J. Fry",
	"Rafael Nadal",
	"The Professor",
	"Buzz Lightyear",
	"Tom Hanks",
	"Dwight Schrute",
	"Jim Halpert",
	"McDonalds",
	"Burger King",
	"Wendys",
	"Jimmy John's",
	"Zeus",
	"Hades",
	"Hercules",
	"Bo Jackson",
	"Joe Rogan",
	"Tom Cruise",
	"Coffee",
	"Tea",
	"Shaq",
	"Michael Jordan",
	"Beyonce",
	"Slim Shady",
	"Dr. Dre",
	"Snoop Dogg",
	"Subway",
	"Jersey Mike's",
	"Scientology",
	"John Travolta",
	"Amazon",
	"Google",
	"Apple",
	"Andriod",
}

func Get_Next_Word(lastWord string) string {

	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator

	nextWord := things[rand.Intn(len(things))]
	for nextWord == lastWord {
		nextWord = things[rand.Intn(len(things))]
	}
	//fmt.Println("lastWord : " + lastWord)
	//fmt.Println("nextWord : " + nextWord)

	return nextWord
}
