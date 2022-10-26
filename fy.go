package fy

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net"
	"strings"

	"github.com/o8x/fy/internal/lang"

	"github.com/oschwald/maxminddb-golang"
)

type Origin struct {
	Origin    string            `json:"origin"`
	Country   string            `json:"country"`
	Province  string            `json:"province"`
	City      string            `json:"city"`
	ASN       string            `json:"asn"`
	Longitude string            `json:"longitude"`
	Latitude  string            `json:"latitude"`
	Headers   map[string]string `json:"headers,omitempty" xml:"-"`
}

var Default = New()

type FindYou struct {
	city *maxminddb.Reader
	asn  *maxminddb.Reader
}

func New() *FindYou {
	MustLoadDB()

	return &FindYou{
		city: InitCity(),
		asn:  InitASN(),
	}
}

func (f *FindYou) LookupASN(ip net.IP, result interface{}) error {
	return f.asn.Lookup(ip, result)
}

func (f *FindYou) LookupCity(ip net.IP, result interface{}) error {
	return f.city.Lookup(ip, result)
}

func FormatXML(origin *Origin) string {
	indent, err := xml.MarshalIndent(origin, "", "    ")
	if err != nil {
		indent, _ = xml.Marshal(origin)
	}
	return string(indent)
}

func FormatJSON(origin *Origin) string {
	marshal, _ := json.Marshal(origin)
	buf := bytes.NewBuffer(nil)
	if err := json.Indent(buf, marshal, "", "    "); err != nil {
		return string(marshal)
	}

	return buf.String()
}

func FormatMultiline(origin *Origin) string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("origin: %s \n", origin.Origin))
	builder.WriteString(fmt.Sprintf("country: %s \n", origin.Country))
	builder.WriteString(fmt.Sprintf("province: %s \n", origin.Province))
	builder.WriteString(fmt.Sprintf("city: %s \n", origin.City))
	builder.WriteString(fmt.Sprintf("asn: %s \n", origin.ASN))
	builder.WriteString(fmt.Sprintf("lon/lat: %s/%s", origin.Longitude, origin.Latitude))

	var headers []string
	for k, v := range origin.Headers {
		headers = append(headers, fmt.Sprintf("header.%s: %s", k, v))
	}

	if headers != nil {
		builder.WriteString("\n")
		builder.WriteString(strings.Join(headers, "\n"))
	}

	return builder.String()
}

func FormatText(origin *Origin, trans *lang.Translate) string {
	if origin.Country == "" {
		origin.Country = trans.Unknown
	}

	if origin.Province == "" {
		origin.Province = trans.Unknown
	}

	if origin.City == "" {
		origin.City = trans.Unknown
	}

	if origin.ASN == "" {
		origin.ASN = trans.Unknown
	}

	if origin.Longitude == "" {
		origin.Longitude = "0.000000"
	}

	if origin.Latitude == "" {
		origin.Latitude = "0.000000"
	}

	builder := strings.Builder{}
	builder.WriteString(trans.TextPrefix)
	builder.WriteString(origin.Origin)
	builder.WriteString(trans.TextFrom)
	builder.WriteString(fmt.Sprintf("%s ", origin.Country))
	builder.WriteString(fmt.Sprintf("%s ", origin.Province))
	builder.WriteString(fmt.Sprintf("%s ", origin.City))
	builder.WriteString(fmt.Sprintf("%s ", origin.ASN))
	builder.WriteString(fmt.Sprintf("%s/%s", origin.Longitude, origin.Latitude))

	return builder.String()
}

func LookupASN(addr net.IP) (*ASN, error) {
	var al ASN
	if err := Default.LookupASN(addr, &al); err != nil {
		return nil, err
	}

	al.Name = al.ISPOrganization
	return &al, nil
}

func LookupCity(addr net.IP) (*GeoCity, error) {
	var cl GeoCity

	if err := Default.LookupCity(addr, &cl); err != nil {
		return nil, err
	}

	return &cl, nil
}

func LookupIP(addr net.IP, lang string) *Origin {
	city, err := LookupCity(addr)
	if err != nil {
		return nil
	}

	asn, _ := LookupASN(addr)
	origin := Origin{
		Origin:    addr.String(),
		Country:   city.Country.Names[lang],
		City:      city.City.Names[lang],
		Longitude: fmt.Sprintf("%f", city.Location.Longitude),
		Latitude:  fmt.Sprintf("%f", city.Location.Latitude),
	}

	if asn != nil {
		origin.ASN = asn.Name
	}

	if len(city.Subdivisions) > 0 {
		origin.Province = city.Subdivisions[0].Names[lang]
	}

	return &origin
}
