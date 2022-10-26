package fy

import (
	"log"

	"github.com/oschwald/maxminddb-golang"
)

//type Names struct {
//	De   string `maxminddb:"de" json:"de"`
//	En   string `maxminddb:"en" json:"en"`
//	Fr   string `maxminddb:"fr" json:"fr"`
//	Ja   string `maxminddb:"ja" json:"ja"`
//	Ru   string `maxminddb:"ru" json:"ru"`
//	ZhCN string `maxminddb:"zh-CN" json:"zh-CN"`
//}

type Names map[string]string

type City struct {
	GeonameId int   `maxminddb:"geoname_id" json:"geoname_id"`
	Names     Names `maxminddb:"names" json:"names"`
}

type Continent struct {
	Code      string `maxminddb:"code" json:"code"`
	GeonameId int    `maxminddb:"geoname_id" json:"geoname_id"`
	Names     Names  `maxminddb:"names" json:"names"`
}

type Country struct {
	IsoCode   string `maxminddb:"iso_code" json:"iso_code"`
	GeonameId int    `maxminddb:"geoname_id" json:"geoname_id"`
	Names     Names  `maxminddb:"names" json:"names"`
}
type RegisteredCountry struct {
	GeonameId int    `maxminddb:"geoname_id" json:"geoname_id"`
	IsoCode   string `maxminddb:"iso_code" json:"iso_code"`
	Names     Names  `maxminddb:"names" json:"names"`
}

type Location struct {
	Aadius    int     `maxminddb:"aadius" json:"aadius"`
	Latitude  float64 `maxminddb:"latitude" json:"latitude"`
	Longitude float64 `maxminddb:"longitude" json:"longitude"`
	TimeZone  string  `maxminddb:"time_zone" json:"time_zone"`
}

type Subdivision struct {
	GeonameId int    `maxminddb:"geoname_id" json:"geoname_id"`
	IsoCode   string `maxminddb:"iso_code" json:"iso_code"`
	Names     Names  `maxminddb:"names" json:"names"`
}

type GeoCity struct {
	City              City              `maxminddb:"city" json:"city"`
	Continent         Continent         `maxminddb:"continent" json:"continent"`
	Location          Location          `maxminddb:"location" json:"location"`
	RegisteredCountry RegisteredCountry `maxminddb:"registered_country" json:"registered_country"`
	Country           Country           `maxminddb:"country" json:"country"`
	Subdivisions      []Subdivision     `maxminddb:"subdivisions" json:"subdivisions"`
}

func InitCity() *maxminddb.Reader {
	db, err := maxminddb.FromBytes(mmdbBytes)
	if err != nil {
		log.Fatal("init failed", err)
	}

	return db
}
