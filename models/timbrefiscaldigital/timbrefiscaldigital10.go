package timbrefiscaldigital

import (
	"encoding/xml"
	"time"

	"github.com/SaulEnriqueMR/sat-serializer-xml-golang/models/helpers"
)

// Namespace es tfd.

// TimbreFiscalDigital10 Versión 1.0 del complemento TimbreFiscalDigital.
type TimbreFiscalDigital10 struct {
	Version          string    `xml:"version,attr" bson:"Version" json:"Version"`
	Uuid             string    `xml:"UUID,attr" bson:"Uuid" json:"Uuid"`
	FechaTimbrado    time.Time `bson:"FechaTimbrado" json:"FechaTimbrado"`
	SelloCfd         string    `xml:"selloCFD,attr" bson:"SelloCfd" json:"SelloCfd"`
	NoCertificadoSat string    `xml:"noCertificadoSAT,attr" bson:"NoCertificadoSat" json:"NoCertificadoSat"`
	SelloSat         string    `xml:"selloSAT,attr" bson:"SelloSat" json:"SelloSat"`
}

func (t *TimbreFiscalDigital10) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type tmp struct {
		Version          string `xml:"version,attr"`
		Uuid             string `xml:"UUID,attr"`
		FechaRaw         string `xml:"FechaTimbrado,attr"`
		SelloCfd         string `xml:"selloCFD,attr"`
		NoCertificadoSat string `xml:"noCertificadoSAT,attr"`
		SelloSat         string `xml:"selloSAT,attr"`
	}

	var aux tmp

	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	parsed, err := helpers.ParseDatetime(aux.FechaRaw)
	if err != nil {
		return err
	}

	t.Version = aux.Version
	t.Uuid = aux.Uuid
	t.FechaTimbrado = parsed
	t.SelloCfd = aux.SelloCfd
	t.NoCertificadoSat = aux.NoCertificadoSat
	t.SelloSat = aux.SelloSat

	return nil
}
