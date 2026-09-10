package cfdi

import "encoding/xml"

type Comprobante30 struct {
	XMLName xml.Name `xml:"http://www.sat.gob.mx/cfd/3 Comprobante"`
	Version string   `xml:"version,attr" bson:"Version" json:"Version"`
}
