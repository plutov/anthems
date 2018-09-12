package app

import "unicode"

var supportedCountries = []string{
	"afghanistan", "albania", "algeria", "andorra", "angola", "anguilla",
	"antarctica", "antigua and barbuda", "argentina", "armenia", "aruba", "australia", "austria",
	"azerbaijan", "bahamas", "bahrain", "bangladesh", "barbados", "belarus", "belgium", "belize",
	"benin", "bermuda", "bhutan", "bolivia", "bosnia and herzegowina", "botswana", "bouvet island",
	"brazil", "british indian ocean territory", "brunei darussalam", "bulgaria", "burkina faso",
	"burundi", "cambodia", "cameroon", "canada", "cape verde", "cayman islands",
	"central african republic", "chad", "chile", "china", "christmas island",
	"cocos (keeling) islands", "colombia", "comoros", "congo", "congo, the democratic republic of the",
	"cook islands", "costa rica", "cote d'ivoire", "croatia (hrvatska)", "cuba", "cyprus",
	"czech republic", "denmark", "djibouti", "dominica", "dominican republic", "east timor",
	"ecuador", "egypt", "el salvador", "equatorial guinea", "eritrea", "estonia", "ethiopia",
	"falkland islands (malvinas)", "faroe islands", "fiji", "finland", "france",
	"france metropolitan", "french guiana", "french polynesia", "french southern territories",
	"gabon", "gambia", "georgia", "germany", "ghana", "gibraltar", "greece", "greenland",
	"grenada", "guadeloupe", "guam", "guatemala", "guinea", "guinea-bissau", "guyana", "haiti",
	"heard and mc donald islands", "holy see (vatican city state)", "honduras", "hong kong", "hungary",
	"iceland", "india", "indonesia", "iran (islamic republic of)", "iraq", "ireland", "israel",
	"italy", "jamaica", "japan", "jordan", "kazakhstan", "kenya", "kiribati",
	"korea, democratic people's republic of", "korea, republic of", "kuwait", "kyrgyzstan",
	"lao, people's democratic republic", "latvia", "lebanon", "lesotho", "liberia",
	"libyan arab jamahiriya", "liechtenstein", "lithuania", "luxembourg", "macau",
	"macedonia, the former yugoslav republic of", "madagascar", "malawi", "malaysia", "maldives",
	"mali", "malta", "marshall islands", "martinique", "mauritania", "mauritius", "mayotte",
	"mexico", "micronesia, federated states of", "moldova, republic of", "monaco", "mongolia",
	"montserrat", "morocco", "mozambique", "myanmar", "namibia", "nauru", "nepal", "netherlands",
	"netherlands antilles", "new caledonia", "new zealand", "nicaragua", "niger", "nigeria", "niue",
	"norfolk island", "northern mariana islands", "norway", "oman", "pakistan", "palau", "panama",
	"papua new guinea", "paraguay", "peru", "philippines", "pitcairn", "poland", "portugal",
	"puerto rico", "qatar", "reunion", "romania", "russian federation", "rwanda",
	"saint kitts and nevis", "saint lucia", "saint vincent and the grenadines",
	"samoa", "san marino", "sao tome and principe", "saudi arabia", "senegal",
	"seychelles", "sierra leone", "singapore", "slovakia (slovak republic)",
	"slovenia", "solomon islands", "somalia", "south africa",
	"south georgia and the south sandwich islands", "spain", "sri lanka", "st. helena",
	"st. pierre and miquelon", "sudan", "suriname", "svalbard and jan mayen islands", "swaziland",
	"sweden", "switzerland", "syrian arab republic", "taiwan, province of china", "tajikistan",
	"tanzania, united republic of", "thailand", "togo", "tokelau", "tonga", "trinidad and tobago",
	"tunisia", "turkey", "turkmenistan", "turks and caicos islands", "tuvalu", "uganda", "ukraine",
	"united arab emirates", "united kingdom", "united states of america", "united states minor outlying islands",
	"uruguay", "uzbekistan", "vanuatu", "venezuela", "vietnam", "virgin islands (british)",
	"virgin islands (u.s.)", "wallis and futuna islands", "western sahara", "yemen", "yugoslavia",
	"zambia", "zimbabwe",
}

var replaceMap = map[string]string{
	"argentina":                "Argentina_-_Long",
	"united states of america": "U.S.A",
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func ucfirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToUpper(v))
		return u + str[len(u):]
	}

	return ""
}
