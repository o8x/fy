//go:build embed

package fy

import (
	_ "embed"
)

//go:embed lib/GeoLite2-ASN.mmdb
var asnBytes []byte

//go:embed lib/GeoLite2-City.mmdb
var mmdbBytes []byte

func MustLoadDB() {

}
