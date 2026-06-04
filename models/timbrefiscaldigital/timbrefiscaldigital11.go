package timbrefiscaldigital

import (
	"encoding/xml"
	"strings"
	"time"

	"github.com/SaulEnriqueMR/sat-serializer-xml-golang/helpers"
)

// Namespace es tfd.

// TimbreFiscalDigital11 Versión 1.1 del complemento TimbreFiscalDigital.
type TimbreFiscalDigital11 struct {
	Version          string    `xml:"Version,attr" bson:"Version" json:"Version"`
	Uuid             string    `xml:"UUID,attr" bson:"UUID" json:"UUID"`
	FechaTimbrado    time.Time `bson:"FechaTimbrado" json:"FechaTimbrado"`
	RfcProvCertif    string    `xml:"RfcProvCertif,attr" bson:"RfcProvCertif" json:"RfcProvCertif"`
	Leyenda          *string   `xml:"Leyenda,attr" bson:"Leyenda,omitempty" json:"Leyenda,omitempty"`
	SelloCFD         string    `xml:"SelloCFD,attr" bson:"SelloCFD" json:"SelloCFD"`
	NoCertificadoSAT string    `xml:"NoCertificadoSAT,attr" bson:"NoCertificadoSAT" json:"NoCertificadoSAT"`
	SelloSAT         string    `xml:"SelloSAT,attr" bson:"SelloSAT" json:"SelloSAT"`
}

func (t *TimbreFiscalDigital11) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type tmp struct {
		Version          string  `xml:"Version,attr"`
		Uuid             string  `xml:"UUID,attr"`
		FechaTimbradoRaw string  `xml:"FechaTimbrado,attr"`
		RfcProvCertif    string  `xml:"RfcProvCertif,attr"`
		Leyenda          *string `xml:"Leyenda,attr"`
		SelloCfd         string  `xml:"SelloCFD,attr"`
		NoCertificadoSat string  `xml:"NoCertificadoSAT,attr"`
		SelloSat         string  `xml:"SelloSAT,attr"`
	}
	var aux tmp
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	parsed, err := helpers.ParseDatetime(aux.FechaTimbradoRaw)
	if err != nil {
		return err
	}
	t.Version = aux.Version
	t.Uuid = strings.ToUpper(aux.Uuid)
	t.FechaTimbrado = parsed
	t.SelloCFD = aux.SelloCfd
	t.NoCertificadoSAT = aux.NoCertificadoSat
	t.SelloSAT = aux.SelloSat
	return nil
}
