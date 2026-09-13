package cfdi

import (
	"github.com/SaulEnriqueMR/sat-serializer-xml-golang/serializers/cfdi/complementos"
)

type ComplementoConcepto struct {
}

type Complemento struct {
	Pagos10  *complementos.Pagos10  `xml:"http://www.sat.gob.mx/Pagos Pagos,omitempty"   json:"Pagos10,omitempty"`
	Pagos20  *complementos.Pagos20  `xml:"http://www.sat.gob.mx/Pagos20 Pagos,omitempty" json:"Pagos20,omitempty"`
	Nomina12 *complementos.Nomina12 `xml:"http://www.sat.gob.mx/nomina12 Nomina,omitempty"  json:"Nomina12,omitempty"`
}

type Addenda struct {
}
