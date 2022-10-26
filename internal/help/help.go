package help

import "strings"

func GetHelpText() string {
	builder := strings.Builder{}
	builder.WriteString("options: \n")
	builder.WriteString("	lang [option]   : de,en,fr,ja,ru,zh-CN default: zh-CN\n")
	builder.WriteString("	format [option] : text, json, xml, multiline,ml default: text\n")
	builder.WriteString("	ip [option]     : query this IP address instead of the client IP\n")
	builder.WriteString("	trace [option]  : result with request header when format in json or multiline\n")
	builder.WriteString("	help [option]   : display this message\n")
	builder.WriteString("\n")
	builder.WriteString("example: \n")
	builder.WriteString("	- ?help\n")
	builder.WriteString("	- ?format=json\n")
	builder.WriteString("	- ?format=json&trace\n")
	builder.WriteString("	- ?lang=en&format=xml&ip=8.8.8.8\n")
	builder.WriteString("\n")

	return builder.String()
}
