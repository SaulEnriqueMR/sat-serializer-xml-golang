package pago

import "encoding/xml"

type Pagos10 struct {
	XMLName xml.Name `xml:"http://www.sat.gob.mx/Pagos Pagos"`
	Version string   `xml:"Version,attr" bson:"Version" json:"Version"`
}
