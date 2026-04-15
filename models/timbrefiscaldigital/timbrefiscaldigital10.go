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
	Fecha            string    `xml:"FechaTimbrado,attr"`
	FechaTimbrado    time.Time `bson:"FechaTimbrado" json:"FechaTimbrado"`
	SelloCfd         string    `xml:"selloCFD,attr" bson:"SelloCfd" json:"SelloCfd"`
	NoCertificadoSat string    `xml:"noCertificadoSAT,attr" bson:"NoCertificadoSat" json:"NoCertificadoSat"`
	SelloSat         string    `xml:"selloSAT,attr" bson:"SelloSat" json:"SelloSat"`
}

func (t *TimbreFiscalDigital10) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias TimbreFiscalDigital10
	var aux Alias

	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	fechaTimbrado, err := helpers.ParseDatetime(aux.Fecha)
	if err != nil {
		return err
	}

	*t = TimbreFiscalDigital10(aux)
	t.FechaTimbrado = fechaTimbrado

	return nil
}
