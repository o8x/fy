package fy

import (
	"log"

	"github.com/oschwald/maxminddb-golang"
)

type ASN struct {
	ISPNumber       int    `maxminddb:"autonomous_system_number" json:"autonomous_system_number"`
	ISPOrganization string `maxminddb:"autonomous_system_organization" json:"autonomous_system_organization"`
	Name            string `json:"name"`
}

func InitASN() *maxminddb.Reader {
	db, err := maxminddb.FromBytes(asnBytes)
	if err != nil {
		log.Fatal("init asnBytes failed", err)
	}

	return db
}
