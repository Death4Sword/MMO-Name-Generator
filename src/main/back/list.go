package back

var dwarfNameArr = [180]string{
	"Dronglik",
	"Dwalin",
	" Damin",
	" Drong",
	"Drongli",
	" Dwalik",
	" Duregar",
	"Dern",
	"Durgrim",
	"Durin",
	"Drumin",
	"Dimrond",
	"Dimzad",
	" Daled",
	"Durgin",
	"Drokk",
	"Dertrain",
	" Daled",
	" Drakki",
	"Darbli",
	"Dorin",
	"Dargo",
	"Durmak",
	"Drang",
	" Drong",
	"Durzak",
	"Durak",
	"Durgin",
	"Dwinbar",
	"Dagmar",
	"Dorag",
	"Dormar",
	" Dorn",
	"Durbar",
	"Durifer",
	"Dolgr",
	"Dovvet",
	"Durlin",
	"Danrun",
	"Dodrum",
	"Durlin",
	"Drokki",
	"Darent",
	"Dovvent",
	"Drongvor",
	"Duregar",
	" Dorrin",
	"Durgan",
	"Drugni",
	"Dunden",
	" Dundin",
	" Dagren",
	"Duric",
	"Elgrom",
	"Elgroth",
	"Eldarig",
	"Eisenbjorn",
	"Elgroth",
	"Eldarig",
	"Furagrum",
	"Furtan",
	"Furgil",
	"Fudrin",
	"Finn",
	"Firgil",
	"fenni",
	"Fimbur",
	"Fungg",
	"Fondar",
	"Falco",
	"Fech",
	"Ferh",
	"Finbar",
	"Fijar",
	"Framm",
	"Fraunk",
	"Fumok",
	"Fimarin",
	"Furen",
	"fin",
	"Furakrag",
	"Furd",
	"Fimki",
	"Furmir",
	"Furakrag",
	"Finarin",
	"Furmir",
	"Firengul",
	" Faramin",
	"Farlin",
	"Farnir",
	"Fengast",
	"Fredi",
	"Goddi",
	"Gomrund",
	"Gadrin",
	"Gorazin",
	"Gorim",
	"Gorm",
	"Gottri",
	"Gimnir",
	"Golendhil",
	"Garil",
	" Grumdin",
	" Grimbul",
	"Garaith",
	" Gudrum",
	"Gorrin",
	"Gudii",
	"Gorgi",
	"Groth",
	"Grindol",
	"Grom",
	"Grond",
	"Groth",
	"Grum",
	"Gumli",
	"Grundi",
	"Grung",
	"GuttriGotrek",
	"Grodrik",
	"Gorri",
	"Gadrin",
	"Gumli",
	"Godrik",
	"Gorri",
	"Grimbul",
	"Garagrim",
	"Gorazin",
	"Grimni",
	"Grimbeard",
	"Godir",
	"Grombrindal",
	"Grim",
	"Grimli",
	"Gimlin",
	" Grongi",
	"Grondik ",
	"Grumbar",
	"Grumhilde",
	"Grom",
	"Gemmund",
	"Geirrod",
	"Gerar",
	"Gloier",
	"Gofnyr",
	"Gorwinay",
	"Gothrom",
	"Grennel",
	"Gruff",
	"Grimsiad",
	"Grummore",
	"Grummer",
	"Guirod",
	"Gulnyr",
	"Guthorm",
	"Gymir",
	" Grimgor",
	" Gantor",
	"Golgin",
	"Grundi",
	"Gronfin",
	"Grimgor",
	" Grumir",
	"Grendel",
	"Grumdnir",
	"Gargrim",
	" Gunnolf",
	"Gantor",
	" Gav",
	"Grogrim",
	"Grodun",
	" Glod",
	"Grulfdok",
	"Gavrii",
	"Grogrund",
	" Grodi",
	" Gobrik",
	"Grunak",
	"Galin",
}

var elfNameArr = [91]string{
	"Aradehel",
	"Arwen",
	"Belwen",
	"Balwena",
	"Curufinwë",
	"Celenear",
	"Danica",
	"Daenara",
	"Elbereth",
	"Estë",
	"Fatris",
	"Fëanturi",
	"Heldaria",
	"Indis",
	"Itarillë",
	"Kardryar",
	"Laurelin",
	"Lindorie",
	"Luinil",
	"Miriel",
	"Niniel",
	"Nandil",
	"Oilossë",
	"Qorwïn",
	"Rian",
	"Serindë",
	"Siltiame",
	"Trisphyra",
	"Tinûviel",
	"Than",
	"Urilana",
	"Valstina",
	"Valgella",
	"Varda",
	"Virani",
	"Wilwarin",
	"Yavana",
	"Yndreth",
	"Anardil",
	"Atanatar",
	"Aldaron",
	"Argwaen",
	"Adanedhel",
	"Amlach",
	"Boromir",
	"Bronweg",
	"Beiven",
	"Beivenor",
	"Ballar",
	"Calion",
	"Cyrion",
	"Celeborn",
	"Denethor",
	"Dior",
	"Daurin",
	"Elros",
	"Elronde",
	"Finarlin",
	"Faramir",
	"Gondolin",
	"Gurthang",
	"Haldir",
	"Fantur",
	"Isildur",
	"Handir",
	"Kahe",
	"Kadath",
	"Lath",
	"Legolas",
	"Lomion",
	"Maeglin",
	"Maglin",
	"Nenuial",
	"Nartshseg",
	"Oriour",
	"Orodruin",
	"Peredhel",
	"Panthael",
	"Rana",
	"Saruman",
	"Saro",
	"Thalion",
	"Urith",
	"Uilos",
	"Valwë",
	"Wilwarin",
	"Yenlu",
	"Yelhorn",
	"Yinven",
	"Zinwraek",
	"Zumcas",
}
