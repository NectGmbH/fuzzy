package normalize

import (
	"strings"
)

// ReplaceUnicodeToICAO replaces Unicode characters to ICAO-MRZ accepted ones
// as defined in ICAO doc 9303, Part 3, Section 6
func ReplaceUnicodeToICAO() Fn {
	r := strings.NewReplacer(
		"À", "A",
		"Á", "A",
		"Â", "A",
		"Ã", "A",
		"Ä", "AE",
		"Å", "AA",
		"Æ", "AE",
		"Ç", "C",
		"È", "E",
		"É", "E",
		"Ê", "E",
		"Ë", "E",
		"Ì", "I",
		"Í", "I",
		"Î", "I",
		"Ï", "I",
		"Ð", "D",
		"Ñ", "N",
		"Ò", "O",
		"Ó", "O",
		"Ô", "O",
		"Õ", "O",
		"Ö", "OE",
		"Ø", "OE",
		"Ù", "U",
		"Ú", "U",
		"Û", "U",
		"Ü", "UE",
		"Ý", "Y",
		"Þ", "TH",
		"ß", "SS",
		"Ā", "A",
		"Ă", "A",
		"Ą", "A",
		"Ć", "C",
		"Ĉ", "C",
		"Ċ", "C",
		"Č", "C",
		"Ď", "D",
		"Đ", "D",
		"Ē", "E",
		"Ĕ", "E",
		"Ė", "E",
		"Ę", "E",
		"Ě", "E",
		"Ĝ", "G",
		"Ğ", "G",
		"Ġ", "G",
		"Ģ", "G",
		"Ĥ", "H",
		"Ħ", "H",
		"Ĩ", "I",
		"Ī", "I",
		"Ĭ", "I",
		"Į", "I",
		"İ", "I",
		"ı", "I",
		"Ĳ", "IJ",
		"Ĵ", "J",
		"Ķ", "K",
		"Ĺ", "L",
		"Ļ", "L",
		"Ľ", "L",
		"Ŀ", "L",
		"Ł", "L",
		"Ń", "N",
		"Ņ", "N",
		"Ň", "N",
		"Ŋ", "N",
		"Ō", "O",
		"Ŏ", "O",
		"Ő", "O",
		"Œ", "OE",
		"Ŕ", "R",
		"Ŗ", "R",
		"Ř", "R",
		"Ś", "S",
		"Ŝ", "S",
		"Ş", "S",
		"Š", "S",
		"Ţ", "T",
		"Ť", "T",
		"Ŧ", "T",
		"Ũ", "U",
		"Ū", "U",
		"Ŭ", "U",
		"Ů", "U",
		"Ű", "U",
		"Ų", "U",
		"Ŵ", "W",
		"Ŷ", "Y",
		"Ÿ", "Y",
		"Ź", "Z",
		"Ż", "Z",
		"Ž", "Z",
		"Ё", "E",
		"Ђ", "D",
		"Є", "IE",
		"Ѕ", "DZ",
		"І", "I",
		"Ї", "I",
		"Ј", "J",
		"Љ", "LJ",
		"Њ", "NJ",
		"Ќ", "K",
		"Ў", "U",
		"Џ", "DZ",
		"А", "A",
		"Б", "B",
		"В", "V",
		"Г", "G",
		"Д", "D",
		"Е", "E",
		"Ж", "ZH",
		"З", "Z",
		"И", "I",
		"Й", "I",
		"К", "K",
		"Л", "L",
		"М", "M",
		"Н", "N",
		"О", "O",
		"П", "P",
		"Р", "R",
		"С", "S",
		"Т", "T",
		"У", "U",
		"Ф", "F",
		"Х", "KH",
		"Ц", "TS",
		"Ч", "CH",
		"Ш", "SH",
		"Щ", "SHCH",
		"Ъ", "IE",
		"Ы", "Y",
		"Э", "E",
		"Ю", "IU",
		"Я", "IA",
		"Ѫ", "U",
		"Ѵ", "Y",
		"Ґ", "G",
		"Ғ", "G",
		"Һ", "C",
		"ء", "XE",
		"آ", "XAA",
		"أ", "XAE",
		"ؤ", "U",
		"إ", "I",
		"ئ", "XI",
		"ا", "A",
		"ب", "B",
		"ة", "XTA",
		"ت", "T",
		"ث", "XTH",
		"ج", "J",
		"ح", "XH",
		"خ", "XKH",
		"د", "D",
		"ذ", "XDH",
		"ر", "R",
		"ز", "Z",
		"س", "S",
		"ش", "XSH",
		"ص", "XSS",
		"ض", "XDZ",
		"ط", "XTT",
		"ظ", "XZZ",
		"ع", "E",
		"غ", "G",
		"ـ", "",
		"ف", "F",
		"ق", "Q",
		"ك", "K",
		"ل", "L",
		"م", "M",
		"ن", "N",
		"ه", "H",
		"و", "W",
		"ى", "XAY",
		"ي", "Y",
		"ً", "",
		"ٌ", "",
		"ٍ", "",
		"َ", "",
		"ُ", "",
		"ِ", "",
		"ّ", "",
		"ْ", "",
		"ٰ", "",
		"ٱ", "XXA",
		"ٹ", "XXT",
		"ټ", "XRT",
		"پ", "P",
		"ځ", "XKE",
		"څ", "XXH",
		"چ", "XC",
		"ڈ", "XXD",
		"ډ", "XDR",
		"ڑ", "XXR",
		"ړ", "XRR",
		"ږ", "XRX",
		"ژ", "XJ",
		"ښ", "XXS",
		"ڜ", "",
		"ڢ", "",
		"ڧ", "",
		"ڨ", "",
		"ک", "XKK",
		"ګ", "XXK",
		"ڭ", "XNG",
		"گ", "XGG",
		"ں", "XNN",
		"ڼ", "XXN",
		"ھ", "XDO",
		"ۀ", "XYH",
		"ہ", "XXG",
		"ۂ", "XGE",
		"ۃ", "XTG",
		"ی", "XYA",
		"ۍ", "XXY",
		"ې", "Y",
		"ے", "XYB",
		"ۓ", "XBE",
	)

	return func(s string) string {
		return r.Replace(s)
	}
}
