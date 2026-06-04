package timbrefiscaldigital

import (
	"encoding/xml"
	"strings"
	"time"

	"github.com/SaulEnriqueMR/sat-serializer-xml-golang/helpers"
)

// Namespace es tfd.

// TimbreFiscalDigital10 Versión 1.0 del complemento TimbreFiscalDigital.
type TimbreFiscalDigital10 struct {
	Version          string    `xml:"version,attr" bson:"Version" json:"Version"`
	Uuid             string    `xml:"UUID,attr" bson:"Uuid" json:"Uuid"`
	FechaTimbrado    time.Time `bson:"FechaTimbrado" json:"FechaTimbrado"`
	SelloCFD         string    `xml:"selloCFD,attr" bson:"SelloCFD" json:"SelloCFD"`
	NoCertificadoSAT string    `xml:"noCertificadoSAT,attr" bson:"NoCertificadoSAT" json:"NoCertificadoSAT"`
	SelloSAT         string    `xml:"selloSAT,attr" bson:"SelloSAT" json:"SelloSAT"`
}

func (t *TimbreFiscalDigital10) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type tmp struct {
		Version          string `xml:"version,attr"`
		Uuid             string `xml:"UUID,attr"`
		FechaTimbradoRaw string `xml:"FechaTimbrado,attr"`
		SelloCfd         string `xml:"selloCFD,attr"`
		NoCertificadoSat string `xml:"noCertificadoSAT,attr"`
		SelloSat         string `xml:"selloSAT,attr"`
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
