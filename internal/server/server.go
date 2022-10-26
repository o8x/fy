package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/o8x/fy"
	"github.com/o8x/fy/internal/help"
	"github.com/o8x/fy/internal/utils"
)

func HomeFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	log.Printf("request from uri: %s", r.RequestURI)

	if utils.KeyInQueryString(r, "help") {
		utils.WriteResponse(w, help.GetHelpText())
		return
	}

	ip := utils.ParseKeyFromQueryString(r, "ip", "")
	if ip == "" {
		ip = utils.ParseIPFromRequest(r)
	}

	lang := utils.ParseKeyFromQueryString(r, "lang", "zh-CN")
	format := utils.ParseKeyFromQueryString(r, "format", "text")

	ori := fy.LookupIP(net.ParseIP(ip), lang)
	if ori == nil {
		utils.WriteResponse(w, fmt.Sprintf("invalid origin: %s", ip))
		return
	}

	if utils.KeyInQueryString(r, "trace") {
		var headers = map[string]string{}
		for k, v := range r.Header {
			if strings.HasPrefix(k, "Cf-") {
				continue
			}

			if len(v) == 0 {
				headers[k] = v[0]
			} else {
				headers[k] = strings.Join(v, " ; ")
			}
		}

		ori.Headers = headers
	}

	res, err := utils.FormatOutput(ori, lang, format)
	if err != nil {
		utils.WriteResponse(w, err.Error())
		return
	}

	switch format {
	case "json":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	case "xml":
		w.Header().Set("Content-Type", "application/xml; charset=utf-8")
	}

	utils.WriteResponse(w, res)
}

func ListenAndServe() {
	http.HandleFunc("/", HomeFunc)

	listenOn := ":12345"
	log.Printf("server listen on: %s", listenOn)
	if err := http.ListenAndServe(listenOn, nil); err != nil {
		log.Fatal(err)
	}
}
