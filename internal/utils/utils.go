package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/o8x/fy"
	"github.com/o8x/fy/internal/lang"
)

func WriteResponse(w http.ResponseWriter, content string) {
	_, _ = io.WriteString(w, content)
}

func ParseKeyFromQueryString(r *http.Request, key, defVal string) string {
	uri, err := url.ParseRequestURI(r.RequestURI)
	if err != nil {
		return defVal
	}

	if uri.Query().Has(key) {
		return uri.Query().Get(key)
	}
	return defVal
}

func KeyInQueryString(r *http.Request, key string) bool {
	uri, err := url.ParseRequestURI(r.RequestURI)
	if err != nil {
		return false
	}

	return uri.Query().Has(key)
}

func ParseIPFromRequest(r *http.Request) string {
	remoteIP := ""
	if r.Header.Get("Cdn-Loop") == "cloudflare" {
		remoteIP = r.Header.Get("Cf-Connecting-Ip")
		if remoteIP == "" {
			remoteIP = r.Header.Get("X-Forwarded-For")
		}
	}

	if remoteIP == "" {
		remoteIP = r.Header.Get("X-Real-IP")
	}

	if remoteIP == "" {
		remoteIP = strings.Split(r.RemoteAddr, ":")[0]
	}

	return remoteIP
}

func FormatOutput(ori *fy.Origin, language, format string) (string, error) {
	t := lang.GetTranslate(language)
	if t == nil {
		return "", fmt.Errorf("invalid lang: %s", language)
	}

	var res string
	switch format {
	case "text":
		res = fy.FormatText(ori, t)
	case "json":
		res = fy.FormatJSON(ori)
	case "xml":
		res = fy.FormatXML(ori)
	case "multiline", "ml":
		res = fy.FormatMultiline(ori)
	default:
		return "", fmt.Errorf("invalid format: %s", format)
	}

	return fmt.Sprintf("%s\n", res), nil
}
