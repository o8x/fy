package lang

type Translate struct {
	TextPrefix string `json:"text_prefix"`
	TextFrom   string `json:"text_from"`
	Unknown    string `json:"unknown"`
}

var langs = map[string]Translate{
	"de":    {"ORIGIN:", " FROM:", "Unknown"},
	"en":    {"ORIGIN:", " FROM:", "Unknown"},
	"fr":    {"ORIGIN:", " FROM:", "Unknown"},
	"ja":    {"ORIGIN:", " FROM:", "Unknown"},
	"ru":    {"ORIGIN:", " FROM:", "Unknown"},
	"zh-CN": {"当前IP：", " 来自：", "未知"},
}

func GetTranslate(lang string) *Translate {
	t, ok := langs[lang]
	if ok {
		return &t
	}
	return nil
}
