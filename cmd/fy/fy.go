package main

import (
	"fmt"
	"net"
	"os"

	"github.com/o8x/fy"
	"github.com/o8x/fy/internal/server"
	"github.com/o8x/fy/internal/utils"

	"github.com/spf13/cobra"
)

func main() {
	format := ""
	bootServer := false
	lang := ""
	ip := ""

	cmd := &cobra.Command{
		Use: "query ip",
		Args: func(cmd *cobra.Command, args []string) error {
			if bootServer || ip != "" {
				return nil
			}
			return fmt.Errorf("miss ip or server argment")
		},
		Run: func(cmd *cobra.Command, domains []string) {
			if bootServer {
				server.ListenAndServe()
				return
			}

			ori := fy.LookupIP(net.ParseIP(ip), lang)
			if ori == nil {
				os.Exit(1)
			}

			res, err := utils.FormatOutput(ori, lang, format)
			if err != nil {
				os.Exit(1)
			}
			fmt.Print(res)
		},
		Example: `    fycli -s
    fycli -a -l en -f json 8.8.8.8
    go run github.com/o8x/fy/cmd/fy --server
    go run github.com/o8x/fy/cmd/fy -a 8.8.8.8
`,
	}

	cmd.Flags().StringVarP(&format, "format", "f", "text", "text, json, xml, multiline")
	cmd.Flags().StringVarP(&lang, "lang", "l", "zh-CN", "de,en,fr,ja,ru,zh-CN")
	cmd.Flags().StringVarP(&ip, "ip", "a", "", "query this IP address")
	cmd.Flags().BoolVarP(&bootServer, "server", "s", false, "start web server")
	_ = cmd.Execute()
}
