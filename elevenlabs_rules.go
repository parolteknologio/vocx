package vocx

// TODO: Add fragement grouping to avoid needing the "sssij" reflexive rule.

const defaultRules = `
{
	"letters": {
		"a": "a",
		"b": "b",
		"c": "ts",
		"ĉ": "cz",
		"d": "d",
		"e": "e",
		"f": "f",
		"g": "g",
		"ĝ": "dż",
		"h": "h",
		"ĥ": "ch",
		"i": "i",
		"j": "y",
		"ĵ": "rz",
		"k": "k",
		"l": "l",
		"m": "m",
		"n": "n",
		"o": "o",
		"p": "p",
		"r": "r",
		"s": "s",
		"ŝ": "sz",
		"t": "t",
		"u": "u",
		"ŭ": "ł",
		"v": "w",
		"z": "z"
	},
	"fragments": [
		{ "match": "tsx", "replace": "cz" },
		{ "match": "gx", "replace": "dż" },
		{ "match": "hx", "replace": "ch" },
		{ "match": "yx", "replace": "rz" },
		{ "match": "sx", "replace": "sz" },
		{ "match": "ux", "replace": "ł" },
		{ "match": "atsij", "replace": "atssij" },
		{ "match": "^ekzij", "replace": "ekzji" },
		{ "match": "tssijl", "replace": "tssil" },
		{ "match": "ijuy", "replace": "iuyy" },
		{ "match": "ijeh", "replace": "ije" },
		{ "match": "sijlo", "replace": "ssilo" },
		{ "match": "^sij", "replace": "syy" },
		{ "match": "tsij", "replace": "tssij" },
		{ "match": "sij", "replace": "ssij" },
		{ "match": "sssij", "replace": "ssij" },
		{ "match": "rijpozij", "replace": "ryypozyj" },
		{ "match": "zijs", "replace": "zyjs" }
	],
	"overrides": [
		{ "eo": "ok", "pl": "ohk" },
		{ "eo": "s-ro", "pl": "sjijnjoro" },
		{ "eo": "s-ino", "pl": "sjijnjorijno" },
		{ "eo": "ktp", "pl": "ko-to-po" },
		{ "eo": "k.t.p", "pl": "ko-to-po" },
		{ "eo": "atm", "pl": "antałtagmeze" },
		{ "eo": "ptm", "pl": "posttagmeze" },
		{ "eo": "bv", "pl": "bonvolu" }
	],
	"numbers": {
		"0": "nulo",
		"1": "unu",
		"2": "du",
		"3": "tri",
		"4": "kvar",
		"5": "kvin",
		"6": "ses",
		"7": "sep",
		"8": "ohk",
		"9": "nał",
		"10": "dek",
		"100": "tsent",
		"1000": "mijl",
		"1000000": "mijlijono"
	}
} 
