package iso6391

// Language is an ISO 639-1 language with code, name and native name.
type Language struct {
	Code       string
	Name       string
	NativeName string
}

// Languages is a map of all ISO 639-1 languages using the two character lowercase language code as key.
var Languages = map[string]Language{
	"aa":    {Code: "aa", Name: "Afar", NativeName: "Afaraf"},
	"ab":    {Code: "ab", Name: "Abkhaz", NativeName: "аҧсуа бызшәа"},
	"ae":    {Code: "ae", Name: "Avestan", NativeName: "avesta"},
	"af":    {Code: "af", Name: "Afrikaans", NativeName: "Afrikaans"},
	"ak":    {Code: "ak", Name: "Akan", NativeName: "Akan"},
	"am":    {Code: "am", Name: "Amharic", NativeName: "አማርኛ"},
	"an":    {Code: "an", Name: "Aragonese", NativeName: "aragonés"},
	"ar":    {Code: "ar", Name: "Arabic", NativeName: "اللغة العربية"},
	"as":    {Code: "as", Name: "Assamese", NativeName: "অসমীয়া"},
	"av":    {Code: "av", Name: "Avaric", NativeName: "авар мацӀ"},
	"ay":    {Code: "ay", Name: "Aymara", NativeName: "aymar aru"},
	"az":    {Code: "az", Name: "Azerbaijani", NativeName: "azərbaycan dili"},
	"ba":    {Code: "ba", Name: "Bashkir", NativeName: "башҡорт теле"},
	"be":    {Code: "be", Name: "Belarusian", NativeName: "беларуская мова"},
	"bg":    {Code: "bg", Name: "Bulgarian", NativeName: "български език"},
	"bh":    {Code: "bh", Name: "Bihari", NativeName: "भोजपुरी"},
	"bi":    {Code: "bi", Name: "Bislama", NativeName: "Bislama"},
	"bm":    {Code: "bm", Name: "Bambara", NativeName: "bamanankan"},
	"bn":    {Code: "bn", Name: "Bengali", NativeName: "বাংলা"},
	"bo":    {Code: "bo", Name: "Tibetan Standard", NativeName: "བོད་ཡིག"},
	"br":    {Code: "br", Name: "Breton", NativeName: "brezhoneg"},
	"bs":    {Code: "bs", Name: "Bosnian", NativeName: "bosanski jezik"},
	"ca":    {Code: "ca", Name: "Catalan", NativeName: "català"},
	"ce":    {Code: "ce", Name: "Chechen", NativeName: "нохчийн мотт"},
	"ch":    {Code: "ch", Name: "Chamorro", NativeName: "Chamoru"},
	"co":    {Code: "co", Name: "Corsican", NativeName: "corsu"},
	"cr":    {Code: "cr", Name: "Cree", NativeName: "ᓀᐦᐃᔭᐍᐏᐣ"},
	"cs":    {Code: "cs", Name: "Czech", NativeName: "čeština"},
	"cu":    {Code: "cu", Name: "Old Church Slavonic", NativeName: "ѩзыкъ словѣньскъ"},
	"cv":    {Code: "cv", Name: "Chuvash", NativeName: "чӑваш чӗлхи"},
	"cy":    {Code: "cy", Name: "Welsh", NativeName: "Cymraeg"},
	"da":    {Code: "da", Name: "Danish", NativeName: "dansk"},
	"de":    {Code: "de", Name: "German", NativeName: "Deutsch"},
	"dv":    {Code: "dv", Name: "Divehi", NativeName: "Dhivehi"},
	"dz":    {Code: "dz", Name: "Dzongkha", NativeName: "རྫོང་ཁ"},
	"ee":    {Code: "ee", Name: "Ewe", NativeName: "Eʋegbe"},
	"el":    {Code: "el", Name: "Greek", NativeName: "Ελληνικά"},
	"en":    {Code: "en", Name: "English", NativeName: "English"},
	"eo":    {Code: "eo", Name: "Esperanto", NativeName: "Esperanto"},
	"es":    {Code: "es", Name: "Spanish", NativeName: "Español"},
	"et":    {Code: "et", Name: "Estonian", NativeName: "eesti"},
	"eu":    {Code: "eu", Name: "Basque", NativeName: "euskara"},
	"fa":    {Code: "fa", Name: "Persian", NativeName: "فارسی"},
	"ff":    {Code: "ff", Name: "Fula", NativeName: "Fulfulde"},
	"fi":    {Code: "fi", Name: "Finnish", NativeName: "suomi"},
	"fj":    {Code: "fj", Name: "Fijian", NativeName: "Vakaviti"},
	"fo":    {Code: "fo", Name: "Faroese", NativeName: "føroyskt"},
	"fr":    {Code: "fr", Name: "French", NativeName: "Français"},
	"fy":    {Code: "fy", Name: "Western Frisian", NativeName: "Frysk"},
	"ga":    {Code: "ga", Name: "Irish", NativeName: "Gaeilge"},
	"gd":    {Code: "gd", Name: "Scottish Gaelic", NativeName: "Gàidhlig"},
	"gl":    {Code: "gl", Name: "Galician", NativeName: "galego"},
	"gn":    {Code: "gn", Name: "Guaraní", NativeName: "Avañeẽ"},
	"gu":    {Code: "gu", Name: "Gujarati", NativeName: "ગુજરાતી"},
	"gv":    {Code: "gv", Name: "Manx", NativeName: "Gaelg"},
	"ha":    {Code: "ha", Name: "Hausa", NativeName: "هَوُسَ"},
	"he":    {Code: "he", Name: "Hebrew", NativeName: "עברית"},
	"hi":    {Code: "hi", Name: "Hindi", NativeName: "हिन्दी"},
	"ho":    {Code: "ho", Name: "Hiri Motu", NativeName: "Hiri Motu"},
	"hr":    {Code: "hr", Name: "Croatian", NativeName: "hrvatski jezik"},
	"ht":    {Code: "ht", Name: "Haitian", NativeName: "Kreyòl ayisyen"},
	"hu":    {Code: "hu", Name: "Hungarian", NativeName: "magyar"},
	"hy":    {Code: "hy", Name: "Armenian", NativeName: "Հայերեն"},
	"hz":    {Code: "hz", Name: "Herero", NativeName: "Otjiherero"},
	"ia":    {Code: "ia", Name: "Interlingua", NativeName: "Interlingua"},
	"id":    {Code: "id", Name: "Indonesian", NativeName: "Indonesian"},
	"ie":    {Code: "ie", Name: "Interlingue", NativeName: "Interlingue"},
	"ig":    {Code: "ig", Name: "Igbo", NativeName: "Asụsụ Igbo"},
	"ii":    {Code: "ii", Name: "Nuosu", NativeName: "ꆈꌠ꒿ Nuosuhxop"},
	"ik":    {Code: "ik", Name: "Inupiaq", NativeName: "Iñupiaq"},
	"io":    {Code: "io", Name: "Ido", NativeName: "Ido"},
	"is":    {Code: "is", Name: "Icelandic", NativeName: "Íslenska"},
	"it":    {Code: "it", Name: "Italian", NativeName: "Italiano"},
	"iu":    {Code: "iu", Name: "Inuktitut", NativeName: "ᐃᓄᒃᑎᑐᑦ"},
	"ja":    {Code: "ja", Name: "Japanese", NativeName: "日本語"},
	"jv":    {Code: "jv", Name: "Javanese", NativeName: "basa Jawa"},
	"ka":    {Code: "ka", Name: "Georgian", NativeName: "ქართული"},
	"kg":    {Code: "kg", Name: "Kongo", NativeName: "Kikongo"},
	"ki":    {Code: "ki", Name: "Kikuyu", NativeName: "Gĩkũyũ"},
	"kj":    {Code: "kj", Name: "Kwanyama", NativeName: "Kuanyama"},
	"kk":    {Code: "kk", Name: "Kazakh", NativeName: "қазақ тілі"},
	"kl":    {Code: "kl", Name: "Kalaallisut", NativeName: "kalaallisut"},
	"km":    {Code: "km", Name: "Khmer", NativeName: "ខេមរភាសា"},
	"kn":    {Code: "kn", Name: "Kannada", NativeName: "ಕನ್ನಡ"},
	"ko":    {Code: "ko", Name: "Korean", NativeName: "한국어"},
	"kr":    {Code: "kr", Name: "Kanuri", NativeName: "Kanuri"},
	"ks":    {Code: "ks", Name: "Kashmiri", NativeName: "कश्मीरी"},
	"ku":    {Code: "ku", Name: "Kurdish", NativeName: "Kurdî"},
	"kv":    {Code: "kv", Name: "Komi", NativeName: "коми кыв"},
	"kw":    {Code: "kw", Name: "Cornish", NativeName: "Kernewek"},
	"ky":    {Code: "ky", Name: "Kyrgyz", NativeName: "Кыргызча"},
	"la":    {Code: "la", Name: "Latin", NativeName: "latine"},
	"lb":    {Code: "lb", Name: "Luxembourgish", NativeName: "Lëtzebuergesch"},
	"lg":    {Code: "lg", Name: "Ganda", NativeName: "Luganda"},
	"li":    {Code: "li", Name: "Limburgish", NativeName: "Limburgs"},
	"ln":    {Code: "ln", Name: "Lingala", NativeName: "Lingála"},
	"lo":    {Code: "lo", Name: "Lao", NativeName: "ພາສາ"},
	"lt":    {Code: "lt", Name: "Lithuanian", NativeName: "lietuvių kalba"},
	"lu":    {Code: "lu", Name: "Luba-Katanga", NativeName: "Tshiluba"},
	"lv":    {Code: "lv", Name: "Latvian", NativeName: "latviešu valoda"},
	"mg":    {Code: "mg", Name: "Malagasy", NativeName: "fiteny malagasy"},
	"mh":    {Code: "mh", Name: "Marshallese", NativeName: "Kajin M̧ajeļ"},
	"mi":    {Code: "mi", Name: "Māori", NativeName: "te reo Māori"},
	"mk":    {Code: "mk", Name: "Macedonian", NativeName: "македонски јазик"},
	"ml":    {Code: "ml", Name: "Malayalam", NativeName: "മലയാളം"},
	"mn":    {Code: "mn", Name: "Mongolian", NativeName: "Монгол хэл"},
	"mr":    {Code: "mr", Name: "Marathi", NativeName: "मराठी"},
	"ms":    {Code: "ms", Name: "Malay", NativeName: "هاس ملايو‎"},
	"mt":    {Code: "mt", Name: "Maltese", NativeName: "Malti"},
	"my":    {Code: "my", Name: "Burmese", NativeName: "ဗမာစာ"},
	"na":    {Code: "na", Name: "Nauru", NativeName: "Ekakairũ Naoero"},
	"nb":    {Code: "nb", Name: "Norwegian Bokmål", NativeName: "Norsk bokmål"},
	"nd":    {Code: "nd", Name: "Northern Ndebele", NativeName: "isiNdebele"},
	"ne":    {Code: "ne", Name: "Nepali", NativeName: "नेपाली"},
	"ng":    {Code: "ng", Name: "Ndonga", NativeName: "Owambo"},
	"nl":    {Code: "nl", Name: "Dutch", NativeName: "Nederlands"},
	"nn":    {Code: "nn", Name: "Norwegian Nynorsk", NativeName: "Norsk nynorsk"},
	"no":    {Code: "no", Name: "Norwegian", NativeName: "Norsk"},
	"nr":    {Code: "nr", Name: "Southern Ndebele", NativeName: "isiNdebele"},
	"nv":    {Code: "nv", Name: "Navajo", NativeName: "Diné bizaad"},
	"ny":    {Code: "ny", Name: "Chichewa", NativeName: "chiCheŵa"},
	"oc":    {Code: "oc", Name: "Occitan", NativeName: "occitan"},
	"oj":    {Code: "oj", Name: "Ojibwe", NativeName: "ᐊᓂᔑᓈᐯᒧᐎᓐ"},
	"om":    {Code: "om", Name: "Oromo", NativeName: "Afaan Oromoo"},
	"or":    {Code: "or", Name: "Oriya", NativeName: "ଓଡ଼ିଆ"},
	"os":    {Code: "os", Name: "Ossetian", NativeName: "ирон æвзаг"},
	"pa":    {Code: "pa", Name: "Panjabi", NativeName: "ਪੰਜਾਬੀ"},
	"pi":    {Code: "pi", Name: "Pāli", NativeName: "पाऴि"},
	"pl":    {Code: "pl", Name: "Polish", NativeName: "język polski"},
	"ps":    {Code: "ps", Name: "Pashto", NativeName: "پښتو"},
	"pt":    {Code: "pt", Name: "Portuguese", NativeName: "Português"},
	"qu":    {Code: "qu", Name: "Quechua", NativeName: "Runa Simi"},
	"rm":    {Code: "rm", Name: "Romansh", NativeName: "rumantsch grischun"},
	"rn":    {Code: "rn", Name: "Kirundi", NativeName: "Ikirundi"},
	"ro":    {Code: "ro", Name: "Romanian", NativeName: "Română"},
	"ru":    {Code: "ru", Name: "Russian", NativeName: "Русский"},
	"rw":    {Code: "rw", Name: "Kinyarwanda", NativeName: "Ikinyarwanda"},
	"sa":    {Code: "sa", Name: "Sanskrit", NativeName: "संस्कृतम्"},
	"sc":    {Code: "sc", Name: "Sardinian", NativeName: "sardu"},
	"sd":    {Code: "sd", Name: "Sindhi", NativeName: "सिन्धी"},
	"se":    {Code: "se", Name: "Northern Sami", NativeName: "Davvisámegiella"},
	"sg":    {Code: "sg", Name: "Sango", NativeName: "yângâ tî sängö"},
	"si":    {Code: "si", Name: "Sinhala", NativeName: "සිංහල"},
	"sk":    {Code: "sk", Name: "Slovak", NativeName: "slovenčina"},
	"sl":    {Code: "sl", Name: "Slovene", NativeName: "slovenski jezik"},
	"sm":    {Code: "sm", Name: "Samoan", NativeName: "gagana faa Samoa"},
	"sn":    {Code: "sn", Name: "Shona", NativeName: "chiShona"},
	"so":    {Code: "so", Name: "Somali", NativeName: "Soomaaliga"},
	"sq":    {Code: "sq", Name: "Albanian", NativeName: "Shqip"},
	"sr":    {Code: "sr", Name: "Serbian", NativeName: "српски језик"},
	"ss":    {Code: "ss", Name: "Swati", NativeName: "SiSwati"},
	"st":    {Code: "st", Name: "Southern Sotho", NativeName: "Sesotho"},
	"su":    {Code: "su", Name: "Sundanese", NativeName: "Basa Sunda"},
	"sv":    {Code: "sv", Name: "Swedish", NativeName: "svenska"},
	"sw":    {Code: "sw", Name: "Swahili", NativeName: "Kiswahili"},
	"ta":    {Code: "ta", Name: "Tamil", NativeName: "தமிழ்"},
	"te":    {Code: "te", Name: "Telugu", NativeName: "తెలుగు"},
	"tg":    {Code: "tg", Name: "Tajik", NativeName: "тоҷикӣ"},
	"th":    {Code: "th", Name: "Thai", NativeName: "ไทย"},
	"ti":    {Code: "ti", Name: "Tigrinya", NativeName: "ትግርኛ"},
	"tk":    {Code: "tk", Name: "Turkmen", NativeName: "Türkmen"},
	"tl":    {Code: "tl", Name: "Tagalog", NativeName: "Wikang Tagalog"},
	"tn":    {Code: "tn", Name: "Tswana", NativeName: "Setswana"},
	"to":    {Code: "to", Name: "Tonga", NativeName: "faka Tonga"},
	"tr":    {Code: "tr", Name: "Turkish", NativeName: "Türkçe"},
	"ts":    {Code: "ts", Name: "Tsonga", NativeName: "Xitsonga"},
	"tt":    {Code: "tt", Name: "Tatar", NativeName: "татар теле"},
	"tw":    {Code: "tw", Name: "Twi", NativeName: "Twi"},
	"ty":    {Code: "ty", Name: "Tahitian", NativeName: "Reo Tahiti"},
	"ug":    {Code: "ug", Name: "Uyghur", NativeName: "ئۇيغۇرچە‎"},
	"uk":    {Code: "uk", Name: "Ukrainian", NativeName: "Українська"},
	"ur":    {Code: "ur", Name: "Urdu", NativeName: "اردو"},
	"uz":    {Code: "uz", Name: "Uzbek", NativeName: "Ўзбек"},
	"ve":    {Code: "ve", Name: "Venda", NativeName: "Tshivenḓa"},
	"vi":    {Code: "vi", Name: "Vietnamese", NativeName: "Tiếng Việt"},
	"vo":    {Code: "vo", Name: "Volapük", NativeName: "Volapük"},
	"wa":    {Code: "wa", Name: "Walloon", NativeName: "walon"},
	"wo":    {Code: "wo", Name: "Wolof", NativeName: "Wollof"},
	"xh":    {Code: "xh", Name: "Xhosa", NativeName: "isiXhosa"},
	"yi":    {Code: "yi", Name: "Yiddish", NativeName: "ייִדיש"},
	"yo":    {Code: "yo", Name: "Yoruba", NativeName: "Yorùbá"},
	"za":    {Code: "za", Name: "Zhuang", NativeName: "Saɯ cueŋƅ"},
	"zh-cn": {Code: "zh-cn", Name: "Simplified Chinese", NativeName: "简体中文"},
	"zh-tw": {Code: "zh-tw", Name: "Traditional Chinese", NativeName: "繁體中文"},
	"zu":    {Code: "zu", Name: "Zulu", NativeName: "isiZulu"},
}
