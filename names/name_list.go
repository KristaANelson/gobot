package names

import (
	"math/rand"
	"strings"
)

var firstNames = []string{
	"Launa",
	"Tobias",
	"Helga",
	"Lonna",
	"Arvilla",
	"Ricky",
	"Tanner",
	"Valentina",
	"Sherley",
	"Shira",
	"Joye",
	"Mozelle",
	"Floria",
	"Hosea",
	"Boyd",
	"Robert",
	"Robbie",
	"Isela",
	"Rana",
	"Waylon",
	"Dan",
	"Cecilia",
	"Nu",
	"Karoline",
	"Delilah",
	"Nicki",
	"Lani",
	"Annabel",
	"Milissa",
	"Marjory",
	"Lilli",
	"Queenie",
	"Terence",
	"Yukiko",
	"Holley",
	"Genoveva",
	"Normand",
	"Juli",
	"Eveline",
	"Elvia",
	"Raisa",
	"Jess",
	"Clair",
	"Fabian",
	"Lino",
	"Maryellen",
	"Vernell",
	"Wendy",
	"Karena",
	"Joanie",
	"Lanell",
	"Mathilda",
	"Sanjuana",
	"Oma",
	"Rhiannon",
	"Jack",
	"Gabriele",
	"Orval",
	"Lynnette",
	"Felipe",
	"Rayna",
	"Leonor",
	"Soo",
	"Melodie",
	"Vivien",
	"Elli",
	"Jenee",
	"Rufus",
	"Meghann",
	"Julia",
	"Jane",
	"Tonisha",
	"Kenyatta",
	"Armandina",
	"Ardelia",
	"Hobert",
	"Exie",
	"Todd",
	"Lashaun",
	"Susana",
	"Krystina",
	"Cleta",
	"Esta",
	"Emely",
	"Vickie",
	"Loma",
	"Tegan",
	"Charlette",
	"Amparo",
	"Flor",
	"Gala",
	"Gudrun",
	"Candice",
	"Juana",
	"Lucille",
	"Amie",
	"Christen",
	"Albertina",
	"Frederica",
	"Claris",
	"Masako",
	"Blondell",
	"Fausto",
	"Paul",
	"Tess",
	"Brenton",
	"Carolee",
	"Nena",
	"Grover",
	"Loren",
	"Errol",
	"Delmy",
	"Piper",
	"Alma",
	"Rocky",
	"Lisha",
	"Loriann",
	"Shay",
	"Dillon",
	"Hillary",
	"Jule",
	"Emil",
	"Vinnie",
	"Numbers",
	"Berna",
	"Candis",
	"Cyndi",
	"Jonah",
	"Sarina",
	"Kayce",
	"Maud",
	"Annice",
	"Gary",
	"Lan",
	"Bethanie",
	"Dane",
	"Keturah",
	"Judith",
	"Maxie",
	"Leisha",
	"Na",
	"George",
	"Mandi",
	"Lacresha",
	"Corie",
	"Cletus",
	"Shanita",
	"Neoma",
	"Aura",
	"Wilhemina",
}
var lastNames = []string{
	"Marc",
	"Morgado",
	"Sine",
	"Oberg",
	"Said",
	"Nicols",
	"Asay",
	"Samuelson",
	"Lusby",
	"Dudash",
	"Even",
	"Say",
	"Bothwell",
	"Barkley",
	"Prall",
	"Paulino",
	"Loughman",
	"Birkholz",
	"Dunnigan",
	"Hayek",
	"Hardiman",
	"Porch",
	"Pavlak",
	"Sakai",
	"Vinci",
	"Ernst",
	"Homan",
	"Yates",
	"Halbrook",
	"Besaw",
	"Serpa",
	"Goncalves",
	"Kraatz",
	"Laycock",
	"Zentz",
	"Strelow",
	"Kinyon",
	"Dudney",
	"Sabat",
	"Boden",
	"Knouse",
	"Pinkard",
	"Mikula",
	"Chapdelaine",
	"Wayt",
	"Cryer",
	"Dalbey",
	"Crespo",
	"Lever",
	"Hundley",
	"Janzen",
	"Homes",
	"Eberhard",
	"Frisch",
	"Wiest",
	"Pichette",
	"Dement",
	"Tomberlin",
	"Mallari",
	"Stumpff",
	"Murry",
	"Frierson",
	"Dimauro",
	"Cashwell",
	"Quezada",
	"Miracle",
	"Asmus",
	"Sheffler",
	"Kehrer",
	"Hensler",
	"Coward",
	"Owsley",
	"Vankeuren",
	"Zuckerman",
	"Helbert",
	"Emmer",
	"Sedberry",
	"Beamer",
	"Grand",
	"Landy",
	"Gonsoulin",
	"Costantino",
	"Mcnaught",
	"Braddock",
	"Lacourse",
	"Willner",
	"Thresher",
	"Dumas",
	"Sommers",
	"Bogart",
	"Olivarria",
	"Kaiser",
	"Topping",
	"Creasy",
	"Heuser",
	"Studer",
	"Mcavoy",
	"Walcott",
	"Varian",
	"Berta",
	"Leight",
	"Swain",
	"Rapier",
	"Earhart",
	"Hohler",
	"Rumsey",
	"Rubino",
	"Dillman",
	"Kucharski",
	"Dodge",
	"Cappello",
	"Sidwell",
	"Frew",
	"Hopf",
	"Costner",
	"Gravitt",
	"Bodnar",
	"Duffel",
	"Sachs",
	"Jeon",
	"Bostic",
	"Rega",
	"Cuomo",
	"Tovey",
	"Mcgowan",
	"Cheatam",
	"Jeppesen",
	"Pennell",
	"Heppner",
	"Hodgin",
	"Bernal",
	"Gallman",
	"Guss",
	"Seidel",
	"Montesano",
	"Weed",
	"Burkley",
	"Strohl",
	"Nickle",
	"Bohland",
	"Ivey",
	"Fogal",
	"Powel",
	"Fessler",
	"Friddle",
	"Huseby",
	"Keep",
	"Pickard",
	"Krach",
	"Mclennon",
}

func randomName() string {
	firstName := firstNames[rand.Int()%len(firstNames)]
	lastName := lastNames[rand.Int()%len(lastNames)]
	return strings.Join([]string{firstName, lastName}, " ")
}
