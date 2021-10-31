package config

//THIS SECTION OF CODE HAS NOT BEEN IMPLEMENTED YET

type Person struct {
	Name            string
	Gender          string
	Race            string
	Class           string
	Alignment       string
	Flaws           []string
	Personality     string
	Ideals          []string
	Bonds           []string
	KnownLang       []string
	Weapons         []string
	Equipment       []string
	Mannerisms      []string
	Occupation      string
	Secrets         []string
	UsefulKnowledge []string
	Story           string
}

type Faction struct {
	Name            string
	Goals           []string
	Methods         []string
	Members         []Person
	Secrets         []string
	UsefulKnowledge []string
}

type Location struct {
	Name       string
	Location   string
	History    string
	Size       string
	Government string
	Ruler      Person
	Alliances  []string
	Enemies    []string
	Currency   string
	Problems   []string
	POI        []Person
}

type Kingdom struct {
	Name               string
	Capital            Location
	Ruler              Person
	Government         string
	Age                int
	OfficialLanguages  []string
	SecondaryLanguages []string
	History            string
	Size               string
	PlacesOfInterest   []Location
}

type Continent struct {
	Name     string
	Kingdoms []Kingdom
	Histroy  string
}

type Items struct {
	Name      string
	Age       int
	Abilities []string
	Curses    []string
	Location  Location
	Story     string
}

type God struct {
	God        Person
	Histroy    string
	Worshipers []Person
}

type World struct {
	Name       string
	History    string
	Continents []Continent
	ExtraLaws  []string
}
