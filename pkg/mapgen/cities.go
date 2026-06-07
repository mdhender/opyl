// Copyright (c) 2026 Michael D Henderson. All rights reserved.

package mapgen

import "github.com/mdhender/opyl/internal/infra/prng"

type CityNameGenerator struct {
	list []string
}

// NewCityNameGenerator uses the rng to return a generator seeded with a
// random list of names. Call cng.Name to obtain a randomly selected name.
func NewCityNameGenerator(rng *prng.PRNG) *CityNameGenerator {
	cng := &CityNameGenerator{list: append([]string{},
		"Abreway", "Aburg", "Accolon", "Aclomark", "Adalgar", "Adee", "Adelaide", "Adric", "Aelbarik",
		"Aeltra", "Aelwynne", "Aerilon", "Aesa", "Aethelarn", "Aeweald", "Agate City", "Agnar", "Aindorf",
		"Akyket", "Alaris", "Albrand", "Ald Ruhn", "Ald Sotha", "Aldain", "Aldcliff", "Aldmaple", "Aldmead",
		"Alera", "Alicor City", "Alldus", "Alsvider", "Alwood", "Amatin", "Anarch", "Anarth", "Andrasreth",
		"Anghurst", "Anniun", "Antbin", "Anyo", "Aodhagan", "Apore", "Appleville", "Aquarin", "Aral",
		"Aramoor", "Areth Pirn", "Argon City", "Arkaune", "Arkfire City", "Arlington", "Asbo", "Ashalmawia", "Ashland",
		"Assarinibib", "Astryde", "Atdon", "Aten", "Atlantis", "Attewelle", "Auburn", "Aumenburg", "Aure",
		"Aurion", "Axeth", "Axheim", "Ayastowe", "Ayham", "Azland", "Azmar", "Azure City", "Babrook",
		"Baelraeth", "Baeron", "Balatea", "Balderon", "Balmora", "Bamfirth", "Bampu", "Banwes", "Baraxes",
		"Barebranch", "Barghan", "Barren City", "Barrowwald", "Baseltded", "Bawside", "Bayarth", "Bayfield", "Bayrun",
		"Beacon Hill", "Befield", "Beggar’s Hole", "Belcoast", "Belgarth", "Beltran", "Belwe", "Benfield", "Benoic",
		"Benpretbrook", "Beorttun", "Berandes", "Bernthe", "Bestead", "Betport", "Bicanton", "Binacre", "Binshor",
		"Birdan", "Black Hill", "Black Hollow", "Blacksoul", "Blaglen", "Ble'eck", "Blerglust", "Blue Field", "Bluebeach",
		"Blukirk", "Blyn", "Boatwright", "Boden City", "Boheim", "Bonehorn", "Boryn", "Braddach", "Bragge",
		"Branora", "Branraker", "Branwen", "Brapslung", "Briar Glen", "Brickelwhyte", "Brine City", "Brisbane", "Bristol",
		"Briwater", "Broken Shield", "Brookmoor", "Buckmoth", "Buelthane", "Bul Isra", "Bullmar", "Buren", "Burlington",
		"Burnside", "Butterice", "Caase", "Cabury", "Caden", "Caer City", "Cafan", "Caglex", "Calanthe",
		"Caldera", "Caliron", "Calldwr", "Canberra", "Cano", "Caracatus", "Carim", "Carmisa", "Carran",
		"Carru", "Carslenford", "Cartseth", "Casthe", "Castle Hill", "Catchclaw", "Cawode", "Cear", "Cedon",
		"Cedway", "Celestra", "Cellangham", "Celoydorf", "Cengrove", "Centerville", "Cerbruk", "Cerbury", "Ceremon",
		"Cerglen", "Cerkas", "Cerrin City", "Cesirun", "Chagrad", "Chater", "Che'aldgost", "Chebas", "Chedytown",
		"Cherfren", "Chiron", "Chronos", "Cigry", "Cilburn", "Cile City", "Cinbach", "Cirrele", "Ciryon",
		"City of Fire", "City of the Shadows", "Claunecar", "Clay City", "Clayton", "Cleanbones", "Clearham Downs", "Cleveland", "Clinstan",
		"Clinton", "Clyf", "Clywd", "Coalfell", "Coel", "Coinbalth", "Coldrose", "Coldshadow", "Coldshore",
		"Condmedic", "Congisfirth", "Cormyr", "Cospera", "Cotys", "Courtmarsh", "Craeven", "Cragghe", "Crasspest",
		"Crimir", "Crystalshadow", "Cullfield", "Culshire", "Custyn", "Cyrdfel", "Cyrkarth", "Dacborath", "Dacria",
		"Daerte", "Dafarik", "Dafyd", "Dagonfel", "Dagr", "Dagrove", "Daiwick", "Dalamar", "Damarck",
		"Danfen", "Dantor", "Daon", "Darcon", "Darkbone City", "Darkmage", "Darkwell", "Darkwind", "Darwin",
		"Dasbach", "Dathoth", "Davchar", "Dayton", "Deathfall", "Dedonburn", "Dedpool", "Deepcrest", "Deephall Point",
		"Deepmoor", "Deirburn", "Delde City", "Delgrove", "Dellgate", "Delorn", "Delphys", "Denaste", "Dencede",
		"Dendest", "Dengelfel", "Deorward", "Dergost", "Derokin", "Descarq", "Detheim", "Deybank", "Dezarne",
		"Diaratyh", "Diarmaoid", "Dindale", "Dirbrand", "Disprelfield", "Distran", "Diuran", "Dogrock", "Dolgan",
		"Domin", "Doonatel", "Dorenth", "Dorhaven", "Dostborough", "Dover", "Doxca", "Doycro", "Draeden",
		"Dragonmarsh", "Drassa", "Drinishok", "Dry Gulch", "Dryope", "Duddaleah", "Dungon", "Durshire", "Dwalin",
		"Dyot", "Eadak", "Eadgyth", "Eadweard", "Earnberict", "Easthaven", "Eathelin", "Ebonheart", "Ebow City",
		"Echule", "Ecrin", "Eddra", "Edgegate", "Edorin", "Ekkel", "Eknanbor", "Elakain", "Elantir",
		"Elbramair", "Eldead", "Elderon", "Elysson", "Emerald City", "Emulpool", "Engion", "Enheim", "Enless",
		"Eoghan", "Eorforwic", "Erabenimsun", "Erast", "Erbham", "Eregdor", "Eribank", "Eron", "Esdros",
		"Esme's Rot", "Esnar", "Essault", "Estbeorn", "Estercoast", "Eststead", "Ethna", "Etnenk", "Etranth",
		"Evinob", "Faelgrar", "Faerwald", "Fairburn Point", "Fairmeadow", "Fairview", "Falconlake", "Falensarano", "Fallash Bridge",
		"Fallhedge", "Fallville", "Falo'a", "Falrepent", "Fanborough", "Fandrall", "Far Water", "Fargate", "Farnor",
		"Farshadow", "Fassen", "Faxbury", "Faycastle", "Fayfair", "Fearn", "Fearshadow", "Fedresheath", "Fegwern",
		"Feno", "Filugrave", "Firebend", "Fistrock", "Flairown", "Flat City", "Flelheim", "Fletacre", "Fondcot",
		"Fool's March", "Fornil", "Forterk", "Franklin", "Freywall", "Frostford", "Fyxinca", "Gaelen", "Gaethaa",
		"Gallys", "Gandar", "Gane", "Ganelon", "Ganith", "Garen's Well", "Garlupool", "Garn City", "Gaytforth",
		"Gealkend", "Geatan", "Geirrod", "Gentlewind", "Georgetown", "Gildeath", "Ginisis", "Gisapool", "Glasscliff",
		"Glastowe", "Glatchan", "Glingedheath", "Gnarr", "Gocin", "Godehard", "Gofannon", "Goldcrest", "Goldenleaf",
		"Goldlyn", "Goodan", "Gostarbach", "Gowerd", "Gralnen", "Gravecarn", "Greendell", "Greenhill", "Greenville",
		"Greton", "Greymage", "Greymarsh", "Greywater Edge", "Grimwall", "Grindor", "Groltain", "Gutar", "Gwayhne",
		"Haakon", "Habyn", "Hacranbrook", "Haele", "Haemfrith", "Haemin", "Haertlinde", "Hahel", "Haim",
		"Halfolk", "Halhere", "Hamish", "Haran", "Harmakros", "Harn", "Harshire", "Hasfolk", "Havale",
		"Hawkwind", "Hebost", "Hermrord", "Hestan", "Hifro", "Highdale", "Highmeadow", "Hildieth", "Hillfar",
		"Hlormaren", "Hobart", "Hogsfeet", "Holamayn", "Holith", "Hollian", "Hollyhead", "Holthasburg", "Hoochillwick",
		"Hornmar", "Hostyyk", "Hrodowald", "Hudson", "Huffimstowe", "Hull", "Hultor", "Hwen", "Hykirk",
		"Hyksos", "Iaxil", "Icebarrow", "Icefay", "Icemeet", "Iche", "Ickasu", "Illa", "Illinod",
		"Ilya", "Iniera", "Innsmouth", "Inos", "Inte", "Iprial", "Ironforge", "Ironplow", "Ironville Crossing",
		"Irragin", "Ixdencer", "Jackson", "Jancastle", "Janlyn", "Jarren's Outpost", "Jasand", "Jedarhe", "Jeling",
		"Jhena", "Jina", "Jongvale", "Jontmac", "Ka'oma", "Kara's Vale", "Kerreck", "Khartag", "Khuul",
		"Kingston", "Kior", "Kipamod", "Knife's Edge", "Koal", "Kohgoruhn", "Kouglen", "Krallides", "Kuneack",
		"Lakeshore", "Landpond", "Las Vegas", "Lassan", "Lawpest", "Laxton", "Le'oght", "Lebrus", "Leehaven",
		"Leeside", "Lenham", "Lethrys", "Lexington", "Linism", "Linland", "Linmeadow", "Lintown", "Litysh",
		"Liwald", "Lochfort", "Lochhurst", "Lookfar", "Lorbeach", "Lorbin", "Lullin", "Lyborough", "Madison",
		"Mageland", "Maire", "Mallon", "Mallowbrook", "Manchester", "Manmint", "Mantooth", "Maplehurst", "Marath",
		"Marblemoor", "Marbleton", "Margate", "Marion", "Marr Gan", "Marren's Eve", "Masellil", "Meadowlake", "Meaple",
		"Melbourne", "Merribourne", "Meshburn", "Meton", "Metropolis", "Metwan", "Mide", "Mikum", "Milford",
		"Millitburn", "Millstone", "Milton", "Mimea", "Minhanstowe", "Miranth", "Mishgrave", "Misiport", "Mitu",
		"Mompi", "Moonbright", "Moonfire", "Moonmoth", "Mora", "Morcrest", "Mount Pleasant", "Mount Vernon", "Mountmend",
		"Movawood", "Mowbach", "Muqueling", "Murplant", "Mutzcat", "Myhra", "Mysa", "Mytchville", "Nadin",
		"Nantasarn", "Narlenrun", "Nassic", "Nearon", "Nepill", "Nergwern", "Nespho", "Nestan", "Netgrove",
		"Neuson", "New Cresthill", "Newbald", "Newhaw", "Newleaf", "Newport", "Newton", "Newtown",
		"Nightfrost", "Nophalis", "Norbank", "Norbus", "Norgar", "Norratyn", "Northhollow", "Northmold", "Northpass",
		"Notlbrob", "Nuchuleft", "Nuncarth", "Nuxvar", "Nyssa", "Oakborough", "Oakheart", "Oakland", "Oaldar",
		"Oar's Rest", "Obraed", "Ociera", "Ocshire", "Odana", "Odar", "Odasgunn", "Odrosal", "Odwulf",
		"Oftar", "Old Ashton", "Oldcastle", "Oldel", "Oldshade", "Ollaneg", "Ollayos", "Olmar",
		"Olon", "Oltpest", "Omournil", "Onbruk", "Orbost", "Orfler", "Orianna", "Orness", "Orre",
		"Orrinshire", "Orwald", "Osgea", "Ossa", "Othkar", "Othon", "Oxford", "Ozorak", "Ozryn",
		"Pactra", "Pallia", "Panplara", "Pantarastar", "Pavv", "Peadar", "Peash", "Pella's Wish", "Pelra",
		"Pelthros", "Penbarn", "Penci", "Penrili", "Pentara", "Pentgaland", "Perchhead", "Perendor", "Pereswyff",
		"Perth", "Pesteir", "Petelinus", "Phames", "Piaside", "Pictar", "Pildor", "Pinnella Pass", "Pirn",
		"Plagcath", "Plemarun", "Plinsaway", "Po'asta", "Poltgobi", "Pomlinfolk", "Pontent", "Pothbaz",
		"Praice", "Pran", "Presrenfa", "Prothla", "Proudrock", "Proupol", "Prymarsh", "Prytani", "Prywyn",
		"Pultack", "Pygate", "Quagcry", "Quan Ma", "Queenstown", "Quickrock", "Raigor", "Rairkvale", "Ralsinpe",
		"Ramshorn", "Raskold", "Rathisa", "Ravenbow", "Raypond", "Rayth", "Reaver", "Red Hawk", "Redcliff",
		"Redwine", "Remdam", "Rertstet", "Resh", "Rhifirth", "Rhonius", "Riabury", "Rimmon", "Rindalsem",
		"Riveredge", "Riverside", "Riverton", "Riverwind", "Rolkfield", "Rorcy", "Roseglass", "Rosewall", "Roshun",
		"Rotshaw", "Ruanrath", "Runesward", "Rushownstad", "Rustan", "Rylla", "Ryshet", "Rysshop", "Rytor",
		"Ryvwy", "Saker Keep", "Salach", "Salamus", "Salem", "Sallen", "Saltstone", "Samca", "Sandrith",
		"Sardis", "Sarindor", "Sartheim", "Sasbank", "Scadee", "Schaldhan", "Scrushfield", "Seafort", "Seamarsh",
		"Seameet", "Secordrus", "Sedeor", "Semfirth", "Sentrun", "Serra", "Sesklos", "Sessgate", "Sestun",
		"Seyda", "Shadowdale", "Shadowmoor Downs", "Shadowpond Point", "Shavnor", "Shencodo", "Shermer", "Ship's Haven", "Shotluth",
		"Shrecrun", "Shull", "Silver City", "Silverkeep", "Silverpond Crags", "Sinyanwood", "Skyllith", "Slamer", "Sloosio",
		"Smant", "Smarmark", "Snake's Canyon", "Snowland", "Snowmelt", "Sompishaw", "Sontsil", "Sorale", "Sout",
		"South Warren", "Spedgost", "Splexnaiss", "Splustap", "Springfield", "Squall's End", "Squangrir", "Squikgost", "Stannecpan",
		"Starrycastle", "Stattown", "Steherwood", "Stenral", "Sti'yl", "Stickgate", "Stilom", "Stoneby", "Stonehand",
		"Stonle", "Stonstoke", "Storre", "Strongby", "Sudabuk", "Summerby", "Summergrass", "Summermarsh", "Sunwater",
		"Suran", "Swagrave", "Sweetwood", "Swinebroth", "Swordbreak", "Swynfield", "Sydney", "Syssale", "Tallyn",
		"Tanglen", "Taniholm", "Tarrin", "Tarvik", "Tastan", "Techanal", "Teffolk", "Telasero", "Telkna",
		"Tendaughters", "Tensa", "Tenwood", "Teran", "Tese", "Tetaun", "Thaholm", "Theeltmil", "Thorbed",
		"Thorneclay", "Thornfield", "Three Streams", "Throll", "Thyria", "Tickben", "Tijinggrad", "Tildor", "Tirne",
		"Tirport", "Tonkcent", "Toppe", "Tosdis", "Tosla", "Tramrare", "Trawald", "Trekrun", "Trion",
		"Trotwood", "Trucdon", "Trudid", "Trullion", "Tureynulal", "Twilight City", "Twisernvale", "Tywy", "Ubbin Falls",
		"Uinan", "Ula'ree", "Ulgor", "Ullast", "Unhink", "Urke", "Urshilaku", "Urvil", "Ushoul",
		"Usibel", "Ussun", "Uvirith", "Valacre", "Valenvelvar", "Valwick", "Vamstead", "Veamer", "Velishire",
		"Velothi", "Vemynal", "Veritas", "Vernolt", "Vertloch Bridge", "Vertmount Downs", "Viberg", "Vikos", "Violl's Garden",
		"Vister", "Vivec", "Vo'irnil", "Volmoria", "Vorasen", "Vril", "Wacot", "Waleoshire", "Wann",
		"Washington", "Watercoast", "Wavemeet", "Wavenhill", "Wellspring", "Wemau", "Wensol", "West Ford", "Westbay",
		"Westen", "Westertown", "Westervale", "Wetrock", "Wheatland", "Whedorf", "Whitehollow", "Whitepine", "Whiteridge",
		"Whitewell Land", "Wildebush", "Wildefort", "Wilea", "Willowdale", "Winchester", "Windrip", "Windwhisper", "Winewood",
		"Winpher", "Winterfeather", "Winterness", "Wintervale", "Wissgate", "Wistan", "Wistleigh", "Witchlyn", "Wolfden",
		"Wolfhair", "Wolfkeed", "Woodbush", "Woodend", "Woodside", "Wurcot", "Wyndu", "Xan's Bequest", "Xynnar",
		"Yarrin", "Ys", "Zaal", "Zaina", "Zao Ying", "Zapstowe", "Zeffari", "Zisbach", "Zumka",
		"Zurgonipal",
	)}
	rng.Shuffle(len(cng.list), func(i, j int) {
		cng.list[i], cng.list[j] = cng.list[j], cng.list[i]
	})
	return cng
}

// Name returns the name popped from the end of the list of names
func (cng *CityNameGenerator) Name() string {
	if cng == nil || len(cng.list) == 0 {
		panic("out of names")
	}
	name := cng.list[len(cng.list)-1]
	cng.list = cng.list[:len(cng.list)-1]
	return name
}
