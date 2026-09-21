package cfdi

import (
	"github.com/SaulEnriqueMR/sat-serializer-xml-golang/serializers/cfdi/complementos"
)

type ComplementoConcepto struct {
}

type Complemento struct {
	Pagos10  *complementos.Pagos10  `xml:"http://www.sat.gob.mx/Pagos Pagos,omitempty" bson:"Pagos10"  json:"Pagos10,omitempty"`
	Pagos20  *complementos.Pagos20  `xml:"http://www.sat.gob.mx/Pagos20 Pagos,omitempty" bson:"Pagos20" json:"Pagos20,omitempty"`
	Nomina11 *complementos.Nomina11 `xml:"http://www.sat.gob.mx/nomina Nomina,omitempty" bson:"Nomina11" json:"Nomina11,omitempty"`
	Nomina12 *complementos.Nomina12 `xml:"http://www.sat.gob.mx/nomina12 Nomina,omitempty" bson:"Nomina12" json:"Nomina12,omitempty"`
}

type Addenda struct {
}
