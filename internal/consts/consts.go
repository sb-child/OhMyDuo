package consts

type MyDuoCharacters uint8

const (
	Duo      MyDuoCharacters = iota
	Bea      MyDuoCharacters = iota
	Vikram   MyDuoCharacters = iota
	Oscar    MyDuoCharacters = iota
	Junior   MyDuoCharacters = iota
	Eddy     MyDuoCharacters = iota
	Zari     MyDuoCharacters = iota
	Lily     MyDuoCharacters = iota
	Lin      MyDuoCharacters = iota
	Lucy     MyDuoCharacters = iota
	Falstaff MyDuoCharacters = iota
)

type MyDuoLanguages uint8

const (
	English MyDuoLanguages = iota
)

type MyDuoElements struct {
	Rounded        bool            `json:"r"`
	Character      MyDuoCharacters `json:"c"`
	Language       MyDuoLanguages  `json:"l"`
	OriginText     string          `json:"o"`
	TranslatedText string          `json:"t"`
}

type SpiltTextPiece struct {
	Text    string
	Unicode bool
}
