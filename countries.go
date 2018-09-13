package app

var supportedCountries = []string{
	"afghanistan", "albania", "algeria", "andorra", "angola",
	"antigua and barbuda", "argentina", "armenia", "aruba", "australia", "austria",
	"azerbaijan", "bahamas", "bahrain", "bangladesh", "barbados", "belarus", "belgium", "belize",
	"benin", "bhutan", "bolivia", "bosnia and herzegowina", "botswana",
	"brazil", "british indian ocean territory", "bulgaria", "burkina faso",
	"burundi", "cambodia", "cameroon", "canada", "cape verde",
	"central african republic", "chad", "chile", "china",
	"colombia", "comoros", "congo",
	"costa rica", "croatia", "cuba", "cyprus",
	"czech", "denmark", "djibouti", "dominica", "dominican republic", "east timor",
	"ecuador", "egypt", "el salvador", "equatorial guinea", "eritrea", "estonia", "ethiopia",
	"fiji", "finland", "france",
	"gabon", "gambia", "georgia", "germany", "ghana", "greece", "greenland",
	"grenada", "guatemala", "guinea", "guinea-bissau", "guyana", "haiti",
	"honduras", "hungary",
	"iceland", "india", "indonesia", "iran", "iraq", "ireland", "israel",
	"italy", "jamaica", "japan", "jordan", "kazakhstan", "kenya", "kiribati",
	"north korea", "south korea", "kuwait", "kyrgyzstan",
	"laos", "latvia", "lebanon", "lesotho", "liberia",
	"libya", "liechtenstein", "lithuania", "luxembourg",
	"macedonia", "madagascar", "malawi", "malaysia", "maldives",
	"mali", "malta", "mauritania", "mauritius",
	"mexico", "micronesia", "moldova", "monaco", "mongolia",
	"morocco", "mozambique", "myanmar", "namibia", "nauru", "nepal", "netherlands",
	"netherlands antilles", "new zealand", "nicaragua", "niger", "nigeria", "niue",
	"norway", "oman", "pakistan", "palau", "panama",
	"papua new guinea", "paraguay", "peru", "philippines", "poland", "portugal",
	"puerto rico", "qatar", "romania", "russian federation", "rwanda",
	"saint lucia", "saint vincent and grenadines",
	"samoa", "san marino", "sao tome and principe", "saudi arabia", "senegal",
	"seychelles", "sierra leone", "singapore", "slovakia",
	"slovenia", "solomon islands", "somalia", "south africa",
	"spain", "sri lanka", "sudan", "suriname", "swaziland",
	"sweden", "switzerland", "syria", "taiwan", "tajikistan",
	"tanzania", "thailand", "togo", "tonga", "trinidad and tobago",
	"tunisia", "turkmenistan", "tuvalu", "uganda", "ukraine",
	"united arab emirates", "united kingdom", "united states of america",
	"uruguay", "uzbekistan", "vanuatu", "vatican city", "venezuela", "vietnam", "wales", "yemen",
	"zambia", "zimbabwe",
}

// Dialogflow returns formalized country name, but file name is different
var replaceMap = map[string]string{
	"argentina":                "Argentina - Long",
	"united states of america": "U.S.A",
	"russian federation":       "Russia",
	"bosnia and herzegowina":   "Bosnia-Herzegovina",
	"brazil":                   "Brazil - Long",
	"congo":                    "Congo, Republic of",
	"peru":                     "Peru - Long",
	"papua new guinea":         "Papau New Guinea",
	"botswana":                 "Bostwana",
	"guinea-bissau":            "Guinea-Bissau",
	"azerbaijan":               "Azerbaidjan",
}
