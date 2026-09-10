package cfdi

import "encoding/xml"

type Comprobante10 struct {
	XMLName xml.Name `xml:"Comprobante"`
	Version string   `xml:"version,attr" bson:"Version" json:"Version"`
}
