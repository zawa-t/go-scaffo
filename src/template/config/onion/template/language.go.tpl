package lang

import "{{ .BaseImportPath }}/env"

const (
	_   = iota
	ENG // English
	JPN // Japanese
	RUS // Russian
)

var LangList = map[string]int{
	"eng": ENG,
	"jpn": JPN,
	"rus": RUS,
}

func Language() int {
	if name, ok := LangList[env.Lang]; ok {
		return name
	}
	return ENG
}
