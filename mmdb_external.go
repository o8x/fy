//go:build external

package fy

import (
	"log"
	"os"
	"path/filepath"
)

var asnBytes []byte

var mmdbBytes []byte

func MustLoadDB() {
	var err error

	dir, err := os.UserHomeDir()
	baseDir := filepath.Join(dir, "lib")

	log.Printf("scan mmdb files from path: %s", baseDir)

	if asnBytes, err = os.ReadFile(filepath.Join(baseDir, "GeoLite2-ASN.mmdb")); err != nil {
		panic(err)
	}

	if mmdbBytes, err = os.ReadFile(filepath.Join(baseDir, "GeoLite2-City.mmdb")); err != nil {
		panic(err)
	}
}
