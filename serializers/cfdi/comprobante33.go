package cfdi

import "encoding/xml"

type Comprobante33 struct {
	XMLName xml.Name `xml:"http://www.sat.gob.mx/cfd/3 Comprobante"`
	Version string   `xml:"Version,attr" bson:"Version" json:"Version"`
}
