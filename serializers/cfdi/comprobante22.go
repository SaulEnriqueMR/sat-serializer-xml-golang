package cfdi

import "encoding/xml"

type Comprobante22 struct {
	XMLName xml.Name `xml:"http://www.sat.gob.mx/cfd/2 Comprobante"`
	Version string   `xml:"version,attr" bson:"Version" json:"Version"`
}
