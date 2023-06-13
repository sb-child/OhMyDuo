package consts

import (
	"encoding/base64"
	"unicode"

	"github.com/gogf/gf/v2/encoding/gjson"
)

func MyDuoCharactersFromString(s string) MyDuoCharacters {
	switch capitalizeFirst(s) {
	case Duo.ToString():
		return Duo
	case Bea.ToString():
		return Bea
	case Vikram.ToString():
		return Vikram
	case Oscar.ToString():
		return Oscar
	case Junior.ToString():
		return Junior
	case Eddy.ToString():
		return Eddy
	case Zari.ToString():
		return Zari
	case Lily.ToString():
		return Lily
	case Lin.ToString():
		return Lin
	case Lucy.ToString():
		return Lucy
	case Falstaff.ToString():
		return Falstaff
	}
	return Duo
}

func (x MyDuoCharacters) ToString() string {
	switch x {
	case Duo:
		return "Duo"
	case Bea:
		return "Bea"
	case Vikram:
		return "Vikram"
	case Oscar:
		return "Oscar"
	case Junior:
		return "Junior"
	case Eddy:
		return "Eddy"
	case Zari:
		return "Zari"
	case Lily:
		return "Lily"
	case Lin:
		return "Lin"
	case Lucy:
		return "Lucy"
	case Falstaff:
		return "Falstaff"
	}
	return "Duo"
}

func MyDuoLanguageFromString(s string) MyDuoLanguages {
	switch capitalizeFirst(s) {
	case "english", "en":
		return English
	}
	return English
}

func (x MyDuoLanguages) ToString() string {
	switch x {
	case English:
		return "English"
	}
	return "English"
}

func capitalizeFirst(s string) string {
	for i, c := range s {
		if unicode.IsLetter(c) {
			return string(unicode.ToUpper(c)) + s[i+1:]
		}
	}
	return s
}

func (x MyDuoElements) ToUrl(base string) string {
	enc, err := gjson.Encode(x)
	if err != nil {
		return base + "/_"
	}
	return base + "/b/" + base64.RawURLEncoding.EncodeToString(enc)
}
